package init

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.SetLogger(
		logs.AdapterFile,
		fmt.Sprintf(
			`{"filename":"%s","level":%d,"maxlines":%d,"maxsize":%d,"daily":%t,"maxdays":%d,"color":true,"perm":%s}`,
			beego.AppConfig.DefaultString("log::LogFile", "main.log"),
			beego.AppConfig.DefaultInt("log::Level", 7),
			beego.AppConfig.DefaultInt("log::MaxLines", 10000),
			beego.AppConfig.DefaultInt64("log::MaxSize", 1<<30),
			beego.AppConfig.DefaultBool("log::Daily", true),
			beego.AppConfig.DefaultInt("log::MaxDays", 7),
			beego.AppConfig.DefaultString("log::Perm", "0664"),
		),
	)
}
