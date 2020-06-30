package models

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"todolist/config"
	"todolist/db"
	"todolist/user"
	"todolist/utils"
)

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

type TaskForm struct {
	ID           int
	Name         string
	Status       int
	StartTime    string
	CompleteTime string
	DeadlineTime string
	Content      string
	User         int
}

func init() {
	// 设置log
	logfile, _ := os.OpenFile("task.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	defer logfile.Close()
	log.SetOutput(logfile)
}

func (t *Task) Index(w http.ResponseWriter, r *http.Request) {
	tasks := make([]Task, 0, 20)
	rows, err := db.Config.DB.Query(config.SqlQueryAllTask)
	if err == nil {
		for rows.Next() {
			var task Task
			var status int
			var description interface{}
			err := rows.Scan(&task.ID, &task.Name, &status, &task.StartTime, &task.CompleteTime, &task.DeadlineTime, &task.User, &description)

			task.Status = config.StatusMap[status]
			if desc, ok := description.([]byte); ok {
				task.Content = string(desc)
			} else {
				task.Content = "--"
			}

			if err == nil {
				tasks = append(tasks, task)
			} else {
				log.Println(err)
			}
		}
	}

	funcs := template.FuncMap{
		"datetime": func(t *time.Time) string {
			if t == nil {
				return "--"
			}
			return t.Format(config.TimeLayout)
		},
		"status": func(status int) string {
			return config.StatusMap[status]
		},
	}

	tpl := template.Must(template.New("tpl").Funcs(funcs).ParseFiles("views/index.html"))
	tpl.ExecuteTemplate(w, "index.html", struct {
		Tasks []Task
	}{tasks})
}

// 添加任务
func (t *Task) AddTask(w http.ResponseWriter, r *http.Request) {
	var (
		errMsg = make(map[string]string)
		task   TaskForm
	)

	// account := user.NewUser()
	users := user.GetAccounts()

	if r.Method == http.MethodPost {
		name := strings.TrimSpace(r.PostFormValue("name"))
		content := strings.TrimSpace(r.PostFormValue("content"))
		startTime := strings.TrimSpace(r.PostFormValue("start_time"))
		completeTime := strings.TrimSpace(r.PostFormValue("complete_time"))
		deadlineTime := strings.TrimSpace(r.PostFormValue("deadline_time"))
		suid := strings.TrimSpace(r.PostFormValue("user"))
		uid, _ := strconv.Atoi(suid)
		// 检查任务名称
		if err := utils.CheckTaskName(name); err != nil {
			errMsg["name"] = err.Error()
			fmt.Println(errMsg)
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

		task = TaskForm{
			Name:         string(name),
			Content:      string(content),
			StartTime:    st,
			DeadlineTime: dt,
			User:         uid,
		}

		if len(errMsg) == 0 {
			if completeTime != "" {
				ct := strings.ReplaceAll(completeTime, "T", " ")
				db.Config.DB.Exec(config.SqlCreateTaskWithCt, task.Name, task.Content, st, ct, dt, uid)
			} else {
				db.Config.DB.Exec(config.SqlCreateTask, task.Name, task.Content, st, dt, uid)
			}
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}

	tpl := template.Must(template.ParseFiles("views/task/add.html"))
	tpl.ExecuteTemplate(w, "add.html", struct {
		Task  TaskForm
		User  []user.User
		Error map[string]string
	}{task, users, errMsg})
}

// 查询任务
func (t *Task) QueryTask(w http.ResponseWriter, r *http.Request) {
	var (
		task        Task
		status      int
		ok          bool
		error       string
		description interface{}
	)

	tid := r.FormValue("id")
	if id, err := strconv.Atoi(tid); err == nil {
		task.ID = id
	}
	row := db.Config.DB.QueryRow(config.SqlQueryTaskWithUserName, tid)
	err := row.Scan(&task.Name, &status, &task.StartTime, &task.CompleteTime, &task.DeadlineTime, &description, &task.User)

	if err == nil {
		// 	检查任务描述
		if description != nil {
			cnt, _ := description.([]byte)
			task.Content = string(cnt)
		}

		task.Status = config.StatusMap[status]
		ok = true
	} else if err == sql.ErrNoRows {
		error = "该任务不存在!"
	} else {
		error = err.Error()
	}

	funcs := template.FuncMap{
		"datetime": func(t *time.Time) string {
			if t == nil {
				return "--"
			}
			return t.Format(config.TimeLayout)
		},
	}

	tpl := template.Must(template.New("tpl").Funcs(funcs).ParseFiles("views/task/query.html"))
	tpl.ExecuteTemplate(w, "query.html", struct {
		Task
		Error string
		OK    bool
	}{task, error, ok})
}

// // 修改任务
func (t *Task) ModifyTask(w http.ResponseWriter, r *http.Request) {
	var (
		task TaskForm
		// account          user.User
		tempStartTime    string
		tempDeadlineTime string
		errMsg           = make(map[string]string)
	)

	// account := user.NewUser()
	users := user.GetAccounts()

	if r.Method == http.MethodGet {
		id := strings.TrimSpace(r.FormValue("id"))
		row := db.Config.DB.QueryRow(config.SqlQueryTask, id)
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

	} else if r.Method == http.MethodPost {
		id := strings.TrimSpace(r.PostFormValue("id"))
		name := strings.TrimSpace(r.PostFormValue("name"))
		status := strings.TrimSpace(r.PostFormValue("status"))
		startTime := strings.TrimSpace(r.PostFormValue("start_time"))
		deadlineTime := strings.TrimSpace(r.PostFormValue("deadline_time"))
		content := strings.TrimSpace(r.PostFormValue("content"))
		uid := strings.TrimSpace(r.PostFormValue("user"))

		// 检查任务名称
		if err := utils.CheckTaskName(name); err != nil {
			errMsg["name"] = err.Error()
		}

		// 检查截止日期
		dt := strings.ReplaceAll(deadlineTime, "T", " ")
		if err := utils.CheckDeadline(dt); err != nil {
			errMsg["deadlineTime"] = deadlineTime
		}

		// 检查任务描述
		if err := utils.CheckContent(content); err != nil {
			errMsg["content"] = content
		}

		if status == "3" {
			completeTime := time.Now().Format(config.TimeLayout)
			db.Config.DB.Exec(config.SqlUpdateTaskIncludeTime, name, status, startTime, deadlineTime, completeTime, content, uid, id)
		} else {
			db.Config.DB.Exec(config.SqlUpdateTask, name, status, startTime, deadlineTime, content, uid, id)
		}
		http.Redirect(w, r, "/", http.StatusFound)

	}
	tpl := template.Must(template.ParseFiles("views/task/modify.html"))
	tpl.ExecuteTemplate(w, "modify.html", struct {
		Task  TaskForm
		User  []user.User
		Error map[string]string
	}{task, users, errMsg})
}

// 删除任务
func (t *Task) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	db.Config.DB.Exec(config.SqlDeleteTask, id)
	http.Redirect(w, r, "/", http.StatusFound)
}
