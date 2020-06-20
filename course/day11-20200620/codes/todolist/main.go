package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver   = "mysql"
	dbUser     = "golang"
	dbPassword = "golang@2020"
	dbName     = "todolist"
	dbHost     = "localhost"
	dbPort     = 3306
)

const (
	listenAdd = ":8888"
)

const (
	sqlTasks      = "select task.id, task.name, task.status, task.start_time, task.complete_time, task.deadline_time, user.name as user from task left join user on task.user=user.id"
	sqlCreateTask = "insert into task(name, content, deadline_time) values(?, ?, ?)"
	sqlDeleteTask = "delete from task where id = ?"
)

const (
	dateTimeLayout = "2006-01-02 15:04:05"
)

var (
	statusMap = map[int]string{
		0: "新建",
		1: "正在进行",
		2: "暂停",
		3: "完成",
	}
)

type Task struct {
	ID           int
	Name         string
	Status       int
	StartTime    *time.Time
	CompleteTime *time.Time
	DeadlineTime *time.Time
	User         *string
	Content      string
}

func (task *Task) StatusText() string {
	return statusMap[task.Status]
}

type TaskForm struct {
	ID           int
	Name         string
	Status       int
	DeadlineTime string
	Content      string
	User         int
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=PRC", dbUser, dbPassword, dbHost, dbPort, dbName)

	// 打开数据库连接池
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	// 测试数据库连接
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// 显示任务列表
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		tasks := make([]Task, 0, 20)
		rows, err := db.Query(sqlTasks)
		if err == nil {
			for rows.Next() {
				var task Task
				err := rows.Scan(&task.ID, &task.Name, &task.Status, &task.StartTime, &task.CompleteTime, &task.DeadlineTime, &task.User)
				if err == nil {
					tasks = append(tasks, task)
				} else {
					fmt.Println(err)
				}
			}
		}
		funcs := template.FuncMap{
			"datetime": func(t *time.Time) string {
				if t == nil {
					return "--"
				}
				return t.Format(dateTimeLayout)
			},
			"status": func(status int) string {
				//status int => string
				//if
				//switch
				return statusMap[status]
			},
		}
		// 模板函数必须在解析模板之前进行设置
		tmpl := template.Must(template.New("tpl").Funcs(funcs).ParseFiles("views/tasks.html"))
		tmpl.ExecuteTemplate(response, "tasks.html", struct {
			Tasks []Task
		}{tasks})
	})

	http.HandleFunc("/add/", func(response http.ResponseWriter, request *http.Request) {
		var (
			task   TaskForm
			errors = make(map[string]string)
		)
		if request.Method == http.MethodGet {
			// 加载模板
		} else if request.Method == http.MethodPost {
			name := strings.TrimSpace(request.PostFormValue("name"))
			content := strings.TrimSpace(request.PostFormValue("content"))
			deadlineTime := strings.TrimSpace(request.PostFormValue("deadline_time"))
			task = TaskForm{
				Name:         name,
				Content:      content,
				DeadlineTime: deadlineTime,
			}
			nameLength := utf8.RuneCountInString(task.Name)
			if nameLength == 0 {
				errors["name"] = "任务名不能空"
			} else if nameLength > 32 {
				errors["name"] = "任务名不能超过32个字符"
			}

			contentLength := utf8.RuneCountInString(task.Content)
			if contentLength > 512 {
				errors["content"] = "任务描述不能超过512个字符"
			}

			if _, err := time.Parse("2006-01-02", deadlineTime); err != nil {
				errors["deadline_time"] = "任务期限不能为空"
			}

			// 验证完成，无错误
			if len(errors) == 0 {
				db.Exec(sqlCreateTask, task.Name, task.Content, task.DeadlineTime)
				http.Redirect(response, request, "/", http.StatusFound)
			}

		}
		tmpl := template.Must(template.ParseFiles("views/add_task.html"))
		tmpl.ExecuteTemplate(response, "add_task.html", struct {
			Task   TaskForm
			Errors map[string]string
		}{task, errors})

	})

	http.HandleFunc("/delete/", func(response http.ResponseWriter, request *http.Request) {
		id := request.FormValue("id")
		db.Exec(sqlDeleteTask, id)
		http.Redirect(response, request, "/", http.StatusFound)
	})

	http.ListenAndServe(listenAdd, nil)
}
