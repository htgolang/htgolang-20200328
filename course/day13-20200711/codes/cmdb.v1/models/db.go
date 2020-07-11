package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

var db *sql.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/cmdb?charset=utf8mb4&loc=PRC&parseTime=true",
		beego.AppConfig.DefaultString("mysql::User", "golang"),
		beego.AppConfig.DefaultString("mysql::Password", "golang@2020"),
		beego.AppConfig.DefaultString("mysql::Host", "127.0.0.1"),
		beego.AppConfig.DefaultInt("mysql::Port", 3306),
	)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}
