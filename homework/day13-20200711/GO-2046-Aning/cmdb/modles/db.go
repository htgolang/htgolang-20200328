package modles

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

var db *sql.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=PRC&parseTime=true",
		beego.AppConfig.DefaultString("mysql::User", "gostudy"),
		beego.AppConfig.DefaultString("myqsl::Password", "123456q!"),
		beego.AppConfig.DefaultString("mysql::Host", "120.79.60.117"),
		beego.AppConfig.DefaultInt("mysql::Port", 3306),
		beego.AppConfig.DefaultString("mysql::Database", "cmdb"),
	)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

}
