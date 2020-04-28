package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 命令行的任务管理器

// 1. 任务的输入(添加任务)
// 2. 任务列表(任务查询)
// 3. 任务修改
// 4. 任务删除
// 5. 详情

// 任务
// ID, 任务名称， 开始时间， 结束时间，状态， 负责人
// ID, name.,start_time, end_time, status, user
// []map[string][string]

const (
	name      = "name"
	startTime = "startTime"
	endTime   = "endTime"
	status    = "status"
	user      = "user"
)

const (
	statusNew      = "未执行"
	statusComplete = "完成"
	statusBegin    = "开始执行"
	statusPause    = "暂停"
)

var statusChoice []string = []string{statusNew, statusBegin, statusComplete, statusPause}

// var todos = make([]map[string]string, 0)
var todos = []map[string]string{
	{"id": "1", "name": "study", "startTime": "17:00", "endTime": "19:00", "status": statusNew, "user": "devon"},
	{"id": "2", "name": "sleep", "startTime": "19:00", "endTime": "20:00", "status": statusNew, "user": "devon"},
	{"id": "3", "name": "eat", "startTime": "18:00", "endTime": "19:00", "status": statusNew, "user": "devon"},
}
var taskFormat = fmt.Sprintf("|%3s|%20s|%40s|%40s|%40s|%20s|", "ID", "name", startTime, endTime, status, user)

func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}

func genId() int {
	// 生成最大的id
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
	task := make(map[string]string)
	task["id"] = strconv.Itoa(genId())
	task[name] = ""
	task[startTime] = ""
	task[endTime] = ""
	task[status] = statusNew
	task[user] = ""
	return task
}

func printTask(task map[string]string) {
	fmt.Println(strings.Repeat("-", len(taskFormat)))
	fmt.Printf("|%3s|%20s|%40s|%40s|%40s|%20s|\n", task["id"], task[name], task[startTime], task[endTime], task[status], task[user])

}

func add() {
	task := newTask()
	fmt.Println("请输入任务信息:")

	for {
		tempName := input("任务名:")
		if verify_name(tempName) {
			task[name] = tempName
			break
		} else {
			fmt.Println("任务名称已存在!")
		}
	}
	task[startTime] = input("开始时间:")
	task[user] = input("负责人:")
	todos = append(todos, task)
}

func query() {
	q := input("请输入查询信息:")

	fmt.Println(taskFormat)
	for _, todo := range todos {
		if q == "all" || strings.Contains(todo[name], q) {
			printTask(todo)
		}
	}
}

func modify() {
	q := input("请输入需要修改的任务ID:")
	for _, task := range todos {
		if q == task["id"] {
			printTask(task)
			switch input("是否确认修改(y/yes):") {
			case "y", "yes":
				for {
					tempName := input("任务名称:")
					if verify_name(tempName) {
						task[name] = tempName
						break
					} else {
						fmt.Println("任务名称已存在!")
					}
				}
				task[startTime] = input("开始时间:")

				for {
					tempStaus := input("状态:")
					if verify_status(tempStaus) {
						if tempStaus == statusComplete {
							task[endTime] = time.Now().Format("2006-01-02 15:04:05")
						}
						task[status] = tempStaus
						break
					} else {
						fmt.Println("输入的状态值不对!可选范围:", strings.Join(statusChoice, ", "))

					}
				}
				printTask(task)
			default:
				break
			}
		}
	}
}

func remove() {
	q := input("请输入需要删除的任务ID:")
	for index, task := range todos {
		if q == task["id"] {
			printTask(task)
			switch input("是否进行删除(y/yes):") {
			case "y", "yes":
				copy(todos[index:], todos[index+1:])
				newTasks := todos[:len(todos)-1]
				for _, task := range newTasks {
					printTask(task)
				}
			}
		}
	}
}

func verify_name(inputName string) bool {
	for _, task := range todos {
		if inputName == task[name] {
			return false
		}
	}
	return true
}

func verify_status(inputStatus string) bool {
	for _, status := range statusChoice {
		if inputStatus == status {
			return true
		}
	}
	return false
}

func main() {
	methods := map[string]func(){
		"add":    add,
		"query":  query,
		"modify": modify,
		"delete": remove,
	}

	for {
		text := input("请输入操作[add/query/modify/delete/exit]:")
		if text == "exit" {
			break
		}

		if method, ok := methods[text]; ok {
			method()
		} else {
			fmt.Println("输入的指令错误")
		}
	}

}
