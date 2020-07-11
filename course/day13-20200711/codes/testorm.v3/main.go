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

	// ormer.QueryTable("user")
	queryset := ormer.QueryTable(&User{})
	fmt.Println(queryset.Count())

	var users []*User
	queryset.All(&users)
	fmt.Println(users)
	// Where
	// Filter
	// 列名__条件, 对象
	// 条件
	/*
		关系 =([i]exact), <(lt), >(gt), >=(gte), <=(lte)
			in(in)
			like %content%([i]contains) start%([i]startswith) %end([i]endswith)


	*/
	fmt.Println(queryset.Filter("name__iexact", "KK").Count())
	fmt.Println(queryset.Filter("name__contains", "K%K").Count())
	fmt.Println(queryset.Filter("name__startswith", "K%K").Count())
	fmt.Println(queryset.Filter("name__endswith", "K%K").Count())
	fmt.Println(queryset.Filter("id__in", []int{1, 2}).Count())
	fmt.Println(queryset.Filter("id__gt", 5).Filter("id__lt", 10).Count())

	// Exclude
	fmt.Println(queryset.Exclude("name__iexact", "KK").Count())

	// 分页
	queryset.Limit(3).Offset(2).All(&users)
	// 排序
	queryset.OrderBy("Name").All(&users)
	queryset.OrderBy("-Name", "-Tel").All(&users)
	queryset.OrderBy("-Name", "-Tel").One(&users)

	// name like '%kk%' and (tel like '152%' or tel like '158%')
	cond := orm.NewCondition()

	condTel := orm.NewCondition()
	condTel = condTel.Or("tel__istartswith", "152").Or("tel__istartswith", "158")
	cond = cond.And("name__icontains", "kk").AndCond(condTel)

	queryset.SetCond(cond).All(&users, "id")
	fmt.Println(users[0])

	queryset.Filter("id__gt", 10).Update(orm.Params{
		"name":   "kk2",
		"height": orm.ColValue(orm.ColAdd, 10),
	})

	queryset.Filter("id__gt", 120).Delete()

}
