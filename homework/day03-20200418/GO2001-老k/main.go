package main

import (
	. "GO2001-老k/todolist"
	"fmt"
	"os"
)

var (
	todomgr = &TodoMgr{}
)

func showMenu() {
	fmt.Println("欢迎来到todolist系统")
	fmt.Println("1. 添加任务")
	fmt.Println("2. 编辑任务")
	fmt.Println("3. 展示所有任务")
	fmt.Println("4. 删除任务")
	fmt.Println("5. 退出系统")

}

func inputTodolist() *Todo {
	var (
		id        string
		name      string
		startTime string
		endTime   string
		status    string
		user      string
	)
	fmt.Println("请输入id：")
	fmt.Scan(&id)
	fmt.Println("请输入name：")
	fmt.Scan(&name)
	fmt.Println("请输入startTime：")
	fmt.Scan(&startTime)
	fmt.Println("请输入endTime：")
	fmt.Scan(&endTime)
	fmt.Println("请输入status：")
	fmt.Scan(&status)
	fmt.Println("请输入user：")
	fmt.Scan(&user)
	todo := NewTodo(id, name, startTime, endTime, status, user)
	return todo

}

func main() {
	showMenu()
	for {
		var option int
		fmt.Println("请输入：")
		fmt.Scanf("%d", &option)
		switch option {
		case 1:
			todo := inputTodolist()
			todomgr.AddTodo(todo)
		case 2:
			todo := inputTodolist()
			todomgr.EditTodolist(todo)
		case 3:
			todomgr.ShowTodolist()
		case 4:
			todo := inputTodolist()
			todomgr.DeleteTodolist(todo)
		case 5:
			os.Exit(-1)
		}
	}
}
