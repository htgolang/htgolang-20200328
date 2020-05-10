package main

import (
	"fmt"

	"password/utils"

	"github.com/howeyc/gopass"
)

const hash = "mA2hbU$$b61929ec80ce989ef810c250b48a7327"

func main() {
	fmt.Print("请输入密码:")
	password, _ := gopass.GetPasswd()
	if utils.CheckPassword(string(password), hash) {
		fmt.Println("成功")
	} else {
		fmt.Println("失败")
	}

}
