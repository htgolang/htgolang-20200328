package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"github.com/olekukonko/tablewriter"
	"github.com/howeyc/gopass"
)

//TODO LIST a. 认证 密码 =》 内置在程序中 密码 => md5 + salt
//打开程序
//认证: 密码
//password + salt => hashed => 程序内置的hash值比较
//
//认证成功可以进行后续操作，如果连续输入密码超过3次失败，直接退出程序
//
//输入密码: fmt.Scan()
//
//第三方包: gopass


var todos = []map[string]string{
	{"id":"1","name":"读书","startTime":"18:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"2","name":"复习","startTime":"20:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"3","name":"发呆","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"4","name":"练功","startTime":"11:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"5","name":"上山","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"6","name":"买东西","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"7","name":"吃饭","startTime":"21:00","endTime":"","status":statusNew,"user":"chen"},
	{"id":"8","name":"买课本","startTime":"18:00","endTime":"","status":statusNew,"user":"chen"},
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
	tasks := make([]map[string]string,0)
	text = input("请输入查询信息(all 查询所有任务)：")
	for _,todo :=range todos{
		if text =="all" || strings.Contains(todo[name],text){
			//printTask(todo)
			tasks = append(tasks,todo)

		}
	}
	new_tasks := sort_query(tasks)
	tasks_table(new_tasks)
	//fmt.Println(new_tasks)
}

func sort_query(tasks []map[string]string) []map[string]string{
	sort.Slice(tasks, func(i, j int) bool {    // 排序时间升须
		t1,_ := time.Parse("15:04",tasks[i][startTime])
		t2,_:= time.Parse("15:04",tasks[j][startTime])
		return t1.Before(t2)
	})
	return tasks
}
func tasks_table(tasks []map[string]string)  {  //用第三方库tablewriter
	var ss = [][]string{}
	for _,todo := range tasks{
		ss = append(ss,[]string{todo[id],todo[name],todo[startTime],todo[endTime],todo[status],todo[user]})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID","Name","StartTime","EndTime","Status","User"})
	for _,v := range ss{
		table.Append(v)
	}
	table.Render()
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
func check_name(input_name string) bool { //检查任务名字是否已经存在
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

func RandString() string  {
	strs := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var bytes []byte
	for i:=0;i<5;i++{
		bytes = append(bytes,strs[rand.Intn(len(strs))])
	}
	return string(bytes)
} //生成随机盐
func md5_salt(pw ,salt string) string {
	rs := pw + salt
	return fmt.Sprintf("%x",md5.Sum([]byte(rs)))
} //生成混合的密码和盐的16位字符串

func check_passwd()  bool {
	//var pass string
	real_pw := "abc"
	salt := RandString()
	md5salt := md5_salt(real_pw,salt)
	//pass = input("请输入密码：")
	fmt.Println("请输入密码:")
	pass,_ := gopass.GetPasswd()
	input_pass := md5_salt(string(pass),salt)
	if md5salt == input_pass {
		return true
	}
	return false
}  //检查密码
func main()  {
	methods := map[string]func(){
		"add":add,
		"query":query,
		"modify":modify,
		"delete":delete,
	}

	for {
		if check_passwd() {               //默认正确密码设置为 "abc"
			text := input("请输入操作(add/query/modify/delete):")
			if text == "exit" {
				return
			}

			if method, ok := methods[text]; ok {
				method()
			} else {
				fmt.Println("输入指令不正确！")
			}
		}else{
			fmt.Println("密码不正确")
		}
	}
}
