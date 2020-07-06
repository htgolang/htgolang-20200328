package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"todolist/models"
	"todolist/utils"
	"unicode/utf8"
)

func ListTasks(response http.ResponseWriter, rquest *http.Request) {
	var (
		tasks []models.Task
		tmpl  *template.Template
	)
	tasks = models.GetTasks()
	funcMap := template.FuncMap{
		"datetime":   utils.FormatDatetime,
		"status":     utils.FormatTaskStatus,
		"headleuser": utils.FormatUserID,
	}
	tmpl = template.Must(template.New("list.html").Funcs(funcMap).ParseFiles("views/task/list.html"))
	tmpl.Execute(response, struct {
		Tasks []models.Task
	}{
		Tasks: tasks,
	})
}

func AddTask(response http.ResponseWriter, request *http.Request) {
	var (
		task         models.Task
		users        []models.User
		deadlineTime *time.Time
		errors       = make(map[string]string)
	)
	users = models.GetUsers()
	if request.Method == http.MethodGet {

	} else if request.Method == http.MethodPost {
		taskName := strings.TrimSpace(request.PostFormValue("task_name"))
		userID, _ := strconv.Atoi(request.PostFormValue("user_id"))
		deadlinetext := request.PostFormValue("deadline_time")
		describe := request.PostFormValue("describe")
		nameLength := utf8.RuneCountInString(taskName)

		// 验证任务名称不能超过32位
		if nameLength == 0 {
			errors["name"] = "任务名不能为空"
		} else if nameLength > 32 {
			errors["name"] = "任务名称太长，不能超过32位"
		}

		// 描述信息不可以大于512
		if utf8.RuneCountInString(task.Describe) > 512 {
			errors["describe"] = "任务描述不能超过512个字符"
		}

		// 判断时间是否为空，并且不能小于当前时间
		if deadlinetext != "" {
			deadline, _ := time.Parse(utils.DateTimeLayout, deadlinetext)
			deadlineTime = &deadline
			if deadlineTime.Sub(time.Now()) < 0 {
				errors["deadline_time"] = "截止日期必须大于当前时间"
			}
		} else {
			errors["deadline_time"] = "任务期限不能为空"
		}

		// 责任人不能为空
		if userID == 0 {
			errors["userid"] = "必须选择责任人"
		}

		task.Name = taskName
		task.DeadlineTime = deadlineTime
		task.UserID = userID
		task.Describe = describe
		task.Status = 1

		if len(errors) == 0 {
			if err := task.CreateTask(); err == nil {
				http.Redirect(response, request, "/task/list/", http.StatusFound)
			} else {
				errors["submit"] = "提交失败"
			}

		}
	}
	funcMap := template.FuncMap{
		"datetime": utils.FormatDatetime,
	}
	tmpl := template.Must(template.New("add.html").Funcs(funcMap).ParseFiles("views/task/add.html"))
	tmpl.Execute(response, struct {
		Task   models.Task
		Users  []models.User
		Errors map[string]string
	}{
		Task:   task,
		Users:  users,
		Errors: errors,
	})
}

func DelTask(response http.ResponseWriter, request *http.Request) {
	var (
		task models.Task
	)
	task.ID, _ = strconv.Atoi(request.FormValue("id"))
	task.DeleteTask()
	http.Redirect(response, request, "/task/list/", http.StatusFound)
}

func ModTask(response http.ResponseWriter, request *http.Request) {
	var (
		task   models.Task
		users  []models.User
		errors = make(map[string]string)
	)

	task.ID, _ = strconv.Atoi(request.FormValue("id"))
	task.GetTaskById()
	users = models.GetUsers()
	if request.Method == http.MethodGet {

	} else if request.Method == http.MethodPost {
		task.Describe = request.FormValue("describe")
		task.Status, _ = strconv.Atoi(request.FormValue("status"))
		task.UserID, _ = strconv.Atoi(request.FormValue("user_id"))

		// 描述信息不可以大于512
		if utf8.RuneCountInString(task.Describe) > 512 {
			errors["describe"] = "任务描述不能超过512个字符"
		}

		// 责任人不能为空
		if task.UserID == 0 {
			errors["userid"] = "必须选择责任人"
		}

		if len(errors) == 0 {
			if models.StatusMap[task.Status].Name == "正在进行" {
				now := time.Now()
				task.StartTime = &now
			}
			if models.StatusMap[task.Status].Name == "已完成" {
				now := time.Now()
				task.CompleteTime = &now
			}
			if task.UpdateTask() == nil {
				http.Redirect(response, request, "/task/list/", http.StatusFound)
			} else {
				errors["submit"] = "提交失败"
			}

		}
	}

	funcMap := template.FuncMap{
		"status": utils.FormatTaskStatus,
	}

	tmpl := template.Must(template.New("edit.html").Funcs(funcMap).ParseFiles("views/task/edit.html"))
	err := tmpl.Execute(response, struct {
		Task        models.Task
		Users       []models.User
		StatusSlice []int
		Errors      map[string]string
	}{
		Task:        task,
		Users:       users,
		StatusSlice: models.StatusMap[task.Status].Relation,
		Errors:      errors,
	})
	log.Println(err)
}
