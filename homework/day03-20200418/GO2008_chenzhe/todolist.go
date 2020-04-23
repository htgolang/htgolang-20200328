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
	startTime = "start_time"
	endTime   = "end_time"
	status    = "status"
	user      = "user"
)
var todos = []map[string]string{
	{"id": "1", "name": "陪孩子散步", "startTime": "18:00", "endTime": "", "status": statusNew, "user": "kk"},
	{"id": "2", "name": "备课", "startTime": "21:00", "endTime": "", "status": statusNew, "user": "kk"},
	{"id": "3", "name": "复习", "startTime": "09:00", "endTime": "", "status": statusNew, "user": "kk"},
}

const (
	statusNew     = "未执行"
	statusStart = "准备开始"
	statusDoding = "进行中"
	statusCompele = "完成"
)
//用户输入并且返回输入
func input(prompt string) string {
	var text string
	fmt.Println(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}
//获取todo新任务id
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
//显示一个任务
func show(dataIn map[string]string)  {
	fmt.Printf("ID：%s\n",dataIn[id])
	fmt.Printf("NAME：%s\n",dataIn[name])
	fmt.Printf("STARTTIME：%s\n",dataIn[startTime])
	fmt.Printf("ENDTIME：%s\n",dataIn[endTime])
	fmt.Printf("STATUS：%s\n",dataIn[status])
	fmt.Printf("USER：%s\n",dataIn[user])

}
//实例化一个新的任务
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
//新增一个任务
func add() {
	task := newTask()

	fmt.Println("请输入任务信息:")
	for{
		taskName := input("任务名:")
		if !checkTaskName(taskName){
			task[name]=taskName
			break
		}else {
			fmt.Println("任务已重名，请重新输入任务名")
		}
	}


	task[startTime] = time.Now().Format("2006-01-02 15:04:05")
	task[status] = statusNew
	task[user] = input("负责人:")

	todos = append(todos, task)
	fmt.Println("创建任务成功")
}
//编辑一个任务
func edit()  {
	id := input("请输入要编辑任务的id")
	for i,j:=range todos{
		if id == j["id"]{
			show(j)
			fmt.Println("请输入修改信息:")
			for{
				taskName := input("任务名:")
				if !checkTaskName(taskName){
					j[name]=taskName
					break
				}else {
					fmt.Println("任务已重名，请重新输入任务名")
				}
			}
			j[user] = input("负责人:")
			{
				for {
					fmt.Printf("任务状态必须是一下内容之一(未执行/准备开始/进行中/完成)")
					userInStatus := input("任务状态:")
					if userInStatus == statusNew || userInStatus == statusStart || userInStatus == statusDoding || userInStatus == statusCompele {
						if userInStatus == statusCompele {
							j[status] = userInStatus
							j[endTime] = time.Now().Format("2006-01-02 15:04:05")

						}else {
							j[status] = userInStatus
						}
						break
					}else {
						fmt.Printf("重新输入任务状态")
					}
				}
			}
			break
		}
	if i == len(todos)-1{
		fmt.Println("没有这个id")
	}
	}
}
//删除一个任务
func del()  {
	id := input("请输入要删除任务的id")
	for i,j :=range todos{
		if id==j["id"] {
			show(j)
			userInPut := input("以上是任务信息，确认要删除吗？(yes/y)")
			switch userInPut {
			case "yes","y":
				if i==len(todos)-1{
					todos=todos[0:i-1]
				}else {
					copy(todos[i:],todos[i+1:])
					todos=todos[0:len(todos)-1]
				}
			default:

			}
			break
		}
		if i == len(todos)-1{
			fmt.Println("没有这个id")
	}
	}

}
//搜索
func search()  {
	id := input("请输入你要查询的id号或all")
	if id=="all"{
		for _,j:=range todos{
			fmt.Println(strings.Repeat("-",20))
			show(j)
			fmt.Println(strings.Repeat("-",20))
		}
	}else {
		for i,j:=range todos{
			if id==j["id"]{
				show(j)
				break
			}
			if i==len(todos)-1{
				fmt.Println("没有这个id")
			}
		}
	}

}
//退出
func exit()  {
	os.Exit(0)
}
//验证任务名
func checkTaskName(name string) bool  {
	for i,j :=range todos{
		if j["name"]==name{
			return true
		}
		if i==len(todos)-1 {
			return false
		}
		}
	return false
	}



func main() {
	methods := map[string]func(){
		"add":add,
		"edit":edit,
		"del":del,
		"search":search,
		"exit":exit,
	}
	for {
		userinput := input("请输入你的操作(add/edit/del/search/exit)")
		if method,ok:=methods[userinput];ok{
			method()
		}else {
			fmt.Println("没有该指令，请重新输入")
		}
	}


}