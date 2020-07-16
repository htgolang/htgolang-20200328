package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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

func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}

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

func printTask(task map[string]string) {
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("ID:", task[id])
	fmt.Println("任务名:", task[name])
	fmt.Println("开始时间:", task[startTime])
	fmt.Println("完成时间:", task[endTime])
	fmt.Println(strings.Repeat("-", 20))
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

func getIndex(id string) int {
	for index, i := range todos {
		if i["id"] == id {
			return index
		}
	}
	return 0
}

func checkName(old string) bool {
	for _, i := range todos {
		if i["name"] == old {
			return false
		}
	}
	return true
}

func add() {
	task := newTask()

	name_str := "任务名:"
	fmt.Println("请输入任务信息:")

	for {
		task[name] = input(name_str)
		nameStatus := checkName(name)
		if nameStatus {
			break
		} else {
			name_str = "用户名已存在，请重新上输入:"
		}
	}
	task[startTime] = input("开始时间:")
	task[user] = input("负责人:")

	todos = append(todos, task)
	fmt.Println("创建任务成功")
}

func query() {
	// all 显示全部
	q := input("请输入查询信息:")

	for _, todo := range todos {
		if q == "all" || strings.Contains(todo[name], q) {
			printTask(todo)
		}
	}
}

func modify() {
	// 编辑
	for _, i := range todos {
		fmt.Printf("ID: %s, NAME: %s\n", i["id"], i["name"])
	}
	id_str := input("请输入需要编辑的ID信息:")
	status := input(fmt.Sprintf("您选择ID:%s, 是否需要编辑(y/yes)", id_str))
	if status == "y" || status == "yes" || status == "Y" || status == "YES" {
		fmt.Println(status)
		name_str := "请输入用户名:"
		name := input(name_str)
		for {
			nameStatus := checkName(name)
			if nameStatus {
				break
			} else {
				name_str = "用户名已存在，请重新上输入:"
				name = input(name_str)
			}
		}
		status_time := input("请输入开始时间:")
		inputStatus := input("请输入状态:")

		fmt.Println(name, status_time, inputStatus)
		for _, task := range todos {
			if task["id"] == id_str {
				// 进行更新数据
				task["name"] = name
				task["start_time"] = status_time
				task["status"] = inputStatus
				if inputStatus == statusCompele {
					task["end_time"] = time.Now().Format("2001-01-01 15:01:01")
				}
			}
		}

	}
}

func del() {
	// 删除
	for _, i := range todos {
		fmt.Printf("ID: %s, NAME: %s\n", i["id"], i["name"])
	}
	idStr := input("请输入需要删除的ID信息:")
	status := input(fmt.Sprintf("您选择ID:%s, 是否需要删除(y/yes)", idStr))
	if status == "y" || status == "yes" || status == "Y" || status == "YES" {
		for index, task := range todos {
			if task["id"] == idStr {
				fmt.Println(todos[index])
				fmt.Printf("%T\n", task)
			}
			fmt.Println("task->", task[id], idStr)
		}

		index := getIndex(idStr)
		if index == 0 {
			fmt.Println("您输入的id不存在！")
		}
		start_todos := todos[:index]
		end_todos := todos[index+1:]
		todos = start_todos
		todos = append(todos, end_todos...)
		for _, todo := range todos {
			printTask(todo)

		}

	} else {
		return
	}

}

func main() {
	methods := map[string]func(){
		"add":    add,
		"query":  query,
		"modify": modify,
		"delete": del,
	}
	for {
		text := input("请输入操作(add/query/modify/delete/exit/):")
		if text == "exit" {
			break
		}

		if method, ok := methods[text]; ok {
			method()
		} else {
			fmt.Println("输入指令不正确")
		}
	}
}
