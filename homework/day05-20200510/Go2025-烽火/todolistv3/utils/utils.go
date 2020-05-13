package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/howeyc/gopass"
)

const (
	saltLimit      = 6
	passwordLength = 6
)

func init() {
	rand.Seed(time.Now().Unix())
}

// 获取用户输入信息
func Input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}

// md5 hash
func md5text(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

//检查密码
func checkPassword(text, hash string) bool {
	pos := strings.LastIndex(hash, ":")
	if pos < 0 {
		return false
	}
	return hashPassword(text, hash[:pos]) == hash
}

// 生成salt
func generateSalt(limit int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	salt := make([]byte, limit)
	for i := 0; i < limit; i++ {
		salt[i] = chars[rand.Intn(len(chars))]
	}
	return string(salt)
}

// 密码hash
func hashPassword(text, salt string) string {
	if salt == "" {
		salt = generateSalt(saltLimit)
	}
	hash := md5text(fmt.Sprintf("%s:%s", salt, text))
	return fmt.Sprintf("%s:%s", salt, hash)
}

// 非明文显示终端输入信息，获取信息
// 0:设置密码 1:验证密码 2:设置新密码
func getPassword(flag int) string {
	switch flag {
	case 0:
		fmt.Print("请设置密码: ")
	case 1:
		fmt.Print("请输入密码: ")
	case 2:
		fmt.Print("请输入新密码: ")
	}
	passwd, _ := gopass.GetPasswdMasked()
	return string(passwd)
}

//检查文件是否存在
func FileIsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

// 从文件读取内容
func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	txt := make([]byte, 0, 100)
	ctx := make([]byte, 10)

	for {
		n, err := file.Read(ctx)
		if err == io.EOF {
			break
		}
		txt = append(txt, ctx[:n]...)
	}
	return string(txt)
}

// 文件写操作函数
func WriteFile(path, txt string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(txt)
}

//设置密码
func SetPassword(path string, flag int) {
	for {
		password := getPassword(flag)
		if len(password) >= passwordLength {
			salt := generateSalt(saltLimit)
			md5password := hashPassword(password, salt)
			WriteFile(path, md5password)
			fmt.Println("密码设置成功!")
			break
		} else {
			fmt.Printf("密码长度不能少于%d位!\n", passwordLength)
		}
	}
}

// 验证密码
func VerifyPassword(path, password string, limit int) bool {
	hasher := ReadFile(path)
	for count := 0; count < limit; count++ {
		input := getPassword(1)
		if checkPassword(input, hasher) {
			return true
		} else {
			fmt.Printf("密码验证错误，还剩%d次机会!\n", limit-count-1)
		}
	}
	return false
}

// 修改密码
func ChangePassword(path string) {
	switch Input("是否需要修改密码，请确认[y/yes]: ") {
	case "y", "yes":
		SetPassword(path, 2)
	default:
		fmt.Println("取消修改密码!")
	}
}
