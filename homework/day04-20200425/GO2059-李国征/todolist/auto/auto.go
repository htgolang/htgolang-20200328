package auto

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/howeyc/gopass"
)

// 添加登录用户
var users = []map[string]string{
	{"name": "zhangsan", "password": "09cNIzlDZ2^e10adc3949ba59abbe56e057f20f883e"},
	{"name": "lisi", "password": "2Ni9H5&gTN^c33367701511b4f6020ec61ded352059"},
}

func randomCode(n int) []byte {
	// 随机码生成 ^
	baseStr := []byte("1234567890abcdefghigklmnopqrstuvwxyzABCDEFGHIGKLMNOPQRSTUVWXYZ!@#$%&*")
	initStr := make([]byte, 0, len(baseStr))
	for i := 0; i < n; i++ {
		initStr = append(initStr, baseStr[rand.Intn(len(baseStr))])
	}
	return initStr
}

func Md5Salt(pwd string) string {
	// md5 加盐操作
	// 随机码 + “分割” + 密码

	randomStr := randomCode(10)
	middStr := []byte("^")
	pwdStr := []byte(pwd)
	m := md5.Sum(pwdStr)
	md5Str := fmt.Sprintf("%s%s%x", string(randomStr), string(middStr), m)
	return md5Str
}

func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}

func checkInputInfo(name, pwd string) bool {
	// 验证用户名密码是否存在
	for _, user := range users {
		if user["name"] != name {
			return false
		}
		pwd := Md5Salt(pwd)
		inputPWD := strings.Split(pwd, "^")[1]
		userPWD := strings.Split(user["password"], "^")[1]
		if inputPWD != userPWD {
			return false
		}
		return true
	}
	return true
}

func AuthUser() string {
	num := 1
	for {
		name := input("请输入用户名:")
		fmt.Print("请输入密码:")
		pwd, _ := gopass.GetPasswd()
		status := checkInputInfo(name, string(pwd))
		if status == false {
			if num == 3 {
				fmt.Println("登录失败，程序退出!")
				os.Exit(1)
			}
			fmt.Println("输入的用户名或密码错误，请重新输入！已尝试登录", num, "次! 最多3 次机会")
			num++
			continue
		}
		break
	}
	return "认证通过，欢迎使用XXX系统！"
}
