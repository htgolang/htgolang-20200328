package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/models"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
)

type TaskController struct {
	auth.AuthController
}

// 添加任务
func (c *TaskController) Add() {
	task := &models.Task{}
	users := models.GetAllAccounts()

	if c.Ctx.Input.IsPost() {
		c.ParseForm(task)
		models.AddTask(task)
		action := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
		c.Redirect(beego.URLFor(action), http.StatusFound)
	}
	c.Data["task"] = task
	c.Data["users"] = users
	c.TplName = "task/add.html"
}

// 查询任务
func (c *TaskController) Query() {
	query := c.GetString("query")
	tasks := models.QueryTask(query)
	c.Data["tasks"] = tasks
	c.TplName = "task/query.html"
}

// 修改任务
func (c *TaskController) Modify() {
	tid := c.GetString("id")
	task := models.GetTaskById(tid)
	users := models.GetAllAccounts()
	// 设置task对应user的flag标签
	for index, user := range users {
		if task.User == user.ID {
			user.Flag = 1
			users[index] = user
		}
	}
	if c.Ctx.Input.IsPost() {
		c.ParseForm(task)
		models.UpdateTask(task)
		action := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
		c.Redirect(beego.URLFor(action), http.StatusFound)
	}
	fmt.Printf("task : %#v\n", task)
	fmt.Println("task time", *task.StartTime)
	c.Data["users"] = users
	c.Data["task"] = task
	c.TplName = "task/modify.html"
}

// 删除任务

func (c *TaskController) Delete() {
	tid := c.GetString("id")
	models.DeleteTask(tid)
	action := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
	c.Redirect(beego.URLFor(action), http.StatusFound)
}
