package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"todolist/base"
	"todolist/forms"
	"todolist/models"
	"todolist/utils"

	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

type Task struct {
	ID           int
	Name         string
	StartTime    *time.Time
	CompleteTime *time.Time
	DeadlineTime *time.Time
	Status       string
	Content      string
	User         string
}

func (t *TaskController) Index() {
	var tasks = make([]Task, 0, 20)
	if rows, err := models.TDB.Query(base.SqlQueryAllTask); err == nil {
		for rows.Next() {
			var task Task
			var status int

			err := rows.Scan(&task.ID, &task.Name, &status, &task.StartTime, &task.CompleteTime, &task.DeadlineTime, &task.User, &task.Content)
			if err != nil {
				log.Println(err)
			} else {
				task.Status = base.StatusMap[status]
				tasks = append(tasks, task)
			}
		}
	}

	t.Data["tasks"] = tasks
	t.TplName = "index.html"
}

// 添加任务
func (t *TaskController) Add() {
	var (
		task   forms.TaskForm
		errMsg = make(map[string]string)
	)
	users := models.GetAccounts()

	if t.Ctx.Input.IsPost() {
		name := strings.TrimSpace(t.GetString("name"))
		content := strings.TrimSpace(t.GetString("content"))
		startTime := strings.TrimSpace(t.GetString("start_time"))
		completeTime := strings.TrimSpace(t.GetString("complete_time"))
		deadlineTime := strings.TrimSpace(t.GetString("deadline_time"))
		suid := strings.TrimSpace(t.GetString("user"))
		uid, _ := strconv.Atoi(suid)
		// 检查任务名称
		if err := utils.CheckTaskName(name); err != nil {
			errMsg["name"] = err.Error()
		}
		// 检查日期
		st := strings.ReplaceAll(startTime, "T", " ")

		dt := strings.ReplaceAll(deadlineTime, "T", " ")
		if err := utils.CheckDeadline(dt); err != nil {
			errMsg["deadlineTime"] = err.Error()
		}

		// 检查任务描述
		if err := utils.CheckContent(content); err != nil {
			errMsg["content"] = err.Error()
		}

		task.Name = name
		task.Content = content
		task.StartTime = st
		task.DeadlineTime = dt
		task.User = uid

		if len(errMsg) == 0 {
			if completeTime != "" {
				ct := strings.ReplaceAll(completeTime, "T", " ")
				models.TDB.Exec(base.SqlCreateTaskWithCt, task.Name, task.Content, st, ct, dt, uid)
			} else {
				models.TDB.Exec(base.SqlCreateTask, task.Name, task.Content, st, dt, uid)
			}
			t.Redirect("/", http.StatusFound)
		}
	}

	t.Data["task"] = task
	t.Data["users"] = users
	t.Data["errors"] = errMsg
	t.TplName = "task/add.html"
}

// 查询任务
func (t *TaskController) Query() {
	var (
		task    Task
		status  int
		message = map[string]interface{}{
			"flag": false,
		}
	)

	// GET请求带参数
	tid := t.GetString("id")
	if id, err := strconv.Atoi(tid); err == nil {
		task.ID = id
	}

	row := models.TDB.QueryRow(base.SqlQueryTaskWithUserName, tid)
	err := row.Scan(&task.Name, &status, &task.StartTime, &task.CompleteTime, &task.DeadlineTime, &task.Content, &task.User)

	if err == nil {
		task.Status = base.StatusMap[status]
		message["flag"] = true
	} else {
		message["tips"] = "任务不存在"
	}

	t.Data["task"] = task
	t.Data["errors"] = message

	t.TplName = "task/query.html"
}

// // 修改任务
func (t *TaskController) Modify() {
	var (
		task             forms.TaskForm
		tempStartTime    string
		tempDeadlineTime string
		errMsg           = make(map[string]string)
	)

	users := models.GetAccounts()

	if t.Ctx.Input.IsGet() {
		id := strings.TrimSpace(t.GetString("id"))
		row := models.TDB.QueryRow(base.SqlQueryTask, id)
		err := row.Scan(&task.ID, &task.Name, &task.Status, &tempStartTime, &tempDeadlineTime, &task.Content, &task.User)
		if err != nil {
			log.Fatal(err)
		}
		//标记select用户
		for index, user := range users {
			if user.ID == task.User {
				user.Flag = 1
				users[index] = user
			}
		}

		if tempStartTime != "" {
			startTime := strings.ReplaceAll(tempStartTime, ":00Z", "")
			task.StartTime = startTime
		}

		deadlineTime := strings.ReplaceAll(tempDeadlineTime, ":00Z", "")
		task.DeadlineTime = deadlineTime
	} else if t.Ctx.Input.IsPost() {
		id := strings.TrimSpace(t.GetString("id"))
		name := strings.TrimSpace(t.GetString("name"))
		status := strings.TrimSpace(t.GetString("status"))
		startTime := strings.TrimSpace(t.GetString("start_time"))
		deadlineTime := strings.TrimSpace(t.GetString("deadline_time"))
		content := strings.TrimSpace(t.GetString("content"))
		uid := strings.TrimSpace(t.GetString("user"))

		// 检查任务名称
		if err := utils.CheckTaskName(name); err != nil {
			errMsg["name"] = err.Error()
		}

		// 检查截止日期
		dt := strings.ReplaceAll(deadlineTime, "T", " ")
		if err := utils.CheckDeadline(dt); err != nil {
			errMsg["deadlineTime"] = err.Error()
		}

		// 检查任务描述
		if err := utils.CheckContent(content); err != nil {
			errMsg["content"] = content
		}

		if status == "3" {
			completeTime := time.Now().Format(base.TimeLayout)
			models.TDB.Exec(base.SqlUpdateTaskIncludeTime, name, status, startTime, deadlineTime, completeTime, content, uid, id)
		} else {
			models.TDB.Exec(base.SqlUpdateTask, name, status, startTime, deadlineTime, content, uid, id)
		}
		t.Redirect("/", http.StatusFound)

	}

	t.Data["task"] = task
	t.Data["user"] = users
	t.Data["errors"] = errMsg
	fmt.Printf("%#v\n", errMsg)
	t.TplName = "task/modify.html"
}

// 删除任务
func (t *TaskController) Del() {
	id := t.GetString("id")
	models.TDB.Exec(base.SqlDeleteTask, id)
	t.Redirect("/", http.StatusFound)
}
