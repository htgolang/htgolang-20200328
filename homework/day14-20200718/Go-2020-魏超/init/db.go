package init

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	if err := ConnMySQLDB(); err != nil {
		panic("connect database faild.")
	}
}

func ConnMySQLDB() error {
	orm.Debug = beego.AppConfig.DefaultBool("mysql::Debug", false)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=PRC&parseTime=true",
		beego.AppConfig.DefaultString("mysql::User", "root"),
		beego.AppConfig.DefaultString("mysq::Password", ""),
		beego.AppConfig.DefaultString("mysql::Host", "127.0.0.1"),
		beego.AppConfig.DefaultInt("mysql::Port", 3306),
		beego.AppConfig.DefaultString("mysql::DataBase", "cmdb"),
	)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("cmdb", "mysql", dsn)

	// 打印链接信息
	fmt.Printf(
		"connect mysql %s:%d/%s",
		beego.AppConfig.DefaultString("mysql::Host", "127.0.0.1"),
		beego.AppConfig.DefaultInt("mysql::Port", 3306),
		beego.AppConfig.DefaultString("mysql::DataBase", "cmdb"),
	)

	// 测试链接
	if db, err := orm.GetDB("cmdb"); err != nil {
		return err
	} else if err := db.Ping(); err != nil {
		return err
	}
	return nil
}
