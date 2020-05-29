package controllers

import (
	"fmt"
	"myTodolist/commands/config"
	"myTodolist/utils/ioutils"
	"os"
)

//定义一个退出函数
func Logout() {
	os.Exit(0)
}

//定义一个登陆函数,返回一个bool值，true时则登陆成功，false时则登陆失败(退出)。
func Login() bool {
	for i := config.Config.LoginRetry; i>0; i-- {
		pass := ioutils.Password("请输入密码：")
		if pass == "hello" {
			return true
		}
		if i != 1 {
			ioutils.Error(fmt.Sprintf("密码错误，还剩余%d次机会",i-1))
		}
	}
	ioutils.Error(fmt.Sprintf("密码错误超过%d次,程序退出！",config.Config.LoginRetry))
	return false
}

