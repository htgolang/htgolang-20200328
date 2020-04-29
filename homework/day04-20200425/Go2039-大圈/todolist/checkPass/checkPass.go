package checkPass

import (
	"fmt"
	"github.com/howeyc/gopass"
	"todolist/hashValue"
)
var randStrings = "Fo$fcXdDntOjRBy%OauD"
var RealPass = "hello@123"
var count = 0

func CheckPass() bool {
	realhashValue := hashValue.Md5String("hello@123",randStrings)
	for  {
		fmt.Println("请输入登陆系统的密码：")
		InputPass,_ := gopass.GetPasswd()
		inputhashValue := hashValue.Md5String(string(InputPass),randStrings)
		if inputhashValue == realhashValue {
			fmt.Println("密码正确，准许进入系统！")
			return true
		}else {
			fmt.Println("密码错误，不允许进入系统！")
			count++
			if count == 3 {
				fmt.Println("密码错误3次，禁止登陆1小时")
				return false
			}
		}
	}
	return false
}
