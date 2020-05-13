package main

import (
	"fmt"
	"os"
	"todolistv3/task"
	"todolistv3/utils"
)

// 命令行的任务管理器
var (
	passwordFile  = "password.txt"
	passwordLimit = 3
	taskFile      = "tasks.txt"
)

func main() {
	todolist := make([]task.Task, 0)
	if utils.FileIsExists(taskFile) {
		todolist = task.ReadTaskFromFile(taskFile)
	}

	if utils.FileIsExists(passwordFile) {
		password := utils.ReadFile(passwordFile)
		if utils.VerifyPassword(passwordFile, password, passwordLimit) {
			fmt.Println("密码验证成功,请继续后续操作!")
		} else {
			fmt.Println("3次密码验证错误，程序退出!")
			os.Exit(1)
		}
	} else {
		utils.SetPassword(passwordFile, 0)
	}

	for {
		text := utils.Input("请输入操作[1.添加任务 2.查询任务 3.修改任务 4.删除 5.退出]:")
		if text == "5" {
			break
		}

		switch text {
		//add
		case "1":
			newTask := task.AddTask(todolist, passwordFile)
			todolist = append(todolist, newTask)
			task.RecordTask(taskFile, todolist...)
		//query
		case "2":
			task.QueryTaskWithSort(todolist)
		//modify
		case "3":
			task.ModifyTask(todolist, passwordFile)
			task.RecordTask(taskFile, todolist...)
		//delete
		case "4":
			if len(todolist) == 0 {
				fmt.Println("目前没有任何任务记录，请先添加任务，谢谢!")
				continue
			}
			todolist = task.DeleteTask(todolist, passwordFile)
			task.RecordTask(taskFile, todolist...)
		// exit
		case "5":
			break
		default:
			fmt.Println("输入的指令错误")
		}
	}
}
