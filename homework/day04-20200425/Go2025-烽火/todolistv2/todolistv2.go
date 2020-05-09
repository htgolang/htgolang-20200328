package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
	"github.com/olekukonko/tablewriter"
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
	password  = "b0230d6ac1b3b3bcff28ace36d15ad5d" // hello
	salt      = "02P8bA"
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

// 获取用户输入信息
func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}

// 生成最大的id
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

// 新任务初始化设置
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

//渲染输出任务信息
func renderTask(tasks []map[string]string) {
	header := []string{"ID", "name", startTime, endTime, status, user}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetHeader(header)
	for i := 0; i < len(header); i++ {
		table.SetColMinWidth(i, 20)
	}
	for _, task := range tasks {
		table.Append([]string{task["id"], task[name], task[startTime], task[endTime], task[status], task[user]})
	}
	table.Render()
	// table.AppendBulk(content)
}

// 添加一个任务
func addTask() {
	task := newTask()
	fmt.Println("请输入任务信息:")

	for {
		tempName := input("任务名:")
		if verifyName(tempName) {
			task[name] = tempName
			break
		} else {
			fmt.Println("任务名称已存在!")
		}
	}
	task[startTime] = input("开始时间:")
	task[user] = input("负责人:")
	todos = append(todos, task)
	content := []map[string]string{task}
	renderTask(content)
}

// 排序任务信息
func sortTask(tasks []map[string]string, key string) []map[string]string {
	if key == "name" || key == "startTime" {
		sort.Slice(tasks, func(i, j int) bool { return tasks[i][key] < tasks[j][key] })
	}
	return tasks
}

// 查询任务
func queryTaskWithSort() {
	var queryMap = map[string]string{"1": "name", "2": "startTime"}
	q := input("请输入查询信息:")
	content := make([]map[string]string, 0)
	for _, task := range todos {
		if q == "all" || strings.Contains(task[name], q) {
			content = append(content, task)
		}
	}
	if len(content) == 0 {
		fmt.Println("未找到关联任务!")
	} else {
		key := input("请输入排序方式[1.任务名称 2.任务开始时间]:")
		newTasks := sortTask(content, queryMap[key])
		renderTask(newTasks)
	}
}

// 修改任务
func modifyTask() {
	q := input("请输入需要修改的任务ID:")
	for _, task := range todos {
		if q == task["id"] {
			content := []map[string]string{task}
			renderTask(content)
			switch input("是否确认修改(y/yes):") {
			case "y", "yes":
				for {
					tempName := input("任务名称:")
					if verifyName(tempName) {
						task[name] = tempName
						break
					} else {
						fmt.Println("任务名称已存在!")
					}
				}
				task[startTime] = input("开始时间:")

				for {
					tempStaus := input("状态:")
					if verifyStatus(tempStaus) {
						if tempStaus == statusComplete {
							task[endTime] = time.Now().Format("2006-01-02 15:04:05")
						}
						task[status] = tempStaus
						break
					} else {
						fmt.Println("输入的状态值不对!可选范围:", strings.Join(statusChoice, ", "))

					}
				}
				fmt.Println("任务修改完成")
				renderTask(content)
			default:
				fmt.Println("取消修改")
				break
			}
		}
	}
}

// 删除任务
func deleteTask() {
	queryId := input("请输入需要删除的任务ID:")
	for index, task := range todos {
		if queryId == task["id"] {
			content := []map[string]string{task}
			renderTask(content)
			switch input("是否进行删除(y/yes):") {
			case "y", "yes":
				copy(todos[index:], todos[index+1:])
				// newTasks := todos[:len(todos)-1]
				fmt.Printf("任务ID:%s 已删除", queryId)
			default:
				fmt.Println("取消删除")
			}
		}
	}
}

// 验证任务名，确保唯一性
func verifyName(inputName string) bool {
	for _, task := range todos {
		if inputName == task[name] {
			return false
		}
	}
	return true
}

// 验证状态值在可选列表内
func verifyStatus(inputStatus string) bool {
	for _, status := range statusChoice {
		if inputStatus == status {
			return true
		}
	}
	return false
}

// 加盐MD5
func saltMd5(input string) string {
	hasher := md5.New()
	io.WriteString(hasher, input)
	io.WriteString(hasher, salt)

	cryptoPasswd := fmt.Sprintf("%x", hasher.Sum(nil))
	return cryptoPasswd
}

// 非明文显示终端输入信息，获取信息
func getPassword() string {
	fmt.Print("请输入密码: ")
	passwd, _ := gopass.GetPasswdMasked()
	return string(passwd)
}

// 验证密码
func verifyPassword() bool {
	limit := 3
	for count := 0; count < limit; count++ {
		input := getPassword()
		if saltMd5(input) == password {
			return true
		} else {
			fmt.Printf("密码验证错误，还剩%d次机会!\n", limit-count-1)
		}
	}
	return false
}

func main() {
	methods := map[string]func(){
		"1": addTask,
		"2": queryTaskWithSort,
		"3": modifyTask,
		"4": deleteTask,
	}

	if !verifyPassword() {
		fmt.Println("3次密码验证错误，程序退出")
		os.Exit(1)
	}

	for {
		text := input("请输入操作[1.添加任务 2.查询任务 3.修改任务 4.删除 5.退出]:")
		if text == "5" {
			break
		}
		if method, ok := methods[text]; ok {
			method()
		} else {
			fmt.Println("输入的指令错误")
		}
	}

}
