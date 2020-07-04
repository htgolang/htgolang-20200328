package main

import (
	"github.com/astaxie/beego"
)

// 定义控制器HomeController
type HomeController struct {
	beego.Controller
}

// 1. Restful风格的控制器
// api格式
// 一切皆资源 url=> 资源
// 通过动作来表示对资源的操作类型 http method =>
// Post => 创建/更新
// Get => 获取/查询
// Delete => 删除
// Put => 更新

// 可以处理Get请求
func (c *HomeController) Get() {
	c.Ctx.WriteString("Get")
}

// 可以处理Post请求
func (c *HomeController) Post() {
	c.Ctx.WriteString("Post")
}

func (c *HomeController) Delete() {
	c.Ctx.WriteString("Delete")
}

func (c *HomeController) Put() {
	c.Ctx.WriteString("Put")
}

func (c *HomeController) Head() {
	c.Ctx.WriteString("Head")
}

func (c *HomeController) Options() {
	c.Ctx.WriteString("Options")
}

type HostController struct {
	beego.Controller
}

func (c *HostController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("Host:" + id)
}

func main() {
	// url和控制器绑定 => 路由
	beego.Router("/home/", &HomeController{})
	beego.Router("/host/?:id/", &HostController{})
	beego.Run()
}
