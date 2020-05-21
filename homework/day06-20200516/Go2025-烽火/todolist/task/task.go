package task

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"todolist/utils"

	"github.com/olekukonko/tablewriter"
)

type Task struct {
	ID        int
	Name      string
	StartTime *time.Time
	EndTime   *time.Time
	Status    string
	User      string
}

var Todolist = make([]*Task, 0)

// Task结构体初始化函数
func newTask() *Task {
	id := genID()
	now := time.Now()
	end := now.Add(24 * time.Hour)
	return &Task{
		ID:        id,
		Name:      "",
		StartTime: &now,
		EndTime:   &end,
		Status:    statusNew,
		User:      "",
	}
}

func init() {
	loadTask()
}

func loadTask() {
	if utils.FileIsExists(JsonTaskFile) {
		// file, _ := os.Open(JsonTaskFile)
		// defer file.Close()

		// 方法1: 文本存储使用bufio Scanner
		// scanner := bufio.NewScanner(file)
		// for scanner.Scan() {
		// 	line := scanner.Text()
		// 	Todolist = append(Todolist, ParseTask(line))
		// }

		// // 方法2: 文本使用bufio Reader
		// reader := bufio.NewReader(file)
		// for {
		// 	line, err := reader.ReadString('\n')
		// 	if line == "" && err != nil {
		// 		if err == io.EOF {
		// 			break
		// 		}
		// 	}
		// 	Todolist = append(Todolist, ParseTask(line))
		// }

		// 方法3: json存储使用ioutil方式
		jsonText, _ := ioutil.ReadFile(JsonTaskFile)
		err := json.Unmarshal(jsonText, &Todolist)
		if err != nil {
			panic(err)
		}

	}
}

func ParseTask(txt string) *Task {
	content := strings.Split(strings.TrimSpace(txt), ",")
	id, err := strconv.Atoi(content[0])
	if err != nil {
		panic(err)
	}
	return &Task{
		ID:        id,
		Name:      content[1],
		StartTime: string2time(content[2]),
		EndTime:   string2time(content[3]),
		Status:    content[4],
		User:      content[5],
	}
}

// string -> 日期
func string2time(txt string) *time.Time {
	date, err := time.Parse(TimeLayout, txt)
	if err != nil {
		panic(err)
	}
	return &date
}

// 格式化日期->string
func time2string(t *time.Time) string {
	return t.Format(TimeLayout)
}

// 生成tasks最大的id
func genID() int {
	var rt int
	if len(Todolist) == 0 {
		return 1
	}

	for _, task := range Todolist {
		if rt < task.ID {
			rt = task.ID
		}
	}
	return rt + 1
}

// 验证任务名，确保唯一性
func verifyName(inputName string) bool {
	for _, task := range Todolist {
		if inputName == task.Name {
			return false
		}
	}
	return true
}

//渲染输出任务信息
func RenderTask(taskInfo ...*Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetHeader(Header)
	for i := 0; i < len(Header); i++ {
		table.SetColMinWidth(i, 20)
	}
	for _, task := range taskInfo {
		table.Append([]string{
			strconv.Itoa(task.ID),
			task.Name,
			time2string(task.StartTime),
			time2string(task.EndTime),
			task.Status,
			task.User,
		})
	}
	table.Render()
	// table.AppendBulk(content)
}

// 添加任务
func AddTask() {
	task := newTask()
	fmt.Println("请输入任务信息!")

	for {
		tempName := utils.Input("任务名:")
		if verifyName(tempName) {
			task.Name = tempName
			break
		} else {
			fmt.Println("任务名称已存在!")
		}
	}
	task.User = utils.Input("负责人:")
	ChangePassword()
	Todolist = append(Todolist, task)
}

// 排序任务信息
func SortTask(taskInfo []*Task, key string) {
	if key == "name" {
		sort.Slice(taskInfo, func(i, j int) bool { return taskInfo[i].Name < taskInfo[j].Name })
	}
	if key == "startTime" {
		sort.Slice(taskInfo, func(i, j int) bool { return time2string(taskInfo[i].StartTime) < time2string(taskInfo[j].StartTime) })
	}
}

// 查询任务
func QueryTaskWithSort() {
	queryMap := map[string]string{
		"1": "name",
		"2": "startTime",
	}
	if len(Todolist) == 0 {
		fmt.Println("目前没有任何任务记录，请先添加任务，谢谢!")
		return
	}
	filterTasks := make([]*Task, 0)

	q := utils.Input("请输入查询的任务名称:")
	if q == "all" {
		filterTasks = Todolist
	} else {
		for _, task := range Todolist {
			if strings.Contains(task.Name, q) || strings.Contains(task.User, q) || q == strconv.Itoa(task.ID) {
				filterTasks = append(filterTasks, task)
			}
		}
	}

	if len(filterTasks) == 0 {
		fmt.Println("未找到关联任务!")
	} else {
		key := utils.Input("请输入排序方式[1.任务名称 2.任务开始时间]:")
		SortTask(filterTasks, queryMap[key])
		RenderTask(filterTasks...)
	}
}

