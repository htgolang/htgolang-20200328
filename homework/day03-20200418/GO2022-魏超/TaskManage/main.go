package main

import (
	"fmt"

	"github.com/shadow_wei/TaskManage/task"
)

func main() {
	methods := map[string]func(){
		"add":    task.NewTask,
		"query":  task.QueryTask,
		"delete": task.DeleteTask,
		"modify": task.ModifyTask,
	}
	for {
		text := task.Input("请输入操作(add/query/delete/modify)：")
		if text == "exit" {
			break
		}
		method, ok := methods[text]
		if ok {
			method()
		} else {
			fmt.Println("您输入的操作是非法的.")
		}
	}
}
