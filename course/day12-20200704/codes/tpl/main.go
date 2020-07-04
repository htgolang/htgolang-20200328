package main

import (
	"strings"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {
	c.Data["name"] = "kk"
	c.Data["body"] = true
	c.Data["scores"] = []float32{1, 2, 3, 4}
	c.Data["users"] = map[int]string{1: "kk", 2: "wc"}
	c.Data["content"] = "abc.ABC"
	c.TplName = "index.html" // 默认后缀只支持html, tpl
}

func (c *HomeController) Home() {
	// 若无任何响应，则加载控制器名称/Action名.tpl文件显示
}

func main() {

	beego.AddFuncMap("lower", func(in string) string {
		return strings.ToLower(in)
	})

	beego.AutoRouter(&HomeController{})
	beego.Run()
}
