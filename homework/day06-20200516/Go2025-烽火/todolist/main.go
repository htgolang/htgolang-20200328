package main

import (
	"fmt"
	"os"
	"todolist/task"
	"todolist/user"
	"todolist/utils"
)

// 命令行的任务管理器v4版本

func main() {
	if len(user.Accounts) > 0 {
		account := utils.Input("请输入用户名:")
		if user.VerifyPassword(account, user.PasswordLimit) {
			fmt.Println("密码验证成功,请继续后续操作!")
		} else {
			fmt.Println("3次密码验证错误，程序退出!")
			os.Exit(1)
		}
	} else {
		fmt.Println("请注册账号密码!")
		user.AddUser()
	}

	for {
		text := utils.Input("请输入操作[1.添加任务 2.查询任务 3.修改任务 4.删除 5.退出]:")
		if text == "5" {
			break
		}

		switch text {
		//add
		case "1":
			task.AddTask()
			record()
		//query
		case "2":
			task.QueryTaskWithSort()
		//modify
		case "3":
			if task.ModifyTask() {
				record()
			}
		//delete
		case "4":
			if len(task.Todolist) == 0 {
				fmt.Println("目前没有任何任务记录，请先添加任务，谢谢!")
				continue
			}
			if task.DeleteTask() {
				record()
			} else {
				fmt.Println("任务不存在")
			}
		// exit
		case "5":
			break
		default:
			fmt.Println("输入的指令错误")
		}
	}
}

func record() {
	method := utils.Input("请选择持久化任务方式:[1.txt方式 2.gob方式 3.csv方式 4.json方式]")
	switch method {
	case "1":
		task.JsonTask()
		task.TxtTask()
	case "2":
		task.JsonTask()
		task.GobTask()
	case "3":
		task.JsonTask()
		task.CsvTask()
	case "4":
		task.JsonTask()
	default:
		fmt.Println("不支持此种持久化方式!")
	}

}
