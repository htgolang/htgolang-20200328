package module

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int
	Name string  `orm:"size(20)`
	Password string `orm:"size(100)`
}


type Manage struct {
	Id int     //id
	//Name string     `orm:"size(20)` //名字
	StartTime time.Time    `orm:"auto_now_add";type(datetime)`//开始时间
	StopTime time.Time    `orm:"default(' ')`//停止时间
	TaskName string       `orm:"size(20)` //任务名称
	Taskinfo string     `orm:"size(100)`   //任务描述
	TaskStatus int     `orm:"default(0)`    //任务状态  0创建 1进行中 2暂停 3完成 4失败
	//default这里的默认值指的是当字段为null的时候默认为多少 如果字段默认是0那么default就不起到作用
}

func init() {
	//注册数据库 mysql是数据库类型
	orm.RegisterDataBase("default","mysql","root:123456@tcp(39.105.114.198:3306)/operation?charset=utf8&loc=Local")
	//注册表
	orm.RegisterModel(new(User),new(Manage))
	//中间的false 表示如果表存在我就不创建了  如果是true 表示如果表存在就更新他
	//后面的true  表示不存在就创建    如果是false  表示不执行
	orm.RunSyncdb("default",false,true)
}

