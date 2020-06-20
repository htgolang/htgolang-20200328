package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

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
	sqlTasks = "select task.id, task.name, task.status, task.start_time, task.complete_time, task.deadline_time, user.name as user from task left join user on task.user=user.id"
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

	http.ListenAndServe(listenAdd, nil)
}
