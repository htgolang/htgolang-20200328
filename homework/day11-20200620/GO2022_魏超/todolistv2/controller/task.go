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

var handls = map[string]func(http.ResponseWriter, *http.Request){
	"/task/list/": listTasks,
	"/task/add/":  addTask,
	"/task/del/":  delTask,
	"/task/mod/":  modTask,
}

func listTasks(w http.ResponseWriter, r *http.Request) {
	var (
		tasks []models.Task
		tmpl  *template.Template
	)
	tasks = models.GetTasks()
	funcMap := template.FuncMap{
		"datetime":   utils.FormatDatetime,
		"status":     utils.FormatStatus,
		"headleuser": utils.FormatUserID,
	}
	tmpl = template.Must(template.New("tasks.html").Funcs(funcMap).ParseFiles("views/task/tasks.html"))
	err = tmpl.Execute(response, struct {
		Tasks []Task
	}{
		Tasks: tasks,
	})
	if err != nil {
		log.Println(err)
	}
}

func addTask(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		task         models.Task
		users        []models.User
		deadlineTime *time.Time
		errors = make(map[string]string)
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
		describeLength := utf8.RuneCountInString(task.Describe)
		if describeLength > 512 {
			errors["describe"] = "任务描述不能超过512个字符"
		}

		// 判断时间是否为空，并且不能小于当前时间
		if deadlinetext != "" {
			deadline, _ := time.Parse(dateTimeLayout, deadlinetext)
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
		task.Status = 0

		if len(errors) == 0 {
			if err := task.CreateTask(); err == nil {
				http.Redirect(response, request, "/task/list/", http.StatusFound)
			} else {
				errors["subimt"] = "保存数据库失败"
			}

		}

		funcMap := template.FuncMap{
			"datetime": utils.FormatDatetime(t)
		}
		tmpl := template.Must(template.New("add_task.html").Funcs(funcMap).ParseFiles("views/task/add_task.html"))
		err = tmpl.Execute(response, struct {
			Task   Task
			Users  []User
			Errors map[string]string
		}{
			Task:   task,
			Users:  users,
			Errors: errors,
		})
	}
}

func delTask(w http.ResponseWriter, r *http.Request) {
	var (
		task models.Task
	)
	task.ID, _ := strconv.Atoi(request.FormValue("id"))
	task.DeleteTask()
	http.Redirect(response, request, "/task/list/", http.StatusFound)
}

func modTask() {
	var (
		task modles.Task
		users []models.User
		deadlineTime *time.Time
		errors = make(map[string]string)
	)

	users = models.GetUsers()
	if request.Method == http.MethodGet {
		// 加载模板
	} else if request.Method == http.MethodPost {
		task.ID, _ = strconv.Atoi(request.PostFormValue("id"))
		task.Describe = request.PostFormValue("describe")
		task.Status, _ = strconv.Atoi(request.PostFormValue("status"))
		task.UserID, _ = strconv.Atoi(request.PostFormValue("user_id"))

		// 描述信息不可以大于512
		describeLength := utf8.RuneCountInString(task.Describe)
		if describeLength > 512 {
			errors["describe"] = "任务描述不能超过512个字符"
		}

		// 责任人不能为空
		if userID == 0 {
			errors["userid"] = "必须选择责任人"
		}

		if len(errors) == 0 {
			if models.StatusMap[task.Status] == "正在进行" {
				now := time.Now()
				task.StartTime = &now
			}
			if models.StatusMap[task.Status] == "完成" {
				now := time.Now()
				task.CompleteTime = &now
			}
			if err := task.UpdateTask(db); err == nil {
				http.Redirect(response, request, "/task/list/", http.StatusFound)
			} else {
				errors["subimt"] = "保存数据库失败"
			}
			
		}
		tmpl := template.Must(template.New("edit_task.html").ParseFiles("views/task/edit_task.html"))
		err = tmpl.Execute(response, struct {
			Task   Task
			Users  []User
			Status map[int]string
			Errors map[string]string
		}{
			Task:   task,
			Users:  users,
			Status: models.StatusMap,
			Errors: errors,
		})
	}

}
