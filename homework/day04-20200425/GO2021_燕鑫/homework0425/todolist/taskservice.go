package todolist

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type TaskService struct {
	user      string
	tasks     []*task
	taskcount int64
	dburl     string
	lastid    int64
}

func NewTaskService(dburl string, user string) *TaskService {
	tasksrv := &TaskService{dburl: dburl}
	tasksrv.login(user)
	tasksrv.tasks, tasksrv.taskcount = tasksrv.selectAll()
	tasksrv.lastid, _ = strconv.ParseInt((*tasksrv.tasks[len(tasksrv.tasks)-1])["id"], 0, 0)
	return tasksrv
}

func (t *TaskService) GetByFilter(filterstr string) ([]*task, string, error) {
	defer trace(begin())
	filters := strings.Fields(strings.TrimSpace(filterstr))
	var rowcount int64 = 0
	result := make([]*task, 0)
	for _, task := range t.tasks {
		if len(filters) == 0 {
			rowcount = t.taskcount
			result = t.tasks
			break
		}
		for i := 0; i < len(filters); i++ {
			f := strings.Split(filters[i], "=")
			if len(f) != 2 {
				return nil, "", errors.New(`Error: please use the format like "xx=xx yy=yy" to specify filter conditions `)
			}

			if value, ok := (*task)[f[0]]; !ok {
				errstr := fmt.Sprintf("Unknown column '%s' in field list ", f[0])
				return nil, "", errors.New(errstr)
			} else if value == f[1] {
				if i == len(filters)-1 {
					result = append(result, task)
					rowcount++
				}
				continue
			} else {
				break
			}
		}
	}

	switch rowcount {
	case 0:
		fmt.Print("Empty set ")
		return nil, "Empty set ", nil
	case 1:
		fmt.Printf("1 row in set ")
		return result, "1 row in set ", nil
	default:
		resultstr := fmt.Sprintf("%d rows in set ", rowcount)
		return result, resultstr, nil
	}
}

func (t *TaskService) CreateNewTask(taskname string) error {
	if len(taskname) == 0 {
		return errors.New("You must specify a taskname!")
	}
	for _, task := range t.tasks {
		if taskname == (*task)["name"] {
			return errors.New("Duplicated task name!Please change another task name!")
		}
	}
	newtask := task{
		"id":        strconv.FormatInt(t.lastid+1, 10),
		"name":      taskname,
		"starttime": time.Now().Format("2006-01-02 15:04:05"),
		"endtime":   "",
		"status":    "created",
		"users":      t.user,
	}
	t.tasks = append(t.tasks, &newtask)
	t.taskcount++
	t.lastid++
	return nil
}

func (t *TaskService) UpdateTask(taskitem *task, cols ...string) (string, error) {
	if (*taskitem)["user"]!=t.user{
		return "",errors.New("You don't have rights to operate this task!")
	}
	if cols[0] == "" && cols[1] == "" {
		return "", errors.New("Nothing will change!")
	}
	if cols[0] != "" {
		if cols[0] != "created" && cols[0] != "running" && cols[0] != "paused" && cols[0] != "finished" {
			return "", errors.New("Status only can be one of (created,running,paused,finished)")
		}
		(*taskitem)["status"] = cols[0]
		// If status changed to finished,then we must update the endtime
		// If task hasn't been finished,then we must let endtime be empty
		if cols[0] == "finished" {
			(*taskitem)["endtime"] = time.Now().Format("2006-01-02 15:04:05")
		} else {
			(*taskitem)["endtime"] = ""
		}
	}
	if cols[1] != "" {
		(*taskitem)["user"] = cols[1]
	}
	return "Task has changed!", nil
}

func (t *TaskService) DeleteTask(taskid string) (string, error) {
	if _, err := strconv.ParseInt(taskid, 0, 0); err != nil {
		return "", errors.New("ID must be a number!Nothing will change!")
	}
	loopflag := 0
	for index, task := range t.tasks {
		if (*task)["id"] == taskid {
			if (*task)["user"]!=t.user{
				return "",errors.New("You don't have rights to operate this task!")
			}
			for i := index; i < len(t.tasks)-1; i++ {
				t.tasks[i] = t.tasks[i+1]
			}
			break
		}
		loopflag++
	}
	if loopflag == len(t.tasks) {
		return "", errors.New("This id can't be found in recent tasks,nothing will change!")
	}
	t.tasks = t.tasks[:len(t.tasks)-1]
	t.taskcount--
	return "This task has been deleted!", nil
}

func (t *TaskService) GetTaskcount() int64 {
	return t.taskcount
}

func (t *TaskService) GetDBurl() string {
	return t.dburl
}

func (t *TaskService) selectAll() ([]*task, int64) {
	var tableLines int64 = 0
	f, err := os.Open(t.dburl)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}
	defer f.Close()

	tasks := make([]*task, 0)

	br := bufio.NewReader(f)
	for {
		var onetask task
		line, _, eof := br.ReadLine()
		if eof == io.EOF {
			break
		}
		err = json.Unmarshal(line, &onetask)
		tasks = append(tasks, &onetask)
		if err != nil {
			return nil, 0
		}
		tableLines++
	}
	return tasks, tableLines
}

func (t *TaskService) reload() (resultstr string, err error) {
	t.tasks, t.taskcount = t.selectAll()
	t.lastid, err = strconv.ParseInt((*t.tasks[len(t.tasks)-1])["id"], 0, 0)
	if err != nil {
		return resultstr, err
	}
	resultstr = "RELOAD COMPLETE!"
	return resultstr, nil
}

func (t *TaskService) commit() (string, error) {

	f, err := os.OpenFile(TEMPDATAFILE, os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}

	br := bufio.NewWriter(f)
	for _, task := range t.tasks {
		taskline, err := json.Marshal(*task)
		if err != nil {
			return "", err
		}
		_, err = br.Write(taskline)
		if err != nil {
			return "", err
		}
		_, err = br.WriteString("\n")
		if err != nil {
			return "", err
		}
		err = br.Flush()
		if err != nil {
			return "", err
		}
	}
	f.Close()
	err = os.Rename(DBFILE, OLDDATAFILE)
	if err != nil {
		return "", err
	}
	err = os.Rename(TEMPDATAFILE, DBFILE)
	if err != nil {
		return "", err
	}
	return "Data has writen to the disk!You can reload the task list to use the latest data!", nil
}

func (t *TaskService) printLines(results []*task) {
	/*fmt.Printf("%-10s | %-15s | %-20s | %-20s | %-10s | %-10s\n", "id", "name", "starttime", "endtime", "status", "user")
	for _, task := range results {
		fmt.Printf("%-10s | %-15s | %-20s | %-20s | %-10s | %-10s\n", (*task)["id"], (*task)["name"], (*task)["starttime"],
			(*task)["endtime"], (*task)["status"], (*task)["user"])
	}*/

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "name", "starttime", "endtime", "status", "user"})

	for _, task := range results {
		table.Append([]string{(*task)["id"], (*task)["name"], (*task)["starttime"],
			(*task)["endtime"], (*task)["status"], (*task)["user"]})
	}
	table.Render() // Send output
}

func (t *TaskService) login(user string) {
	t.user=user
}
