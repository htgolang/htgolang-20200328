package main

import (
	"fmt"
	"strconv"
)

var todos = []map[string]string{
	{"id": "1", "name": "陪孩子散步", "startTime": "18:00", "endTime": "", "status": statusNew, "user": "kk"},
	{"id": "2", "name": "备课", "startTime": "21:00", "endTime": "", "status": statusNew, "user": "kk"},
	{"id": "4", "name": "复习", "startTime": "09:00", "endTime": "", "status": statusNew, "user": "kk"},
}

const (
	id        = "id"
	name      = "name"
	startTime = "start_time"
	endTime   = "end_time"
	status    = "status"
	user      = "user"
)

const (
	statusNew     = "未执行"
	statusCompele = "完成"
)

func genId() int {
	var rt int
	for _, todo := range todos {
		todoId, _ := strconv.Atoi(todo["id"])
		if rt < todoId {
			rt = todoId
		}
	}
	return rt + 1
}

func newTask() map[string]string {
	// id生成(用todos中最大的ID+1)
	task := make(map[string]string)
	task[id] = strconv.Itoa(genId())
	task[name] = ""
	task[startTime] = ""
	task[endTime] = ""
	task[status] = statusNew
	task[user] = ""
	return task
}

func main() {
	var text string

	task := newTask()

	fmt.Println("请输入任务信息:")
	fmt.Print("任务名: ")
	fmt.Scan(&text)
	task[name] = text

	fmt.Print("开始时间: ")
	fmt.Scan(&text)
	task[startTime] = text

	fmt.Print("负责人:")
	fmt.Scan(&text)
	task[user] = text

	todos = append(todos, task)
	fmt.Println(todos)
}
