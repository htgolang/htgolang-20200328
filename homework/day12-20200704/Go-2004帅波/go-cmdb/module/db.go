package module

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	_ "github.com/astaxie/beego/config/xml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
	//db连接信息
	dbinfo string
	RedisClient *redis.Client
	redisurl string
	redispassword string

)
type Database struct {
	User string
	Password string
	Host string
	Database string
}
func NewDatabase() string{
	database :=   Database{
		User:     beego.AppConfig.String("mysqluser"),
		Password: beego.AppConfig.String("mysqlpass"),
		Host:     beego.AppConfig.String("mysqlurl"),
		Database: beego.AppConfig.String("mysqldb"),
	}
	return  fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",database.User,database.Password,database.Host,database.Database)
}

func init() {
	//注册数据库 mysql是数据库类型
	//拼接数据库连接信息
	dbinfo = NewDatabase()
	//初始化db
	db,err  = gorm.Open("mysql", dbinfo)
	if err != nil {
		fmt.Println("mysql打开失败",err)
		return
	}
	//创建表关联user结构体
	redisurl = beego.AppConfig.String("redisurl")
	redispassword = beego.AppConfig.String("redispassword")

	db.AutoMigrate(&User{})
}




