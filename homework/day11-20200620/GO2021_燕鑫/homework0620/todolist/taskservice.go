package todolist

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"homework0620/logging"
	"strings"
	"time"
)

type TaskService struct {
	db        *gorm.DB
	user      string
	tasks     []*task
	taskcount int64
	lastid    int64
	//add log function
	logger *logging.Logger
}

func NewTaskService(db *gorm.DB) *TaskService {
	tasksrv := &TaskService{db: db}
	tasksrv.logger = logging.NewLogger()
	tasksrv.SetTaskcountAndLastId()
	return tasksrv
}

func (t *TaskService) SetUser(user string) {
	t.user = user
}

func (t *TaskService) GetUser() string {
	return t.user
}

func (t *TaskService) GetByFilter(filterstr string, sortstr string) (tasks []task, resultcnt int64, err error) {
	defer trace(begin())
	defer t.execlog("GetByFilter", &err)

	t.logger.Logging("[INFO]  ", fmt.Sprintf("Select from task: %s", filterstr))

	db := t.db
	filterstmp := strings.Fields(strings.TrimSpace(filterstr))
	tasks = make([]task, 0)
	for _, item := range filterstmp {
		f := strings.Split(item, "=")
		predication := fmt.Sprintf("%s = ?", f[0])
		db = db.Where(predication, f[1])
	}
	if sortstr == "" {
		db = db.Order("id")
	} else {
		db = db.Order(sortstr)
	}
	err = db.Find(&tasks).Error
	if err != nil {
		return nil, 0, err
	}
	rowcount := int64(len(tasks))
	return tasks, rowcount, nil
}

func (t *TaskService) CreateNewTask(taskname string) (err error) {
	defer trace(begin())
	defer t.execlog("CreateNewTask", &err)
	if len(taskname) == 0 {
		return errors.New("You must specify a taskname!")
	}

	newtask := &task{
		Name:      taskname,
		StartTime: time.Now().Format("2006-01-02 15:04:05"),
		EndTime:   "",
		Status:    "created",
		User:      t.user,
	}

	t.db.Create(newtask)
	t.taskcount++
	t.lastid++
	return nil
}

func (t *TaskService) UpdateTask(taskitem *task, cols ...string) (resultstr string, err error) {
	defer trace(begin())
	defer t.execlog("UpdateTask", &err)
	if taskitem.User != t.user {
		return "", errors.New("You don't have rights to operate this task!")
	}
	if cols[0] == "" && cols[1] == "" {
		return "", errors.New("Nothing will change!")
	}
	if cols[0] != "" {
		if cols[0] != "created" && cols[0] != "running" && cols[0] != "paused" && cols[0] != "finished" {
			return "", errors.New("Status only can be one of (created,running,paused,finished)")
		}
		taskitem.Status = cols[0]
		// If status changed to finished,then we must update the endtime
		// If task hasn't been finished,then we must let endtime be empty
		if cols[0] == "finished" {
			taskitem.EndTime = time.Now().Format("2006-01-02 15:04:05")
		} else {
			taskitem.EndTime = ""
		}
	}
	if cols[1] != "" {
		taskitem.User = cols[1]
	}
	t.db.Model(&task{}).Save(taskitem)
	return "Task has changed!", nil
}

func (t *TaskService) DeleteTask(taskitem *task) (resultstr string, err error) {
	defer trace(begin())
	defer t.execlog("DeleteTask", &err)

	if taskitem.User != t.user {
		return "", errors.New("You don't have rights to operate this task!")
	}
	err = t.db.Delete(taskitem).Error
	if err != nil {
		return "", err
	}
	t.taskcount--
	return "This task has been deleted!", nil
}

func (t *TaskService) GetTaskcount() int64 {
	return t.taskcount
}

func (t *TaskService) SetTaskcountAndLastId() {
	var tableLines int64 = 0
	t.db.Model(&task{}).Count(&tableLines)
	task := &task{}
	t.db.Last(task)
	t.lastid = task.Id
	t.taskcount = tableLines
}


func (t *TaskService) login(user string) {
	t.user = user
}

func (t *TaskService) execlog(funcname string, err *error) {
	if (*err) != nil {
		t.logger.Logging("[ERROR]  ", fmt.Sprintf("%s: %s", funcname, (*err).Error()))
	} else {
		t.logger.Logging("[SUCCESS]  ", fmt.Sprintf("%s execute success!", funcname))
	}
}
