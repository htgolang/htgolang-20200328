package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var TasksInfo map[string]map[string]string

const (
	id        = "id"
	name      = "name"
	startTime = "startTime"
	endTime   = "endTime"
	status    = "status"
	user      = "user"
)

const (
	statusNew     = "新创建"
	statusCompele = "已完成"
	statusStart   = "开始执行"
	statusStop    = "暂停"
)

func init() {
	if dbInfo := []byte(ReadDB()); len(dbInfo) > 0 {
		err := json.Unmarshal(dbInfo, &TasksInfo)
		if err != nil {
			panic(fmt.Sprintf("无法从%s反序列化数据,请检查数据的有效性.%s", dbFile, err))
		}
	} else {
		TasksInfo = map[string]map[string]string{}
	}
}

func isExistName(taskName string) bool {
	for _, task := range TasksInfo {
		if taskName == task[name] {
			return true
		}
	}
	return false
}

func Input(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	return strings.TrimSpace(string(data))
}

func generateID() int {
	var rt int
	for taskID, _ := range TasksInfo {
		id, _ := strconv.Atoi(taskID)
		if rt < id {
			rt = id
		}
	}
	return rt + 1
}

func saveDB() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("新建任务失败,%s", err)
		}
	}()
	defer func() {
		err := WriteDB(TasksInfo)
		if err != nil {
			panic(err)
		}
	}()
}

func NewTask() {
	task := make(map[string]string)
	taskID := strconv.Itoa(generateID())
	task[name] = Input("请输入任务名称：")
	task[startTime] = Input("开始时间：")
	task[endTime] = ""
	task[status] = statusNew
	task[user] = Input("责任人：")
	if isExistName(task[name]) {
		fmt.Println("您输入的任务名已存在.")
		return
	}
	if _, err := time.ParseInLocation("2006-01-02 15:04:05", task[startTime], time.Local); err != nil {
		fmt.Printf("输入的时间格式非法，%s，创建任务失败.\n", err)
		return
	}
	TasksInfo[taskID] = task
	saveDB()
}

func QueryTask() {
	info := Input("请输入查询的信息：")
	for taskid, task := range TasksInfo {
		fmt.Println(strings.Repeat("-", 40))
		if info == "all" || strings.Contains(task[name], info) {
			fmt.Println("ID：", taskid)
			fmt.Println("任务名：", task[name])
			fmt.Println("开始时间：", task[startTime])
			fmt.Println("结束时间：", task[endTime])
			fmt.Println("任务状态：", task[status])
			fmt.Println("负责人：", task[user])
		}
	}
}

func DeleteTask() {
	taskID := Input("请输入需要删除的任务ID：")
	task, err := TasksInfo[taskID]
	if err == false {
		fmt.Println("您输入的任务ID不存在.")
	} else {
		fmt.Println("ID：", taskID)
		fmt.Println("任务名：", task[name])
		fmt.Println("开始时间：", task[startTime])
		fmt.Println("结束时间：", task[endTime])
		fmt.Println("任务状态：", task[status])
		fmt.Println("负责人：", task[user])
	}

	if isDelete := Input("是否确认删除当前任务信息(y/yes):"); isDelete == "yes" || isDelete == "y" {
		delete(TasksInfo, taskID)
	}
	saveDB()
}

func ModifyTask() {
	taskID := Input("请输入需要编辑的任务ID：")
	task, err := TasksInfo[taskID]
	if err == false {
		fmt.Println("您输入的任务ID不存在.")
		return
	}
	fmt.Println("ID：", taskID)
	fmt.Println("任务名：", task[name])
	fmt.Println("开始时间：", task[startTime])
	fmt.Println("结束时间：", task[endTime])
	fmt.Println("任务状态：", task[status])
	fmt.Println("负责人：", task[user])

	if isDelete := Input("是否确认修改当前任务信息(yes/y):"); isDelete == "yes" || isDelete == "y" {
		taskName := Input("请输入任务名称：")
		taskStartTime := Input("开始时间：")
		taskStatus := Input("任务状态：")
		if taskName != task[name] && isExistName(taskName) {
			fmt.Println("存在同名任务，修改失败.")
			return
		} else if taskName != task[name] && isExistName(taskName) == false {
			task[name] = taskName
		}
		if _, err := time.ParseInLocation("2006-01-02 15:04:05", taskStartTime, time.Local); err == nil {
			task[startTime] = taskStartTime
		} else {
			fmt.Printf("输入的时间格式非法,%s", err)
			return
		}
		if taskStatus == statusStart || taskStatus == statusStop {
			task[status] = taskStatus
		} else if taskStatus == statusCompele {
			task[status] = taskStatus
			task[endTime] = time.Now().Format("2006-01-02 15:04:05")
		} else {
			fmt.Println("修改成非法的任务状态,修改失败.")
			return
		}
	}
	saveDB()
}
