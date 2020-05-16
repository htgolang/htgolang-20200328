package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/shadow_wei/TaskManage/passwd"
	"github.com/shadow_wei/TaskManage/utils"
)

// userInfo 定义用户信息的字典
var usersInfo map[int]*UserType

var dbFile = "userinfo.db"
var userFormat = []string{"id", "name", "user", "age", "tel", "address", "passwd", "registertime"}

func init() {
	execuFile, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("获取当前执行目录失败,%s\n", err))
	} else {
		dbFile = path.Join(filepath.Dir(execuFile), dbFile)
	}
	if dbInfo := []byte(ReadDB(dbFile)); len(dbInfo) > 0 {
		err = json.Unmarshal(dbInfo, &usersInfo)
		if err != nil {
			panic(fmt.Sprintf("无法从%s反序列化数据,请检查数据的有效性.%s", dbFile, err))
		}
	} else {
		usersInfo = make(map[int]*UserType)
	}
}

// NewUser 创建新用户
func NewUser() {
	var err error
	user := UserType{}
	id := GenerateUserID(usersInfo)
	username := utils.Input("请输入账户名：")
	if IsExitstUser(username) {
		fmt.Println("用户存在.")
		return
	}
	name := utils.Input("请输入用户名：")
	age, err := strconv.Atoi(utils.Input("请输入年龄："))
	if err != nil {
		fmt.Println("您输入的年龄格式不正确.")
		return
	}
	tel, err := strconv.ParseInt(utils.Input("请输入TEL联系方式："), 10, 64)
	if err != nil {
		fmt.Println("您输入的联系方式格式不正确.")
		return
	}
	address := utils.Input("请输入家庭地址：")
	passwd, err := passwd.NewPasswd()
	if err != nil {
		fmt.Println("您输入的密码不符合要求,", err)
		return
	}

	usersInfo[id] = UserType.New(user, id, age, tel, name, username, address, passwd)
	fmt.Println("用户创建成功!!")
	SaveDB(usersInfo, dbFile)
}

// QueryUser 查询用户信息
func QueryUser() {
	userSlice := make([][]string, 0)
	fmt.Println("字段信息：\nID\nName\nUser\nAge\nTel\nAddress\nRegisterTime")
	field := utils.Input("请输入要查询的字段名称：")
	value := utils.Input("请输入查询的信息：")

	for _, user := range usersInfo {
		if value == "all" || user.IsEqual(field, value) {
			userSlice = append(userSlice, user.FormatSlice(userFormat))
		}
	}
	utils.TableFormat(userFormat, userSlice)
}

// DeleteUser 删除用户信息
func DeleteUser() {
	userSlice := make([][]string, 0)
	userID := utils.Input("请输入需要删除的用户ID：")
	id, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println("输入的ID格式不正确.")
		return
	}
	user, ok := usersInfo[id]
	if ok == true {
		userSlice = append(userSlice, user.FormatSlice(userFormat))
	}
	utils.TableFormat(userFormat, userSlice)
	if isDelete := utils.Input("是否确认删除当前任务信息(y/yes):"); isDelete == "yes" || isDelete == "y" {
		delete(usersInfo, id)
	}
	SaveDB(usersInfo, dbFile)
}

// ModifyUser 修改用户信息
func ModifyUser() {
	userSlice := make([][]string, 0)
	userID, err := strconv.Atoi(utils.Input("请输入需要编辑的用户ID："))
	if err != nil {
		fmt.Println("输入的ID格式不正确.")
		return
	}
	user, ok := usersInfo[userID]
	if ok == false {
		fmt.Println("输入的用户ID不存在.")
		return
	}
	userSlice = append(userSlice, user.FormatSlice(userFormat))
	utils.TableFormat(userFormat, userSlice)

	if isDelete := utils.Input("是否确认修改此用户的信息(yes/y):"); isDelete == "yes" || isDelete == "y" {
		for {
			fmt.Println("字段信息：\nID\nName\nUser\nAge\nTel\nAddress\nPasswd")
			field := utils.Input("请输入您要修改的字段名称：")
			err := user.ModifyUser(field)
			if err != nil {
				fmt.Println(err, "修改数据失败.")
			} else {
				SaveDB(usersInfo, dbFile)
				fmt.Println("修改成功.")
			}
			if isExit := utils.Input("是否确定继续修改此用户信息(yes/y):"); !(isExit == "yes" || isExit == "y") {
				break
			}
		}
	}
}

// IsExitstUser 检测用户的账号是否存在
func IsExitstUser(fuser string) bool {
	for _, user := range usersInfo {
		if user.User == fuser {
			return true
		}
	}
	return false
}

// IsExitstID 检测用户ID是否村子
func IsExitstID(id int) bool {
	for _, user := range usersInfo {
		if user.ID == id {
			return true
		}
	}
	return false
}

// GetUserID 通过账户名称得到用户ID
func GetUserID(username string) (int, error) {
	for id, user := range usersInfo {
		if user.User == username {
			return id, nil
		}
	}
	return 0, errors.New("用户不存在.")
}

// GetUserName 根据用户的id获取用户的名称
func GetUserName(id int) string {
	user, _ := usersInfo[id]
	return user.Name
}

// GetUserPass 根据账户名获取用户的passwd信息
func GetUserPasswd(username string) (string, error) {
	for _, user := range usersInfo {
		if user.User == username {
			return user.Passwd, nil
		}
	}
	return "", errors.New("用户不存在.")
}
