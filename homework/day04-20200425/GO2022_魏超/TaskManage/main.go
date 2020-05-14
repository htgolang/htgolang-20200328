package main

import (
	"fmt"
	"strconv"

	"github.com/shadow_wei/TaskManage/passwd"
	"github.com/shadow_wei/TaskManage/task"
	"github.com/shadow_wei/TaskManage/user"
	"github.com/shadow_wei/TaskManage/utils"
)

func main() {
	taskMethods := map[string]func(){
		"add":    task.NewTask,
		"query":  task.QueryTask,
		"delete": task.DeleteTask,
		"modify": task.ModifyTask,
	}

	userMethods := map[string]func(){
		"add":    user.NewUser,
		"query":  user.QueryUser,
		"delete": user.DeleteUser,
		"modify": user.ModifyUser,
	}
	for {
		choiceMode, err := strconv.Atoi(utils.Input("请选择:\n登录:输入1\n注册:输入2\n退出:3\n"))
		if err != nil {
			fmt.Println("您输入的操作符非法.")
			continue
		}
		if choiceMode == 1 {
			pwhash, err := user.GetUserPasswd(utils.Input("请输入账号："))
			if err != nil {
				fmt.Println(err)
				continue
			}
			if passwd.CheckPasswd(pwhash, 3) {
				for {
					opt, err := strconv.Atoi(utils.Input("请选择操作:\n设置用户:1\n操作任务:2\n退出:3\n"))
					if err != nil {
						fmt.Println("您输入的操作符非法.")
						continue
					}
					if opt == 1 {
						for {
							text := utils.Input("请输入操作(add/query/delete/modify/exit)：")
							if text == "exit" {
								break
							}
							method, ok := userMethods[text]
							if ok {
								method()
							} else {
								fmt.Println("您输入的操作是非法的.")
							}
						}
					} else if opt == 2 {
						for {
							text := utils.Input("请输入操作(add/query/delete/modify/exit)：")
							if text == "exit" {
								break
							}
							method, ok := taskMethods[text]
							if ok {
								method()
							} else {
								fmt.Println("您输入的操作是非法的.")
							}
						}
					} else if opt == 3 {
						break
					} else {
						fmt.Println("您输入的操作符不正确,请重新输入.")
					}
				}
			} else {
				fmt.Println("密码输入错误.")
				continue
			}
		} else if choiceMode == 2 {
			user.NewUser()
		} else if choiceMode == 3 {
			break
		} else {
			fmt.Println("您输入的操作符不正确,请重新输入.")
		}
	}

	// for {
	// 	text := task.Input("请输入操作(add/query/delete/modify)：")
	// 	if text == "exit" {
	// 		break
	// 	}
	// 	method, ok := methods[text]
	// 	if ok {
	// 		method()
	// 	} else {
	// 		fmt.Println("您输入的操作是非法的.")
	// 	}
	// }
}
