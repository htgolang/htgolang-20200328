package modles

import (
	"cmdb/utils"
	"fmt"
)

const (
	sqlQueryByName = "select id,name,password from user where name=? "
)

type User struct {
	ID         int
	Name       string
	StaffId    string
	NickName   string
	Password   string
	Gender     int
	Tel        string
	Addr       string
	Email      string
	Department string
	Status     int
}

func GetUserByName(name string) *User {
	user := &User{}
	fmt.Println(name)
	err := db.QueryRow(sqlQueryByName, name).Scan(&user.ID, &user.Name, &user.Password)
	fmt.Println(err)
	if err == nil {
		fmt.Println("modles GetUserByName")
		return user
	}
	return nil
}
func (u *User) ValidPassword(password string) bool {
	fmt.Println(password, u.Password)
	return u.Password == utils.Md5Text(password)
}
