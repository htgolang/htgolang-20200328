package main

import (
	"fmt"
	"todolist/checkPass"
	"todolist/todo"
)
func main() {
	result := checkPass.CheckPass()
	if result {
		fmt.Println("登陆系统成功")
		todo.Todolist()
	}
}
