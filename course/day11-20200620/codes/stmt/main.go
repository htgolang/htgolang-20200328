package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=PRC",
		"golang", "golang@2020", "127.0.0.1", 3306, "todolist")

	db, _ := sql.Open("mysql", dsn)
	start := time.Now()
	stmt, _ := db.Prepare(`insert into account(name, money) values(?, ?)`)
	for i := 0; i < 10000; i++ {
		stmt.Exec(fmt.Sprintf("a_", i), 1000)
	}
	fmt.Println(time.Now().Sub(start))

}
