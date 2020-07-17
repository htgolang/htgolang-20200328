package controllers

import (
	"cmdb/models"
)

// TaskController 任务信息控制器
type TaskController struct {
	authorization
}

// Index 展示任务信息
func (c *TaskController) Index() {
	c.Data["Tasks"] = models.QueryTasks(c.GetString("key"))
	c.TplName = "task/index.html"
}

// Delete 删除任务信息
func (c *TaskController) Delete() {
	c.
}

// Modify 修改任务信息
func (c *TaskController) Modify() {

}

// Add 添加任务信息
func (c *TaskController) Add() {

}
