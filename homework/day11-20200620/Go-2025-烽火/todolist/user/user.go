package user

import (
	"database/sql"
	"log"
	"os"
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

func (u *User) GetAccounts(db *sql.DB) []User {
	var accounts []User
	rows, err := db.Query(sqlQueryAllUser)
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
