package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//1. 任务的输入（添加任务）
//2. 任务列表（任务查询）
//3.任务修改
//4.任务删除
//5. 详情

//任务名称，开始时间，结束时间，状态，负责人
// id,name,startTime,endTime,status,user
// []map[string]string
var todos = []map[string]string{
	{"id":"1","name":"读书","startTime":"18:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"2","name":"复习","startTime":"20:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"3","name":"发呆","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"4","name":"练功","startTime":"11:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"5","name":"上山","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"6","name":"买东西","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"7","name":"吃饭","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"8","name":"买课本","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
}
const  (
	id = "id"
	name = "name"
	startTime = "startTime"
	endTime = "endTime"
	status = "status"
	user = "user"
)
const (
	statusNew = "新创建"
	statusComplete = "完成"
	statusIncomplete = "未完成"

)

func genId() int {
	var rt int
	for _,todo :=range todos{
		todoid,_ := strconv.Atoi(todo["id"])
		if  todoid > rt{
			rt = todoid
		}
	}
	return rt + 1
}
func printTask(task map[string]string)  {
	fmt.Println(strings.Repeat("-",20))
	fmt.Println("id:",task[id])
	fmt.Println("任务名称:",task[name])
	fmt.Println("开始时间:",task[startTime])
	fmt.Println("结束时间:",task[endTime])
	fmt.Println("状态:",task[status])
	fmt.Println("负责人:",task[user])
	fmt.Println(strings.Repeat("-",20))
}
func newTask() map[string]string  {
	task := make(map[string]string)
	task[id] = strconv.Itoa(genId())
	task[name] = ""
	task[startTime] = ""
	task[endTime] = ""
	task[status] = statusNew
	task[user] = ""
	return task
}
func add()  {
	task := newTask()
	fmt.Println("请输入任务信息：")

	task[name] = input("任务名")
	task[startTime] = input("开始时间")
	task[user] = input("负责人")
	todos = append(todos,task)
	fmt.Println("创建任务成功！")
	//fmt.Println(todos)
}
func query()  {
	var text string
	text = input("请输入查询信息(all 查询所有任务)：")
	for _,todo :=range todos{
		if text =="all" || strings.Contains(todo[name],text){
			printTask(todo)
		}
	}
}
func modify()  {
	var id_text,text string
	var flag bool
	id_text = input("请输入修改的任务的ID:")
	for _,todo := range todos{
		if todo[id] == id_text {
			flag = true
			text = input("是否要编辑任务?(y/n)")
			if text == "n"{
				return
			} else if text == "y"{
				name_text := input("请修改任务名字：")
				if check_name(name_text){
					fmt.Println("任务名已存在")
					break
				} else {
					todo[name] = name_text
				}
				todo[startTime] = input("请修改任务开始时间：")
				todo[status] = input("请修改任务的状态")
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
		fmt.Println("该ID不存在")
		return
	}
}

func delete()  {
	var id_text string
	var flag bool
	id_text = input("请输入删除任务的ID")
	for idx,todo := range todos{
		if todo[id] == id_text{
			flag = true
			yes_no := input("是否真的要删除该任务?(y/n)")
			if yes_no == "y"{
				copy(todos[idx:],todos[idx+1:])
				todos = todos[:len(todos)-1]
				//todos = append(todos[:idx],todos[idx+1:]...)
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
		fmt.Println("无效ID不存在")
		return
	}
}
func check_name(input_name string) bool {
	var sig bool
	for _,todo := range todos{
		if input_name == todo[name] {
			sig = true
			break
		}
	}
	return sig
}

func input(prompt string) string  {
	var text string
	fmt.Println(prompt)
	fmt.Scan(&text)
	return text
}

func main()  {
	methods := map[string]func(){
		"add":add,
		"query":query,
		"modify":modify,
		"delete":delete,
	}
	for {
		text := input("请输入操作(add/query/modify/delete):")
		if text == "exit"{
			return
		}

		if method ,ok := methods[text];ok {
			method()
		} else {
			fmt.Println("输入指令不正确！")
		}
	}
}
