package models

import (
	"fmt"
	"time"

	"cmdb/utils"
)

const (
	sqlQueryByName = "select id, name, password from user where name=?"
	sqlQuery       = "select id, staff_id, name, nickname, gender, tel, email, addr, department, status from user"
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
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}

// ValidPassword 验证用户密码是否正确
func (u *User) ValidPassword(password string) bool {
	fmt.Println(password, u.Password)
	return u.Password == utils.Md5Text(password)
}

// GenderText 性别显示
func (u *User) GenderText() string {
	if u.Gender == 0 {
		return "女"
	}
	return "男"
}

// StatusText 状态显示
func (u *User) StatusText() string {
	switch u.Status {
	case 0:
		return "正常"
	case 1:
		return "锁定"
	case 2:
		return "离职"
	}
	return "未知"
}

// GetUserByName 通过用户名获取用户
func GetUserByName(name string) *User {
	user := &User{}
	if err := db.QueryRow(sqlQueryByName, name).Scan(&user.ID, &user.Name, &user.Password); err == nil {
		return user
	}
	return nil
}

// QueryUser 查询用户
func QueryUser(q string) []*User {
	users := make([]*User, 0, 10)
	sql := sqlQuery

	params := []interface{}{}
	q = utils.Like(q)
	if q != "" {
		sql += " WHERE staff_id like ? ESCAPE '/' OR name like ? ESCAPE '/' OR nickname like ? ESCAPE '/' OR tel like ? ESCAPE '/' OR email like ? ESCAPE '/' OR addr like ? ESCAPE '/' OR department like ? ESCAPE '/'"
		params = append(params, q, q, q, q, q, q, q)
	}

	rows, err := db.Query(sql, params...)
	if err != nil {
		return users
	}

	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.ID, &user.StaffID, &user.Name, &user.Nickname, &user.Gender, &user.Tel, &user.Email, &user.Addr, &user.Department, &user.Status); err == nil {
			users = append(users, user)
			fmt.Printf("%#v\n", user)
		}
	}
	return users
}
