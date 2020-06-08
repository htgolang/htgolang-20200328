package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 做一个命令行的任务管理
// 用户管理

// 1. 函数, 输入&输出, 复合数据结构, 基本数据类型
// 2. 了解流程（对数据的操作流程，增、删、改、查）

// 1. 任务的输入(添加任务)
// 2. 任务列表(任务查询)
// 3. 任务修改
// 4. 任务删除
// 5. 详情

// 任务
// ID, 任务名称, 开始时间，结束时间, 状态, 负责人
// ID, name, start_time, end_time, status, user
// []map[string][string]
var todos = []map[string]string{
	{"id": "1", "name": "陪孩子散步", "startTime": "18:00", "endTime": "", "status": statusNew, "user": "Jordan"},
	{"id": "2", "name": "备课", "startTime": "21:00", "endTime": "", "status": statusNew, "user": "Jordan"},
	{"id": "4", "name": "复习", "startTime": "09:00", "endTime": "", "status": statusNew, "user": "Jordan"},
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
	statusNew      = "未执行"
	statusComplete = "已完成"
)

func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}
func getId() int {
	var rt int
	for _, gid := range todos {
		toDoId, _ := strconv.Atoi(gid["id"])
		if rt < toDoId {
			rt = toDoId
		}
	}
	return rt + 1
}

func newTask() map[string]string {
	// id 自动生成,最大的ID+啊
	task := make(map[string]string)
	task[id] = strconv.Itoa(getId())
	task[name] = ""
	task[startTime] = ""
	task[endTime] = ""
	task[status] = statusNew
	task[user] = ""
	return task
}
func printTask(task map[string]string) {
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("ID: ", task[id])
	fmt.Println("任务名: ", task[name])
	fmt.Println("开始时间: ", task[startTime])
	fmt.Println("完成时间: ", task[endTime])
	fmt.Println("负责人: ", task[user])
	fmt.Println(strings.Repeat("-", 20))
}

func add() {
	task := newTask()
	fmt.Println("请输入你的任务信息: ")
	task[name] = input("任务名: ")
	task[startTime] = input("开始时间: ")
	task[user] = input("负责人: ")
	todos = append(todos, task)
	fmt.Println("任务创建成功!")
}

func query() {
	q := input("请输入查询信息：")
	for _, todo := range todos {
		if q == "all" || strings.Contains(todo[name], q) == true {
			printTask(todo)
		}
	}
}

func modify() {
	uid := input("输入要修改用户的ID: ")
	var todo map[string]string
	for _, task := range todos {
		if task[id] == uid {
			todo = task
		}
	}
	if todo != nil {
		fmt.Println("要修改的用户信息: ")
		printTask(todo)
		inputInfo := input("确定修改（Y/y）?")
		if inputInfo == "Y" || inputInfo == "y" {
			todo[name] = input("任务名: ")
			todo[startTime] = input("开始时间: ")
			todo[user] = input("负责人: ")
			fmt.Println("修改成功！")
		} else {
			fmt.Println("指令不正确，退出修改！")
		}
	} else {
		fmt.Println("ID不存在！")
	}

}
func remove() {
	uid := input("输入要删除的ID: ")
	var index int
	var todo map[string]string
	for taskIndex, task := range todos {
		if task[id] == uid {
			todo = task
			index = taskIndex
		}
	}
	if todo != nil {
		fmt.Println("要删除的任务信息: ")
		printTask(todo)
		inputInfo := input("确定是否删除 (Y/y): ")
		if inputInfo == "Y" || inputInfo == "y" {
			todos = append(todos[:index], todos[index+1:]...)
		} else {
			fmt.Println("输入指令不正确，退出删除！")
		}
	} else {
		fmt.Println("ID不存在！")
	}

}

func main() {
	methods := map[string]func(){
		"add":    add,
		"query":  query,
		"modify": modify,
		"delete": remove,
	}
	for {
		text := input("请输入操作：add/query/modify/delete/exit...\n")
		if text == "exit" {
			break
		}
		if method, ok := methods[text]; ok {
			method()
		} else {
			fmt.Println("输入的指令不正确!")
		}
	}
}
