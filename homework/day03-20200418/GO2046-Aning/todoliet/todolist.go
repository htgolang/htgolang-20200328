package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	id        = "id"
	name      = "name"
	startTime = "starttime"
	endTime   = "endtime"
	status    = "status"
	user      = "user"
)
const (
	statusNew     = "未执行"
	statusStart   = "准备开始"
	statusDoding  = "进行中"
	statusCompele = "完成"
)

var todos = []map[string]string{
	{"id": "1", "name": "陪孩子散步", "startTime": "18:00", "endTime": "", "status": statusNew, "user": "kk"},
	{"id": "2", "name": "备课", "startTime": "21:00", "endTime": "", "status": statusNew, "user": "kk"},
	{"id": "3", "name": "复习", "startTime": "09:00", "endTime": "", "status": statusNew, "user": "kk"},
}

func getid() int {
	var big int
	for _, todo := range todos {
		todoid, _ := strconv.Atoi(todo["id"])
		if big < todoid {
			big = todoid
		}
	}
	return big + 1
}
func newtask() map[string]string {
	task := make(map[string]string)
	task["id"] = strconv.Itoa(getid())
	task[name] = ""
	task[startTime] = ""
	task[endTime] = ""
	task[status] = statusNew
	task[user] = ""
	return task
}
func printtask(task map[string]string) {
	fmt.Printf("ID：%s\n", task[id])
	fmt.Printf("NAME：%s\n", task[name])
	fmt.Printf("STARTTIME：%s\n", task[startTime])
	fmt.Printf("ENDTIME：%s\n", task[endTime])
	fmt.Printf("STATUS：%s\n", task[status])
	fmt.Printf("USER：%s\n", task[user])
}
func input(inputt string) string {
	var inputtt string
	fmt.Println(inputt)
	fmt.Scan(&inputtt)
	return strings.TrimSpace(inputtt)
}
func add() {
	task := newtask()
	fmt.Println("输入信息")
	for {
		tempname := input("任务名:")
		task[name] = tempname
	}
	task[startTime] = input("开始时间")
	task[user] = input("负责人")
	todos = append(todos, task)
}
func edit() {
	q := input("请输入需要修改的任务ID:")
	for _, task := range todos {
		if q == task["id"] {
			print(task)
			switch input("是否确认修改(y/yes):") {
			case "y", "yes":
				for {
					tempName := input("任务名称:")
					task[name] = tempName
				}
				task[startTime] = input("开始时间:")
				for {
					qq := input("状态:")
					task[endTime] = time.Now().Format("2006-01-02 15:04:05")
					task[status] = qq
				}
				print(task)
			default:
				break
			}
		}
	}
}

func del() {
	q := input("要删除的")
	for index, task := range todos {
		if q == task["id"] {
			print(task)
			switch input("是否进行删除(y/yes):") {
			case "y", "yes":
				copy(todos[index:], todos[index+1:])
				newTasks := todos[:len(todos)-1]
				for _, task := range newTasks {
					print(task)
				}
			}
		}
	}
}
func search() {
	input("输入查询")
	for _, todo := range todos {
		printtask(todo)
	}
}
func exit() {
	os.Exit(0)
}
func main() {
	op := map[string]func(){
		"add":    add,
		"edit":   edit,
		"del":    del,
		"search": search,
		"exit":   exit,
	}
	for {

		inputt := input("1.add    2.edit    3.del	   4.search   5.exit")
		fmt.Scan(&inputt)
		if opp, ok := op[inputt]; ok {
			opp()
		} else {
			fmt.Println("没有选项")
		}
	}
}
