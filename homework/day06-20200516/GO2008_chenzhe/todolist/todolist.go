package main

import (
	"fmt"
	"todolist/until/auth"
	"todolist/until/task"
)


func main() {

	auth.Auth()
	task.UserDataType()
	task.LoadTask()
	methods := map[string]func(){
		"add":task.Add,
		"edit":task.Edit,
		"del":task.Del,
		"search":task.Search,
		"exit":task.Exit,
		"showall":task.Showall,
	}
	for {
		userinput := auth.Input("请输入你的操作(add/edit/del/search/exit/showall)")
		if method,ok:=methods[userinput];ok{
			method()
		}else {
			fmt.Println("没有该指令，请重新输入")
		}
	}


}