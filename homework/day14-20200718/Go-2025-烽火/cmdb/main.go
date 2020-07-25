package main

import (
	"fmt"
	"log"

	_ "cmdb/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.Debug = true
	// 数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		beego.AppConfig.DefaultString("mysql::user", "devon"),
		beego.AppConfig.DefaultString("mysql::password", "golang@2020"),
		beego.AppConfig.DefaultString("mysql::host", "127.0.0.1"),
		beego.AppConfig.DefaultInt("mysql::port", 3306),
		beego.AppConfig.DefaultString("mysql::dbName", "cmdb"),
	)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)

	if db, err := orm.GetDB("default"); err != nil {
		log.Fatal(err)
	} else if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	beego.Run()
}
