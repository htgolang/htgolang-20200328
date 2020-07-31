package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/services"
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
	task := &forms.FormTask{User: c.LoginUser.ID}
	flash := beego.NewFlash()

	if c.Ctx.Input.IsPost() {
		c.ParseForm(task)
		if c.LoginUser.IsAdmin == 1 {
			flash.Warning("管理员没有权限添加任务")
			flash.Store(&c.Controller)
			c.Redirect(beego.URLFor("TaskController.Query"), http.StatusFound)
			return
		} else {
			err := services.TaskService.Add(task)
			if err == nil {
				action := beego.AppConfig.DefaultString("auth::TaskAction", "TaskController.Query")
				c.Redirect(beego.URLFor(action), http.StatusFound)
				return
			} else {
				flash.Error("添加任务失败")
				flash.Store(&c.Controller)
			}
		}
	}

	c.Data["task"] = task
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "task/add.html"
}

// 查询任务
func (c *TaskController) Query() {
	beego.ReadFromRequest(&c.Controller)

	query := c.GetString("query")
	tasks := services.TaskService.Query(query, c.LoginUser)
	c.Data["tasks"] = tasks
	c.Data["is_admin"] = c.LoginUser.IsAdmin
	c.TplName = "task/query.html"
}

// 修改任务
func (c *TaskController) Modify() {
	users := services.UserService.GetAllAccounts()
	formTask := &forms.FormTask{}
	tid, _ := c.GetInt("id")
	task := services.TaskService.GetById(tid)

	if c.Ctx.Input.IsPost() {
		c.ParseForm(formTask)
		if err := services.TaskService.Update(formTask); err == nil {
			action := beego.AppConfig.DefaultString("auth::TaskAction", "TaskController.Query")
			c.Redirect(beego.URLFor(action), http.StatusFound)
		}
	}
	c.Data["users"] = users
	c.Data["task"] = task
	c.Data["is_admin"] = c.LoginUser.IsAdmin
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "task/modify.html"
}

// 删除任务
func (c *TaskController) Delete() {
	tid, _ := c.GetInt("id")
	if task := services.TaskService.GetById(tid); task != nil {
		if c.LoginUser.ID == task.User {
			if err := services.TaskService.Delete(tid); err == nil {
				action := beego.AppConfig.DefaultString("auth::TaskAction", "TaskController.Query")
				c.Redirect(beego.URLFor(action), http.StatusFound)
			}
		}
	}
}
