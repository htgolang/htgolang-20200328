package main

import (
	"fmt"

	"github.com/astaxie/beego"
)

// 自定义控制器方法&路由规则
type TaskController struct {
	beego.Controller
}

// 添加任务
func (c *TaskController) Add() {
	fmt.Println("add")
	c.Ctx.WriteString("add")
}

// 查询任务*查看任务详情
func (c *TaskController) Query() {
	fmt.Println("query")
	c.Ctx.WriteString("query")
}

// 删除任务
func (c *TaskController) Del() {
	fmt.Println("del")
	c.Ctx.WriteString("del")
}

// 修改任务
func (c *TaskController) Modify() {
	fmt.Println("modify")
	c.Ctx.WriteString("modify")
}

func main() {

	// 自动路由
	// url => 控制 controller/action
	// task => TaskController
	// add => Add方法
	beego.AutoRouter(&TaskController{})
	beego.Run()
}
