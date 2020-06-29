package user

import (
	"log"
	"os"
	"todolist/db"
)

const (
	sqlQueryAllUser = `select * from user`
)

type User struct {
	ID     string
	Name   string
	Status string
}

func NewUser() *User {
	return &User{}
}

func init() {
	// 设置log
	logfile, _ := os.OpenFile("user.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	defer logfile.Close()
	log.SetOutput(logfile)

}

func (u *User) GetAccounts() []User {
	var accounts []User
	rows, err := db.Config.DB.Query(sqlQueryAllUser)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Status)
		if err != nil {
			log.Println(err)
		} else {
			accounts = append(accounts, user)
		}
	}

	return accounts
}
