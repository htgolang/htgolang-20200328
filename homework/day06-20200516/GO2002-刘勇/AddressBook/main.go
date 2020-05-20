package main

import (
	"fmt"
	"homework/htgolang-202003/homework/day06-20200516/GO2002-刘勇/AddressBook/utils"
	"io/ioutil"
	"os"
)

func main() {

	var menu = `
----------------------通讯录----------------------

                   1.  添  加
                   2.  查  询
                   3.  删  除
                   4.  修  改
请输入（1-4或Q退出）:`

	for {

		fmt.Print(menu)
		var user utils.User
		var input string
		fmt.Scan(&input)
		//fmt.Println(input)
		if input == "q" || input == "Q" {
			fmt.Println("谢谢使用，再见")
			os.Exit(0)
		}

		//如果文件不存在，则强制只能添加用户
		if input != "1" && !utils.IsFile(utils.TxtFile) {
			fmt.Println("用户不存在，请先添加用户")
			continue
		}

		//如果是空文件也提示添加用户，拒绝其他操作
		b, _ := ioutil.ReadFile(utils.TxtFile)
		if input != "1" && string(b) == "" {
			fmt.Println("当前无用户，请先添加用户")
			continue
		}

		switch input {
		case "1":
			fmt.Println("添加用户")
			utils.AddUser(&user)
		case "2":
			fmt.Println("查询用户")
			utils.Query()
		case "3":
			fmt.Println("删除用户")
			utils.DelUser(&user)
		case "4":
			fmt.Println("修改用户")
			utils.ModifyUser(&user)
		default:
			fmt.Println("输入错误，请重新输入，退出（q）")
		}
	}
}
