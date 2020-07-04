package models

import (
	"cmdb/utils"
	"fmt"
)

const (
	sqlQueryByName = "select id, name, password from user where name=?"
)

// User 用户对象
type User struct {
	ID         int
	StaffID    string
	Name       string
	Nickname   string
	Password   string
	Gender     int
	Tel        string
	Addr       string
	Email      string
	Department string
	Status     int
}

// ValidPassword 验证用户密码是否正确
func (u *User) ValidPassword(password string) bool {
	fmt.Println(password, u.Password)
	return u.Password == utils.Md5Text(password)
}

// GetUserByName 通过用户名获取用户
func GetUserByName(name string) *User {
	user := &User{}
	if err := db.QueryRow(sqlQueryByName, name).Scan(&user.ID, &user.Name, &user.Password); err == nil {
		return user
	}
	return nil
}
