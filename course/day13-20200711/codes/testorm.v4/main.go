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

	ormer := orm.NewOrm()
	// rawseter := ormer.Raw("insert into user(name, created_at, update_at, birthday) values(?, ?, ?, ?)", "1", "2020-10-01", "2020-10-01", "2020-10-01")
	// fmt.Println(rawseter.Exec())

	// rawseter = ormer.Raw("update user set name=? where id=?", "1xxxx", 1)
	// fmt.Println(rawseter.Exec())

	// rawseter = ormer.Raw("delete from user  where id=?", 2)
	// fmt.Println(rawseter.Exec())
	rawseter := ormer.Raw("select id,name from user")
	var users []*User
	rawseter.QueryRows(&users)
	fmt.Printf("%#v\n", users[0])
	var user *User

	rawseter = ormer.Raw("select id,name from user where id = ?", 100)
	rawseter.QueryRow(&user)
	fmt.Printf("%#v\n", user)

	rawseter = ormer.Raw("select name, count(*) as cnt from user group by name")
	var result []orm.Params
	rawseter.Values(&result)

	fmt.Println(result)

	rawseter = ormer.Raw("select name, count(*) as cnt from user group by name")

	var resultList []orm.ParamsList
	rawseter.ValuesList(&resultList)
	fmt.Println(resultList)

}
