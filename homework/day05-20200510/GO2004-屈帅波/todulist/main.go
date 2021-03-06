package main

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"todulist/controls"
	"todulist/module"
	_ "todulist/module"
	_ "todulist/route"
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
	beego.BConfig.WebConfig.Session.SessionOn = true
	//自定义错误控制器
	beego.ErrorController(&controls.ErrControls{})
	beego.AddFuncMap("TimeForMat",StopTimeFormat)
	beego.AddFuncMap("Roles",module.Role)
	beego.AddFuncMap("Left",IndexLeft)
	beego.AddFuncMap("Right",IndexRight)
	beego.Run()
}
