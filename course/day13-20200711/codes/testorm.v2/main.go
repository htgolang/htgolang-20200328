package main

import (
	"fmt"
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

	dsn := "golang:golang@2020@tcp(localhost:3306)/testorm?charset=utf8mb4&parseTime=true&loc=PRC"
	orm.RegisterDriver("mysql", orm.DRMySQL) // 可省略
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RegisterModel(new(User))

	// 增加

	birthday, _ := time.Parse("2006-01-02", "1988-01-01")
	user := &User{
		Name:     "kk",
		Password: "abc123",
		Gender:   true,
		Tel:      "152xxxx",
		Birthday: &birthday,
	}

	// 数据库连接
	ormer := orm.NewOrm()
	/*
		fmt.Printf("%#v\n", user)
		id, err := ormer.Insert(user) // 插入数据
		fmt.Println(id, err)
		fmt.Printf("%#v\n", user)

		users := make([]*User, 3, 10)
		for i := 0; i < 3; i++ {
			user := &User{
				Name:     fmt.Sprintf("kk-%d", i),
				Password: "abc123",
				Gender:   true,
				Tel:      "152xxxx",
				Birthday: &birthday,
			}
			users[i] = user
		}

		ormer.InsertMulti(2, users)
	*/

	user = &User{ID: 1}
	err := ormer.Read(user)
	fmt.Println(user, err)

	user = &User{ID: 1}
	err = ormer.Read(user)
	fmt.Println(user, err)

	user = &User{Name: "kk", Password: "aaa"}
	err = ormer.Read(user, "Name", "Password")
	fmt.Println(user, err)

	user = &User{ID: 1}
	err = ormer.Read(user)
	fmt.Println(user, err)
	user.Name = "浩泽宇"
	ormer.Update(user, "Name")

	fmt.Println(ormer.Delete(&User{ID: 122}))
	fmt.Println(ormer.Delete(&User{ID: 122}))

	user = &User{
		Name:     "kk",
		Password: "abc123",
		Gender:   true,
		Tel:      "152xxxx",
		Birthday: &birthday,
	}
	ormer.Insert(user)
	user.Name = "kkxxx"
	ormer.Update(user)

	user = &User{
		Name: "kk2",
	}

	ormer.ReadOrCreate(user, "Name")
	fmt.Println(user)
}
