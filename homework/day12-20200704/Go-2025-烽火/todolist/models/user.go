package models

import (
	"log"
	"todolist/base"
)

type User struct {
	ID       int
	Name     string
	Status   int
	Password string
	Flag     int
}

func GetAccounts() []User {
	accounts := make([]User, 0)
	if rows, err := TDB.Query(base.SqlQueryAllUser); err != nil {
		log.Fatal(err)
	} else {
		for rows.Next() {
			user := User{}
			err := rows.Scan(&user.ID, &user.Name, &user.Status)
			if err != nil {
				log.Println(err)
			} else {
				accounts = append(accounts, user)
			}
		}
	}

	return accounts
}
