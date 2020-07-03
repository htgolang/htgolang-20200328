package main

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"github.com/strive-after/go-kubernetes/controls"
	"github.com/strive-after/go-kubernetes/module"
	_ "github.com/strive-after/go-kubernetes/module"
	_ "github.com/strive-after/go-kubernetes/route"
	//_ "github.com/astaxie/beego/session/redis"
)

func StopTimeFormat(stoptime time.Time) string{
	stop := stoptime.Format("2006-01-02")

	if stop == "0001-01-01" {
		return "无"
	}
	return stop
}

func IndexLeft(index int) string{
	num := strconv.Itoa(index-1)
	return  num
}
func IndexRight(index int) string{
	num := strconv.Itoa(index+1)
	return  num
}

func Role(num int) string{
	switch num {
	case 0:
		return "普通用户"
	case 1:
		return "管理员"
	case 2:
		return "超级管理员"
	}
	return ""
}

func main() {
	var port string
	port  = ":"+beego.AppConfig.String("port")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("redisurl")
	//自定义错误控制器
	beego.ErrorController(&controls.ErrControls{})
	beego.AddFuncMap("TimeForMat",StopTimeFormat)
	beego.AddFuncMap("Roles",module.Role)
	beego.AddFuncMap("Left",IndexLeft)
	beego.AddFuncMap("Right",IndexRight)
	beego.Run(port)
}
