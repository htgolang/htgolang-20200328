package main

import (
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename" : "test.log"}`)
	beego.SetLevel(beego.LevelInformational)

	beego.SetLogFuncCall(true)
	beego.Debug("我是一个调试日志")
	beego.Informational("我是一个提醒")
	beego.Error("我是一个错误")
}
