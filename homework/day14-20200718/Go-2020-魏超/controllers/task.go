package controllers

import (
	"cmdb/forms"
	"cmdb/models"
	"net/http"
	"time"

	"cmdb/base/controllers/auth"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

// TaskController 任务信息控制器
type TaskController struct {
	auth.AuthorizationController
}

// Index 展示任务信息
func (c *TaskController) Index() {
	c.Data["Tasks"] = models.QueryTasks(c.GetString("key"))
	c.TplName = "task/index.html"
}

// Delete 删除任务信息
func (c *TaskController) Delete() {
	var (
		taskID int
		err    error
	)
	taskID, err = c.GetInt("tid")
	if err == nil {
		err = models.DeleteUser(taskID)
	}
	c.Redirect(c.URLFor("TaskController.Index"), http.StatusFound)
}

// Modify 修改任务信息
func (c *TaskController) Modify() {
	var (
		errors   map[string]string
		taskForm forms.TaskForm
	)
	if c.Ctx.Input.IsGet() {
		if taskID, err := c.GetInt("pk"); err == nil {
			if task := models.GetTaskByID(taskID); task != nil {
				taskForm.ID = task.ID
				taskForm.Name = task.Name
				taskForm.StatusID = task.StatusID
				if task.StartTime != nil {
					taskForm.StartTime = task.StartTime.Format(forms.DateTimeLayout)
				}
				if task.CompleteTime != nil {
					taskForm.CompleteTime = task.CompleteTime.Format(forms.DateTimeLayout)
				}
				if task.DeadlineTime != nil {
					taskForm.DeadlineTime = task.DeadlineTime.Format(forms.DateTimeLayout)
				}
				taskForm.UserID = task.UserID
				taskForm.Describe = task.Describe
			}
		}
	} else if c.Ctx.Input.IsPost() {
		err := c.ParseForm(taskForm)
		if err != nil {
			logs.Warn("parse form TaskForm error, %s", err)
			errors["submit"] = "server error"
		} else {
			valid := validation.Validation{}
			if ok, err := valid.Valid(&taskForm); err != nil {
				logs.Warn("valid handle error, %s", err)
				errors["submit"] = "server error"
			} else if !ok {
				for _, err := range valid.Errors {
					errors[err.Key] = err.Message
				}
			} else {
				// 针对提交的数据进行逻辑处理
				var (
					startTime    time.Time
					completeTime time.Time
					deadlineTime time.Time
				)

				startTime, _ = time.Parse(forms.DateTimeLayout, taskForm.StartTime)
				completeTime, _ = time.Parse(forms.DateTimeLayout, taskForm.CompleteTime)
				deadlineTime, _ = time.Parse(forms.DateTimeLayout, taskForm.DeadlineTime)

				if models.TaskStatusMap[taskForm.StatusID].Name == "正在进行" {
					startTime = time.Now()
				}

				if models.TaskStatusMap[taskForm.StatusID].Name == "已完成" {
					completeTime = time.Now()
				}

				// 更新数据
				if err := models.UpdateTask(taskForm.ID, orm.Params{
					"name":          taskForm.Name,
					"status_id":     taskForm.StatusID,
					"start_time":    &startTime,
					"complete_time": &completeTime,
					"deadline_time": &deadlineTime,
					"user_id":       taskForm.UserID,
					"describe":      taskForm.Describe},
				); err == nil {
					c.Redirect(c.URLFor("TaskController.Index"), http.StatusFound)
				} else {
					logs.Error("update data error, %s", err)
					errors["submit"] = "update data error"
				}
			}
		}
	}

	if taskStatusMap, ok := models.TaskStatusMap[taskForm.StatusID]; ok {
		c.Data["TaskStatusMap"] = taskStatusMap.Relation
	} else {
		c.Data["TaskStatusMap"] = []int{}
	}
	c.Data["Tasks"] = taskForm
	c.Data["Errors"] = errors
	c.TplName = "task/edit.html"
}

// Add 添加任务信息
func (c *TaskController) Add() {
	var (
		errors   map[string]string
		taskForm forms.TaskForm
	)
	if c.Ctx.Input.IsPost() {
		err := c.ParseForm(taskForm)
		if err != nil {
			logs.Warn("parse form TaskForm error, %s", err)
			errors["submit"] = "server error"
		} else {
			valid := validation.Validation{}
			if ok, err := valid.Valid(&taskForm); err != nil {
				logs.Warn("valid handle error, %s", err)
				errors["submit"] = "server error"
			} else if !ok {
				for _, err := range valid.Errors {
					errors[err.Key] = err.Message
				}
			} else {
				// 针对条数的数据进行逻辑处理
				deadlineTime, _ := time.Parse(forms.DateTimeLayout, taskForm.DeadlineTime)
				if deadlineTime.Sub(time.Now()) < 0 {
					errors["deadline_time"] = "the deadline time must be greater than the current time"
				} else {
					task := models.Task{
						Name:         taskForm.Name,
						StatusID:     taskForm.StatusID,
						StartTime:    nil,
						CompleteTime: nil,
						DeadlineTime: &deadlineTime,
						UserID:       taskForm.UserID,
						Describe:     taskForm.Describe,
					}
					err = models.AddTask(task)
					if err != nil {
						errors["submit"] = "add data err"
					} else {
						c.Redirect(c.URLFor("TaskController.Index"), http.StatusFound)
					}
				}
			}
		}
	}

	c.Data["Tasks"] = taskForm
	c.Data["Errors"] = errors
	c.TplName = "task/add.html"
}
