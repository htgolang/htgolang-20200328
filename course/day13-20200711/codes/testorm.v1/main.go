package main

import (
	"time"

	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

// 变量命名
// 驼峰式 下划线式
type UserModel struct {
	ID        int    `orm:"column(id);pk;auto;"`
	Name      string `orm:"unique;description(姓名)"`
	Gender    bool
	Tel       string     `orm:"index"`
	Height    float32    `orm:"column(height);default(1.8)"`
	AAA       string     `orm:"column(description);size(1024)"`
	BBB       string     `orm:"column(bbb);type(text)"`
	Birthday  *time.Time `orm:"type(date);"`
	CreatedAt *time.Time `orm:"auto_now_add;"`
	UpdateAt  *time.Time `orm:"auto_now;"`
	DeletedAt *time.Time `orm:"null"`
}

func (m *UserModel) TableName() string {
	return "user"
}

func (m *UserModel) TableIndex() [][]string {
	return [][]string{
		{"AAA"},
		{"BBB"},
		{"BBB", "AAA"},
		{"BBB", "Name", "AAA"},
	}
}

func (m *UserModel) TableEngine() string {
	return "myisam"
}

func main() {
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

	orm.RegisterModel(new(UserModel))

	// orm.RunCommand()

	orm.RunSyncdb("default", true, true)
}
