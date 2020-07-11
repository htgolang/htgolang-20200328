package main

import (
	"time"

	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int        `orm:"column(id)"`
	Name      string     `orm:"size(64)"`
	Password  string     `orm:"size(1024)"`
	Gender    bool       `orm:""`
	Tel       string     `orm:""`
	Height    float32    `orm:""`
	Birthday  *time.Time `orm:"type(date);"`
	CreatedAt *time.Time `orm:"auto_now_add;"`
	UpdateAt  *time.Time `orm:"auto_now;"`
	DeletedAt *time.Time `orm:"null"`
}

func (user *User) TableName() string {
	return "kkuser"
}

func main() {
	orm.Debug = true
	// 0. 导入包
	// 1. 注册驱动
	// 2. 注册数据库
	// 3. 定义数据模型 model
	// 4. 注册数据模型
	// 5. 操作
	// 		同步表结构
	//		数据: 增，删，改，查

	dsn1 := "golang:golang@2020@tcp(localhost:3306)/testorm?charset=utf8mb4&parseTime=true&loc=PRC"
	dsn2 := "gostudy:123456q!@tcp(120.79.60.117:3306)/gostudy?charset=utf8mb4&parseTime=true&loc=PRC"

	orm.RegisterDriver("mysql", orm.DRMySQL) // 可省略

	orm.RegisterDataBase("default", "mysql", dsn1)
	orm.RegisterDataBase("aning", "mysql", dsn2)

	orm.RegisterModel(new(User))

	birthday, _ := time.Parse("2006-01-02", "1999-01-01")
	// orm.RunCommand()
	user := &User{
		Name:     "kk",
		Password: "abc123",
		Gender:   true,
		Tel:      "152xxxx",
		Birthday: &birthday,
	}

	ormer1 := orm.NewOrm()
	ormer1.Insert(user)

	ormer2 := orm.NewOrm()
	ormer2.Using("aning")
	ormer2.Insert(user)
}
