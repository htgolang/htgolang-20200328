package main

import (
	//系统
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/session/redis"

	//自己
	"github.com/strive-after/go-cmdb/module"
	"github.com/strive-after/go-cmdb/util"
	_ "github.com/strive-after/go-cmdb/module"
	_ "github.com/strive-after/go-cmdb/route"
)




func main() {
	//未注册  必须注册才能用 有时候session获取的时候  它会提示  gob: name not registered for interface: "github.com/strive-after/go-kubernetes/module.User"
	//那么就注册一下
	gob.Register(module.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("redisurl")
	beego.AddFuncMap("TimeForMat",util.StopTimeFormat)
	beego.AddFuncMap("Roles",util.Role)
	beego.AddFuncMap("Left",util.IndexLeft)
	beego.AddFuncMap("Right",util.IndexRight)
	err := logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)  //separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	if err != nil {
		beego.Error(err)
		return
	}
	beego.Run()
}
