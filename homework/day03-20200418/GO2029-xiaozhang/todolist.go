package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
做一个命令行的任务管理
用户管理

1. 函数, 输入&输出, 复合数据结构, 基本数据类型
2. 了解流程（对数据的操作流程，增、删、改、查）

1. 任务的输入(添加任务)
2. 任务列表(任务查询)
3. 任务修改
4. 任务删除
5. 详情

任务
ID, 任务名称, 开始时间，结束时间, 状态, 负责人
ID, name, start_time, end_time, status, user
[]map[string][string]
*/

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
	statusNew        = "未执行"
	statusComplete   = "完成"
	statusIncomplete = "未完成"
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
		if todoId > rt {
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
	fmt.Println("状态:", task[status])
	fmt.Println("负责人:", task[user])
	fmt.Println(strings.Repeat("-", 20))
}

func newTask() map[string]string {
	// id 生成（用todos中最大的ID+1）
	task := make(map[string]string)
	task[id] = strconv.Itoa(genId())
	task[name] = ""
	task[startTime] = ""
	task[endTime] = ""
	task[status] = statusNew
	task[user] = ""
	return task
}

func add() {
	task := newTask()

	fmt.Println("请输入任务信息:")
	task[name] = input("任务名:")
	task[startTime] = input("开始时间:")
	task[user] = input("负责人:")

	todos = append(todos, task)
	fmt.Println("任务创建成功！")
}

func query() {
	// 显示全部
	q := input("请输入查询信息:")

	for _, todo := range todos {
		if q == "all" || strings.Contains(todo[name], q) {
			printTask(todo)
		}
	}
}

func modify() {
	var id_text, text string
	var flag bool
	id_text = input("请输入修改的任务的ID:")
	for _, todo := range todos {
		if todo[id] == id_text {
			flag = true
			text = input("是否要编辑任务?(y/n)")
			if text == "n" {
				return
			} else if text == "y" || text == "Y" {
				name_text := input("请修改任务名字:")
				if check_name(name_text) {
					fmt.Println("任务名已存在")
					break
				} else {
					todo[name] = name_text
				}
				todo[startTime] = input("请修改任务开始时间:")
				todo[status] = input("请修改任务的状态:")
				if todo[status] == "完成" {
					todo[endTime] = time.Now().Format("2006-01-02 15:04:05")
				}
				fmt.Println("修改成功")
				fmt.Println(todo)
			} else {
				fmt.Println("无效的输入")
				return
			}

		}
	}
	if !flag {
		fmt.Println("该ID不存在！")
		return
	}
}

func delete() {
	var id_text string
	var flag bool
	id_text = input("请输入删除任务的ID:")
	for idx, todo := range todos {
		if todo[id] == id_text {
			flag = true
			yes_no := input("是否真的要删除该任务?(y/n)")
			if yes_no == "y" || yes_no == "Y" {
				copy(todos[idx:], todos[idx+1:])
				todos = todos[:len(todos)-1]
				// todos = append(todos[:idx], todos[idx+1]...)
				fmt.Println("删除成功")
				fmt.Println(todos)
			} else {
				fmt.Println("删除操作取消")
				return
			}
			break
		}
	}
	if !flag {
		fmt.Println("无效，id不存在！")
		return
	}
}
func check_name(input_name string) bool {
	var sig bool
	for _, todo := range todos {
		if input_name == todo[name] {
			sig = true
			break
		}
	}
	return sig
}

func main() {
	methods := map[string]func(){
		"add":    add,
		"query":  query,
		"modify": modify,
		"delete": delete,
	}
	for {
		text := input("请输入操作（add/query/delete/modify/exit...）")
		if text == "exit" {
			break
		}

		if method, ok := methods[text]; ok {
			method()
		} else {
			fmt.Println("输入指令不正确！")
		}
	}
}
