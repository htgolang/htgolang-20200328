package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var TDB *sql.DB

func init() {
	// 数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		beego.AppConfig.DefaultString("mysql::user", "devon"),
		beego.AppConfig.DefaultString("mysql::password", "golang@2020"),
		beego.AppConfig.DefaultString("mysql::host", "127.0.0.1"),
		beego.AppConfig.DefaultInt("mysql::port", 3306),
		beego.AppConfig.DefaultString("mysql::dbName", "todolist"),
	)
	var err error
	TDB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// 测试数据库连接
	if err := TDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
