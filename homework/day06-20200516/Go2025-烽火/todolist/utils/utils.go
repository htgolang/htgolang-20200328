package utils

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io/ioutil"
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
const (
	normal = "普通用户"
	admin  = "管理员"
	root   = "超级管理员"
)

var (
	RoleList = []string{normal, admin, root}
	RoleMap  = map[string]string{
		"1": normal,
		"2": admin,
		"3": root,
	}
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
func CheckPassword(text, hash string) bool {
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
func GetPassword(flag int) string {
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
func ReadPassword(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

// 文件写操作函数
func WriteFile(path, txt string) {
	file, _ := os.Create(path)
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	writer.WriteString(txt)
}

// 设置密码
func SetPassword() string {
	for {
		password := GetPassword(1)
		if len(password) >= passwordLength {
			salt := generateSalt(saltLimit)
			return hashPassword(password, salt)
		} else {
			fmt.Printf("密码长度不能少于%d位!\n", passwordLength)
		}
	}
}

// 设置用户权限
func SetRole() string {
	for {
		roleId := Input("请设置用户权限[1.普通用户 2.管理员 3.超级管理员]: ")
		if role, ok := RoleMap[roleId]; ok {
			return role
		} else {
			fmt.Println("输入错误!")
		}
	}
}

//设置密码
// func SetPassword(path string, flag int) {
// 	for {
// 		password := getPassword(flag)
// 		if len(password) >= passwordLength {
// 			salt := generateSalt(saltLimit)
// 			md5password := hashPassword(password, salt)
// 			WriteFile(path, md5password)
// 			fmt.Println("密码设置成功!")
// 			break
// 		} else {
// 			fmt.Printf("密码长度不能少于%d位!\n", passwordLength)
// 		}
// 	}
// }
