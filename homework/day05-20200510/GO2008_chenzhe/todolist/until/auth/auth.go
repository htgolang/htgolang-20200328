package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/howeyc/gopass"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	saltLen = 6
)

func init()  {
	rand.Seed(time.Now().Unix())
}

//用户输入并且返回输入
func Input(prompt string) string {
	var text string
	fmt.Println(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}

//生成salt
func newSalt()string  {
	salt := make([]byte,0,64)
	var randNum int
	for i:=0;i<saltLen ;  {
		randNum = rand.Intn(57)+65
		if randNum<97 &&randNum>90{
			continue
		}else {
			salt = append(salt,byte(randNum))
			i++
		}
	}

	return string(salt)
}

//注册
func register()  (string,string) {
	var passwd1,passwd2 string
	for {
		passwd1 = Input("请输入新密码")
		passwd2 = Input("请重复输入密码")
		if passwd1 == passwd2{
			fmt.Println("注册成功")
			break
		}else {
			fmt.Println("两次密码不一致，请重新输入")
		}
	}
	return passwd1,newSalt()
}

//加载用户认证文件
func loadAuth()(passwd,salt string)  {
	file,err :=os.OpenFile(`passwd.txt`,os.O_RDWR|os.O_CREATE,os.ModePerm)
	defer file.Close()
	txt := make([]byte, 0, 1024*1024)
	crt := make([]byte, 1024)
	if err == nil{
		for {
			n,err := file.Read(crt)
			if err == io.EOF{
				break
			}
			txt = append(txt,crt[:n]...)
		}
	}
	if len(txt)==0{
		fmt.Println("现在没有密码，请注册")
		passwd,salt := register()
		file.Write([]byte(salt+"@"+hashPasswd(passwd+salt)))
		os.Exit(0)
	}
	passwdSlice :=strings.Split(string(txt),"@")
	return passwdSlice[1],passwdSlice[0]
}

//hash+salt
func hashPasswd(passwd string)string  {
	hash64 := sha256.New()
	hash64.Write([]byte(passwd))
	pass := hex.EncodeToString(hash64.Sum(nil))
	return pass
}

func Auth()  {
	passwd,salt := loadAuth()
	fmt.Println(passwd,salt)
	Auth := false
	for i:=0;i<3;i++ {
		fmt.Println("请输入密码")
		ok,err :=gopass.GetPasswd()
		if err != nil{
			fmt.Println("程序出错了")
			fmt.Println(err)
			os.Exit(2)

		}
		hash64 := sha256.New()
		hash64.Write(ok)
		hash64.Write([]byte(salt))
		pass := hex.EncodeToString(hash64.Sum(nil))
		fmt.Println(string(ok),pass)

		if pass == passwd{
			Auth = true
			fmt.Println("密码正确,登录成功")
			break
		}else {
			fmt.Println("密码错了")
		}

	}
	if !Auth {
		fmt.Println("密码错误3次，退出程序")
		os.Exit(2)
	}
}

