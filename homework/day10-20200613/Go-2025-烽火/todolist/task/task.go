package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
	"todolist/config"
	"todolist/utils"
)

type TaskController struct {
	ID        int
	Name      string
	StartTime *time.Time
	EndTime   *time.Time
	Status    string
	User      string
}

type TaskResponse struct {
	Task []*TaskController
	Ok   bool
}

//json持久化
type JsonPersist struct {
	file        string
	TaskContent []*TaskController
}

func NewJson() *JsonPersist {
	return &JsonPersist{
		file:        config.Config.JsonTaskFile,
		TaskContent: Tasks,
	}
}

func (j *JsonPersist) Record() {
	ctx, err := json.Marshal(j.TaskContent)
	if err != nil {
		panic(err)
	}
	file, _ := os.Create(j.file)
	defer file.Close()
	var buffer bytes.Buffer
	json.Indent(&buffer, ctx, "", "\t")
	buffer.WriteTo(file)
}

var Tasks = make([]*TaskController, 0)
var P = NewJson()

// Task结构体初始化函数
func New() *TaskController {
	id := genID()
	now := time.Now()
	end := now.Add(24 * time.Hour)
	return &TaskController{
		ID:        id,
		StartTime: &now,
		EndTime:   &end,
		Status:    config.StatusNew,
	}
}

func init() {
	loadTask()
	P.TaskContent = Tasks
}

func loadTask() {
	if utils.FileIsExists(config.Config.JsonTaskFile) {
		jsonText, _ := ioutil.ReadFile(config.Config.JsonTaskFile)
		err := json.Unmarshal(jsonText, &Tasks)
		if err != nil {
			panic(err)
		}
	}
}

// 生成tasks最大的id
func genID() int {
	var rt int
	if len(Tasks) == 0 {
		return 1
	}

	for _, t := range Tasks {
		if rt < t.ID {
			rt = t.ID
		}
	}
	return rt + 1
}

// 验证任务名，确保唯一性
func verifyName(name string) bool {
	for _, t := range Tasks {
		if name == t.Name {
			return false
		}
	}
	return true
}

// 添加任务
func (c *TaskController) AddTask(name, startTime, endTime, status, user string) {
	task := New()

	if verifyName(name) {
		task.Name = name
	} else {
		utils.Error("任务名称已存在!")
	}
	task.StartTime = utils.String2time(startTime)
	task.EndTime = utils.String2time(endTime)
	task.Status = config.Config.StatusMap[status]
	task.User = user
	Tasks = append(Tasks, task)
	utils.Success("添加任务成功")
	P.Record()
}

// 排序任务信息
func sortTask(tasks []*TaskController, key string) {
	if key == "name" {
		sort.Slice(tasks, func(i, j int) bool { return tasks[i].Name < tasks[j].Name })
	}
	if key == "startTime" {
		sort.Slice(tasks, func(i, j int) bool {
			return utils.Time2string(tasks[i].StartTime) < utils.Time2string(tasks[j].StartTime)
		})
	}
}

// 查询任务
func ListTask(name string) *TaskResponse {
	response := &TaskResponse{}
	if name == "" {
		return response
	}
	filterTask := make([]*TaskController, 0)
	if name == "all" {
		filterTask = Tasks
	} else {
		for _, task := range Tasks {
			if strings.Contains(task.Name, name) {
				filterTask = append(filterTask, task)
			}
		}
	}
	if len(filterTask) > 0 {
		sortTask(filterTask, name)
		response.Ok = true
		response.Task = filterTask
	}
	fmt.Printf("%3v\n", response)
	return response

}

// 修改任务
func (c *TaskController) Modify(name, status, user string) {
	if len(Tasks) == 0 {
		return
	}
	for _, t := range Tasks {
		if t.Name == name {
			if config.Config.StatusMap[status] == config.StatusComplete {
				now := time.Now()
				t.EndTime = &now
			}
			t.Status = config.Config.StatusMap[status]
			t.User = user
			P.Record()
		}
	}
}

// 删除任务
func (c *TaskController) Delete(name string) {
	for index, t := range Tasks {
		if t.Name == name {
			copy(Tasks[index:], Tasks[index+1:])
			Tasks = Tasks[:len(Tasks)-1]
			P.Record()
		}
	}
}
