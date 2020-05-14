package passwd

import (
	"errors"
	"fmt"

	"github.com/shadow_wei/TaskManage/utils"
)

func CheckPasswd(hash string, checkNum int) bool {
	passwd := utils.PWInput("请输入您的密码：")
	if ComparePasswd(passwd, hash) {
		return true
	}
	for i := 2; i <= checkNum; i++ {
		passwd := utils.PWInput(fmt.Sprintf("密码错误,请第%d次输入您的密码：", i))
		if ComparePasswd(passwd, hash) {
			return true
		}
	}
	return false
}

func NewPasswd() (string, error) {
	newPasswd := utils.PWInput("设置密码：")
	confirmPasswd := utils.PWInput("确认密码：")
	hashPasswd := HashPasswd(newPasswd, "")
	if ComparePasswd(confirmPasswd, hashPasswd) {
		return hashPasswd, nil
	}
	return "", errors.New("密码不一致.")
}
