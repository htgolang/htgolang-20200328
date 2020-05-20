package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	TxtFile = "user.json"
)

type User struct {
	Name    string
	Age     int
	Address string
	Email   string
	Phone   string
	Status  int
}

var users map[int]*User

func init() {
	users = make(map[int]*User, 0)

}

func InputString(prome string) string {
	var input string
	fmt.Print(prome)
	fmt.Scan(&input)
	return strings.TrimSpace(input)
}

func (u *User) InputUser() {
	var err error
	u.Name = InputString("请输入姓名：")

	for {
		u.Age, err = strconv.Atoi(InputString("请输入年龄："))
		if err != nil {
			fmt.Println("输入错误，请重新输入。")
			continue
		}
		break
	}
	u.Email = InputString("请输入Email：")
	u.Phone = InputString("请输入电话号码：")
	u.Address = InputString("请输入地址：")
	u.Status = 0

}

func TxtRead(f string) map[int]*User {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	err = json.Unmarshal(b, &users)
	if err != nil {
		fmt.Println(err)
	}
	return users
}

func TxtWrite(f string) {
	b, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(f, b, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}
}

func AddUser(u *User) {
	if IsFile(TxtFile) {
		users = TxtRead(TxtFile)
	} else {

	}
	//读取文件中的数据到users

	//读取users的长度，确定id号
	k := len(users)
	id := k + 1

	u.InputUser()
	//读取输入赋值
	users[id] = u

	//将修改后的users持久化到文件
	TxtWrite(TxtFile)
}

func Exists(f string) bool {
	_, err := os.Stat(f)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsDir(f string) bool {
	s, err := os.Stat(f)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(f string) bool {
	if !Exists(f) {
		return false
	}
	return !IsDir(f)
}

func Query() {
	users = TxtRead(TxtFile)
	for {
		input := ""
		fmt.Print("请输入你的选择（A，查询所有用户，B,根据条件查询，C返回主菜单）：")
		fmt.Scan(&input)

		switch {
		case input == "A" || input == "a":
			QueryAll(users)
			return
		case input == "B" || input == "b":
			input = ""
			fmt.Print("请输入查询条件：")
			fmt.Scan(&input)
			QuerySome(users, input)
			return
		case input == "C" || input == "c":
			return
		default:
			fmt.Println("输入错误，请新选择")
		}
	}
}

func QueryAll(users map[int]*User) {
	for n, v := range users {
		if v.Status == 0 {
			fmt.Println("ID：", n)
			fmt.Println("姓名：", v.Name)
			fmt.Println("年龄：", v.Age)
			fmt.Println("电话：", v.Phone)
			fmt.Println("地址：", v.Address)
			fmt.Println("邮箱：", v.Email)
			fmt.Println("-------------------------------------------")
		}
	}
	return
}

func QuerySome(users map[int]*User, s string) {
	for n, v := range users {
		if v.Status == 0 {
			if strings.Contains(v.Name, s) || strings.Contains(strconv.Itoa(v.Age), s) || strings.Contains(v.Phone, s) ||
				strings.Contains(v.Address, s) || strings.Contains(v.Email, s) {

				fmt.Println("ID：", n)
				fmt.Println("姓名：", v.Name)
				fmt.Println("年龄：", v.Age)
				fmt.Println("电话：", v.Phone)
				fmt.Println("地址：", v.Address)
				fmt.Println("邮箱：", v.Email)
				fmt.Println("-------------------------------------------")
			}
		}
	}
	return
}

func DelUser(u *User) {
	users = TxtRead(TxtFile)
	input := 0
	fmt.Print("输入需要删除的ID：")
	fmt.Scan(&input)

	//u = &users[input]
	//u.Status = 1
	//users[input] = u

	(*users[input]).Status = 1
	TxtWrite(TxtFile)
	fmt.Println(users[input].Name, "已经删除")
}

func ModifyUser(u *User) {
	users = TxtRead(TxtFile)
	input := 0
	fmt.Print("输入需要修改的ID：")
	fmt.Scan(&input)

	users[input].InputUser()

	TxtWrite(TxtFile)
}