// 修改任务
func ModifyTask() bool {
	flag := false
	if len(Todolist) == 0 {
		fmt.Println("目前没有任何任务记录，请先添加任务，谢谢!")
		return flag
	}
	q := utils.Input("请输入需要修改的任务ID:")
	for index, task := range Todolist {
		if id, _ := strconv.Atoi(q); id == task.ID {
			RenderTask(task)
			switch utils.Input("是否修改,请确认(y/yes):") {
			case "y", "yes":
				for {
					tempName := utils.Input("任务名称:")
					if verifyName(tempName) {
						task.Name = tempName
						break
					} else {
						fmt.Println("任务名称已存在!")
					}
				}
				for {
					tempStatus := utils.Input("状态[1.未执行 2.开始执行 3.暂停 4.完成]:")
					if status, ok := StatusMap[tempStatus]; ok {
						task.Status = status
						now := time.Now()
						task.EndTime = &now
					} else {
						fmt.Println("输入的状态值不对!")
					}
				}
				Todolist[index] = task
				flag = true
				ChangePassword()
				RenderTask(task)
			default:
				fmt.Println("取消修改!")
				break
			}
		}
	}
	return flag
}

// 删除任务
func DeleteTask() bool {
	flag := false
	queryId := utils.Input("请输入需要删除的任务ID:")
	for index, task := range Todolist {
		if qid, _ := strconv.Atoi(queryId); qid == task.ID {
			RenderTask(task)
			switch utils.Input("是否进行删除(y/yes):") {
			case "y", "yes":
				copy(Todolist[index:], Todolist[index+1:])
				Todolist = Todolist[:len(Todolist)-1]
				ChangePassword()
				fmt.Printf("任务ID:%s 已删除\n", queryId)
				flag = true
			default:
				fmt.Println("取消删除!")
				return false
			}
		}
	}
	return flag
}

// 修改密码
func ChangePassword() {
	switch utils.Input("是否需要修改密码，请确认[y/yes]: ") {
	case "y", "yes":
		utils.SetPassword()
	default:
		fmt.Println("取消修改密码!")
	}
}

// txt方式持久化任务信息
func TxtTask() {
	content := make([]string, 0)
	for _, task := range Todolist {
		taskContent := []string{
			strconv.Itoa(task.ID),
			task.Name,
			time2string(task.StartTime),
			time2string(task.EndTime),
			task.Status,
			task.User,
		}
		content = append(content, strings.Join(taskContent, ","))
	}
	utils.WriteFile(TaskFile, strings.Join(content, "\n"))
}

//gob持久化
func GobTask() {
	file, _ := os.Create(GobTaskFile)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	for _, task := range Todolist {
		taskInfo := strings.Join([]string{strconv.Itoa(task.ID), task.Name, time2string(task.StartTime), time2string(task.EndTime), task.Status, task.User}, ",")
		encoder.Encode(taskInfo)
	}
}

//csv持久化
func CsvTask() {
	taskFile := operateCsvFile()
	file, _ := os.Create(taskFile)
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	csvWriter := csv.NewWriter(writer)

	for _, task := range Todolist {
		taskInfo := []string{strconv.Itoa(task.ID), task.Name, time2string(task.StartTime), time2string(task.EndTime), task.Status, task.User}
		csvWriter.Write(taskInfo)
	}
}

// 保留最近N次的csv任务记录
func operateCsvFile() string {
	dir, _ := filepath.Split(CsvTaskFile)
	pattern := dir + "tasks.csv*"
	matchFile := make([]string, 0)
	deleteFile := make([]string, 0)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		ret, _ := filepath.Match(pattern, path)
		if !info.IsDir() && ret {
			matchFile = append(matchFile, path)
		}
		return nil
	})
	sort.Strings(matchFile)

	if len(matchFile) > RetainCsvNum-1 {
		deleteFile = matchFile[:len(matchFile)-2]
	}

	latestfile := matchFile[len(matchFile)-1]
	fileinfo := strings.Split(latestfile, ".")
	lastfileSuffix, _ := strconv.Atoi(fileinfo[len(fileinfo)-1])
	suffix := strconv.Itoa(lastfileSuffix + 1)

	// delete old files
	if len(deleteFile) > 0 {
		for _, path := range deleteFile {
			os.Remove(path)
		}
	}
	return strings.Join([]string{CsvTaskFile, suffix}, ".")

}

//json持久化
func JsonTask() {
	ctx, err := json.Marshal(Todolist)
	if err != nil {
		panic(err)
	}
	file, _ := os.Create(JsonTaskFile)
	defer file.Close()

	var buffer bytes.Buffer
	json.Indent(&buffer, ctx, "", "\t")

	buffer.WriteTo(file)
	// writer := bufio.NewWriter(file)
	// defer writer.Flush()

	// writer.WriteString(string(ctx))

}
