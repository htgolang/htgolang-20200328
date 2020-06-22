package module

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/astaxie/beego/config/xml"
)

func init() {
	//注册数据库 mysql是数据库类型
	//orm.RegisterDataBase("default","mysql","root:123456@tcp(39.105.114.198:3306)/operation?charset=utf8&loc=Local")
	//注册表
	db := Dbjoint()
	orm.RegisterDataBase("default","mysql",db)
	orm.RegisterModel(new(User),new(Manage),new(Group))
	//中间的false 表示如果表存在我就不创建了  如果是true 表示如果表存在就更新他
	//后面的true  表示不存在就创建    如果是false  表示不执行
	orm.RunSyncdb("default",false,true)
}



func Dbjoint() string {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	url := beego.AppConfig.String("mysqlurls")
	dbname := beego.AppConfig.String("mysqldb")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local", mysqluser, mysqlpass, url, dbname)
	return dsn
}
