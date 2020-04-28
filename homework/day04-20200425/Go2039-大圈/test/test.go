package main

import (
	"fmt"
	"github.com/howeyc/gopass"
)


func main() {
	fmt.Print("请输入密码：")
	b, _ := gopass.GetPasswd()
	fmt.Println(string(b))
}
