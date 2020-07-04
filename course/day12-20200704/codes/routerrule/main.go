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
	// GET,HEAD => Query, Post =>Add, Put => Modify, Delete => Del

	// 自定义路由规则

	// 分号路由规则
	beego.Router("/task/", &TaskController{}, "get,head:Query;post:Add;put:Modify;delete:Del")
	beego.Run()

	//url => 函数/方法
}
