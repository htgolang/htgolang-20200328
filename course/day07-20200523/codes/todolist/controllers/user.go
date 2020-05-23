package controllers

import (
	"fmt"
	"os"
	"todolist/config"
	"todolist/utils/ioutils"
	"todolist/views"
)

func Login() bool {
	views.LoginTitle()
	for i := config.Config.LoginRetry; i > 0; i-- {
		txt := ioutils.Password("请输入密码: ")
		if txt == "kk" {
			return true
		}
		if i != 1 {
			ioutils.Error(fmt.Sprintf("密码错误, 剩余登录%d次数", i-1))
		}
	}

	ioutils.Error(fmt.Sprintf("密码错误超过%d, 程序退出", config.Config.LoginRetry))
	return false
}

func Logout() {
	os.Exit(0)
}
