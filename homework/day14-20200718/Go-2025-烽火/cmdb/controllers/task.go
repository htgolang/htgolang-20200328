package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/models"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

const TimeLayout = "2006-01-02 15:04"

type TaskController struct {
	auth.AuthController
}

func formatTime(t string) *time.Time {
	if t != "" {
		ft, _ := time.Parse(TimeLayout, strings.ReplaceAll(t, "T", " "))
		return &ft
	}
	return nil
}

// 添加任务
func (c *TaskController) Add() {
	task := &forms.FormTask{}
	users := models.GetAllAccounts()

	if c.Ctx.Input.IsPost() {
		c.ParseForm(task)
		err := models.AddTask(task)
		if err == nil {
			action := beego.AppConfig.DefaultString("auth::TaskAction", "TaskController.Query")
			c.Redirect(beego.URLFor(action), http.StatusFound)
		}
	}
	c.Data["task"] = task
	c.Data["users"] = users
	c.TplName = "task/add.html"
}

// 查询任务
func (c *TaskController) Query() {
	var tasks []*models.Task
	query := c.GetString("query")
	tasks = models.QueryTask(query)
	c.Data["tasks"] = tasks
	c.TplName = "task/query.html"
}

// 修改任务
func (c *TaskController) Modify() {
	users := models.GetAllAccounts()
	var task *models.Task
	formTask := &forms.FormTask{}
	if c.Ctx.Input.IsPost() {
		c.ParseForm(formTask)
		if err := models.UpdateTask(formTask); err == nil {
			action := beego.AppConfig.DefaultString("auth::TaskAction", "TaskController.Query")
			c.Redirect(beego.URLFor(action), http.StatusFound)
		}
	} else if tid, err := c.GetInt("id"); err == nil {
		task = models.GetTaskById(tid)
	}
	c.Data["users"] = users
	c.Data["task"] = task
	c.TplName = "task/modify.html"
}

// 删除任务

func (c *TaskController) Delete() {
	tid, _ := c.GetInt("id")
	if task := models.GetTaskById(tid); task != nil {
		fmt.Println("-----", c.LoginUser.ID, task.User)
		if c.LoginUser.ID == task.User {
			if err := models.DeleteTask(tid); err == nil {
				action := beego.AppConfig.DefaultString("auth::TaskAction", "TaskController.Query")
				c.Redirect(beego.URLFor(action), http.StatusFound)
			}
		}
	}
}
