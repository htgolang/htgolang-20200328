package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/base/errors"
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
	task := &models.Task{}
	errors := errors.New()
	users := models.GetAllAccounts()

	if c.Ctx.Input.IsPost() {
		st := formatTime(c.GetString("start_time"))
		ct := formatTime(c.GetString("complete_time"))
		dt := formatTime(c.GetString("deadline_time"))
		c.ParseForm(task)
		task.StartTime = st
		task.CompleteTime = ct
		task.DeadlineTime = dt
		err := models.AddTask(task)
		if err == nil {
			action := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
			c.Redirect(beego.URLFor(action), http.StatusFound)
		} else {
			errors.Add("msg", "添加失败!")
		}
	}
	c.Data["task"] = task
	c.Data["users"] = users
	c.Data["errors"] = errors
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
