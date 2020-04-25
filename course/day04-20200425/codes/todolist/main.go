package main

import (
	"fmt"
	mtask "todolist/models/task"
	"todolist/controllers/task"
	// 1. 导入 mod init name + 目录结构
	// 2. 包名与所在目录名称保持一致
	// 3. 调用时用包名.VAR(包名不是文件名)
	"github.com/imsilence/strutil"
)

func main() {
	fmt.Println(mtask.Name)
	fmt.Println(task.Name)
	task.Call()
	fmt.Println(task.Version)
	fmt.Println(strutil.RandString())
}