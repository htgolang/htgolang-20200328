package todolist

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	dburl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=PRC&parseTime=true",
		beego.AppConfig.DefaultString("mysql::User", "root"),
		beego.AppConfig.DefaultString("mysql::Password", "oracle"),
		beego.AppConfig.DefaultString("mysql::Host", "127.0.0.1"),
		beego.AppConfig.DefaultInt("mysql::Port", 3306),
		beego.AppConfig.DefaultString("mysql::DB", "go_todolist"),
	)
	db, _ = gorm.Open("mysql", dburl)
}

type task struct {
	Id        int64  `gorm:"column:Id;PRIMARY_KEY"`
	Name      string `gorm:"column:Name"`
	StartTime string `gorm:"column:StartTime"`
	EndTime   string `gorm:"column:EndTime"`
	Status    string `gorm:"column:Status"`
	User      string `gorm:"column:User"`
}

func (t *task) TableName() string {
	return "task"
}

type user struct {
	Id         int64  `gorm:"column:Id;PRIMARY_KEY"`
	Username   string `gorm:"column:Username"`
	Password   string `gorm:"column:Password"`
	Salt       string `gorm:"column:Salt"`
	CreateTime string `gorm:"column:CreateTime"`
	UpdateTime string `gorm:"column:UpdateTime"`
}

func (u *user) TableName() string {
	return "user"
}
