package user

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shadow_wei/TaskManage/passwd"
	"github.com/shadow_wei/TaskManage/utils"
)

// UserType 用户类型
type UserType struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	User         string `json:"user"`
	Age          int    `json:"Age"`
	Tel          int64  `json:"Tel"`
	Address      string `json:"Address"`
	Passwd       string `json:"passwd"`
	RegisterTime int64  `json:"regtime"`
}

// New 创建userType
func (user UserType) New(id, age int, tel int64, name, username, address, passwd string) *UserType {
	return &UserType{
		ID:           id,
		Name:         name,
		User:         username,
		Age:          age,
		Tel:          tel,
		Address:      address,
		Passwd:       passwd,
		RegisterTime: time.Now().Unix(),
	}
}

// IsEqual 根据属性名和值判断是否相等
func (user UserType) IsEqual(field, value string) bool {
	switch strings.ToLower(field) {
	case "id":
		id, err := strconv.Atoi(value)
		if err == nil {
			if user.ID == id {
				return true
			}
		}
	case "name":
		return strings.Contains(user.Name, value)
	case "user":
		return strings.Contains(user.User, value)
	case "age":
		age, err := strconv.Atoi(value)
		if err == nil {
			if user.Age == age {
				return true
			}
		}
	case "tel":
		tel, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			if user.Tel == tel {
				return true
			}
		}
	case "address":
		return strings.Contains(user.Address, value)
	}
	return false
}

// FormatSlice 转成切片格式
func (user UserType) FormatSlice(format []string) []string {
	userSlice := make([]string, 0)
	for _, filed := range format {
		switch strings.ToLower(filed) {
		case "id":
			userSlice = append(userSlice, strconv.Itoa(user.ID))
		case "name":
			userSlice = append(userSlice, user.Name)
		case "user":
			userSlice = append(userSlice, user.User)
		case "age":
			userSlice = append(userSlice, strconv.Itoa(user.Age))
		case "tel":
			userSlice = append(userSlice, strconv.FormatInt(user.Tel, 10))
		case "address":
			userSlice = append(userSlice, user.Address)
		case "passwd":
			userSlice = append(userSlice, user.Passwd)
		case "registertime":
			userSlice = append(userSlice, time.Unix(user.RegisterTime, 0).Format("2006-01-02 15:04:05"))
		default:
			userSlice = append(userSlice, "")
		}
	}
	return userSlice
}

// ModifyUser 修改用户数据
func (user *UserType) ModifyUser(field string) error {
	var err error
	switch strings.ToLower(field) {
	case "name":
		value := utils.Input("请输入您的账号：")
		if IsExitstUser(value) {
			err = errors.New("账号存在.")
		} else {
			user.Name = value
		}
	case "user":
		value := utils.Input("请输入您的用户名：")
		user.User = value
	case "age":
		value := utils.Input("请输入您的年龄：")
		age, err := strconv.Atoi(value)
		if err == nil {
			user.Age = age
		}
	case "tel":
		value := utils.Input("请输入您的联系方式：")
		tel, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			user.Tel = tel
		}
	case "address":
		value := utils.Input("请输入您的家庭地址：")
		user.Address = value
	case "passwd":
		if passwd.CheckPasswd(user.Passwd, 1) {
			password, err := passwd.NewPasswd()
			if err == nil {
				user.Passwd = password
			}
		} else {
			err = errors.New("密码输入错误无法修改.")
		}
	default:
		err = errors.New(fmt.Sprintf("您输入的%s字段是非法.", field))
	}
	return err
}
