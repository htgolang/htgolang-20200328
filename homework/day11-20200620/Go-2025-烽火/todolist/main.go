package main

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
	"todolist/user"
	"todolist/utils"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	dateTimeLayout = "2006-01-02 15:04"
// )

const (
	dbDriver   = "mysql"
	dbUser     = "devon"
	dbPassword = "golang@2020"
	dbName     = "todolist"
	dbHost     = "127.0.0.1"
	dbPort     = 32769
)

const (
	sqlCreateTask            = `insert into task(name, content, start_time, deadline_time, user) values(?,?,?,?,?)`
	sqlCreateTaskWithCt      = `insert into task(name, content, start_time, end_time, deadline_time, user) values(?,?,?,?,?,?)`
	sqlQueryAllTask          = "select task.id, task.name, task.status, start_time, end_time, deadline_time, user.name, content from task left join user on task.user=user.id"
	sqlDeleteTask            = `delete from task where id=?`
	sqlQueryTask             = `select id, name, status, start_time, deadline_time, content, user from task where id=?`
	sqlQueryTaskWithName     = `select task.id, task.name, task.status, start_time, end_time, deadline_time, content, user.name from task  left join user on task.user=user.id where task.id=?`
	sqlUpdateTaskIncludeTime = `update task set name=?, status=?, start_time=?, deadline_time=?, end_time=?, content=?, user=? where id=?`
	sqlUpdateTask            = `update task set name=?, status=?, start_time=?, deadline_time=?, content=?,user=? where id=?`
	sqlQueryAllUser          = `select * from user`
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

// 命令行的任务管理器
func main() {
	// 设置log
	logfile, _ := os.OpenFile("task.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	defer logfile.Close()
	log.SetOutput(logfile)
	addr := ":9000"

	// 数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	// 测试数据库连接
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	// css image
	http.Handle("/static/", http.FileServer(http.Dir("./views/")))

	// homepage
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tasks := make([]Task, 0, 20)
		rows, err := db.Query(sqlQueryAllTask)
		if err == nil {
			for rows.Next() {
				var t Task
				var status int
				var description interface{}
				err := rows.Scan(&t.ID, &t.Name, &status, &t.StartTime, &t.CompleteTime, &t.DeadlineTime, &t.User, &description)
				t.Status = statusMap[status]
				if desc, ok := description.([]byte); ok {
					t.Content = string(desc)
				} else {
					t.Content = "--"
				}

				if err == nil {
					tasks = append(tasks, t)
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
				return t.Format(utils.TimeLayout)
			},
			"status": func(status int) string {
				return statusMap[status]
			},
		}

		tpl := template.Must(template.New("tpl").Funcs(funcs).ParseFiles("views/index.html"))
		tpl.ExecuteTemplate(writer, "index.html", struct {
			Tasks []Task
		}{tasks})
	})

	// 添加任务
	http.HandleFunc("/add/", func(writer http.ResponseWriter, request *http.Request) {
		var (
			task   TaskForm
			errors map[string]string
		)

		account := user.NewUser()
		users := account.GetAccounts(db)

		if request.Method == http.MethodPost {
			name := strings.TrimSpace(request.PostFormValue("name"))
			content := strings.TrimSpace(request.PostFormValue("content"))
			startTime := strings.TrimSpace(request.PostFormValue("start_time"))
			completeTime := strings.TrimSpace(request.PostFormValue("complete_time"))
			deadlineTime := strings.TrimSpace(request.PostFormValue("deadline_time"))
			suid := strings.TrimSpace(request.PostFormValue("user"))
			uid, _ := strconv.Atoi(suid)

			// 检查任务名称
			if rt, ok := utils.CheckTaskName(name); !ok {
				errors["name"] = rt
			}
			// 检查日期
			st := strings.ReplaceAll(startTime, "T", " ")

			fmt.Printf("uid: %#v\n", suid)
			dt := strings.ReplaceAll(deadlineTime, "T", " ")
			if rt, ok := utils.CheckDeadline(dt); !ok {
				errors["deadline_time"] = rt
			}

			// 检查任务描述
			if rt, ok := utils.CheckContent(content); !ok {
				errors["content"] = rt
			}

			task = TaskForm{
				Name:      string(name),
				Content:   string(content),
				StartTime: st,
				// CompleteTime: ct,
				DeadlineTime: dt,
				User:         uid,
			}

			if len(errors) == 0 {
				if completeTime != "" {
					ct := strings.ReplaceAll(completeTime, "T", " ")
					db.Exec(sqlCreateTaskWithCt, task.Name, task.Content, st, ct, dt, uid)
				} else {
					db.Exec(sqlCreateTask, task.Name, task.Content, st, dt, uid)
				}
				http.Redirect(writer, request, "/", http.StatusFound)
			}
		}

		tpl := template.Must(template.ParseFiles("views/task/add.html"))
		tpl.ExecuteTemplate(writer, "add.html", struct {
			Task   TaskForm
			User   []user.User
			Errors map[string]string
		}{task, users, errors})
	})

	http.HandleFunc("/query/", func(writer http.ResponseWriter, request *http.Request) {
		var (
			task        Task
			status      int
			ok          bool
			error       string
			description interface{}
		)

		tid := request.FormValue("id")
		row := db.QueryRow(sqlQueryTaskWithName, tid)
		err := row.Scan(&task.ID, &task.Name, &status, &task.StartTime, &task.CompleteTime, &task.DeadlineTime, &description, &task.User)
		if err == nil {
			// 	检查任务描述
			if description != nil {
				cnt, _ := description.([]byte)
				task.Content = string(cnt)
			}

			task.Status = statusMap[status]
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
				return t.Format(utils.TimeLayout)
			},
		}

		tpl := template.Must(template.New("tpl").Funcs(funcs).ParseFiles("views/task/query.html"))
		tpl.ExecuteTemplate(writer, "query.html", struct {
			T     Task
			Error string
			OK    bool
		}{task, error, ok})
	})

	// 修改任务
	http.HandleFunc("/modify/", func(writer http.ResponseWriter, request *http.Request) {
		var (
			task             TaskForm
			tempStartTime    string
			tempDeadlineTime string
			// username         string
			errors  map[string]string
			content interface{}
		)

		account := user.NewUser()
		users := account.GetAccounts(db)

		if request.Method == http.MethodGet {
			request.ParseForm()
			id := request.FormValue("id")
			row := db.QueryRow(sqlQueryTask, id)
			err := row.Scan(&task.ID, &task.Name, &task.Status, &tempStartTime, &tempDeadlineTime, &content, &task.User)
			// err := row.Scan(&task.ID, &task.Name, &task.Status, &tempStartTime, &tempDeadlineTime, &content, &username)
			if err != nil {
				log.Fatal(err)
			}
			if content != nil {
				cnt, _ := content.([]byte)
				task.Content = string(cnt)
			}

			if tempStartTime != "" {
				startTime := strings.ReplaceAll(tempStartTime, ":00Z", "")
				task.StartTime = startTime
			}

			deadlineTime := strings.ReplaceAll(tempDeadlineTime, ":00Z", "")
			task.DeadlineTime = deadlineTime
			// fmt.Printf("%#v\n", task)

		} else if request.Method == http.MethodPost {
			id := strings.TrimSpace(request.PostFormValue("id"))
			name := strings.TrimSpace(request.PostFormValue("name"))
			status := strings.TrimSpace(request.PostFormValue("status"))
			startTime := strings.TrimSpace(request.PostFormValue("start_time"))
			deadlineTime := strings.TrimSpace(request.PostFormValue("deadline_time"))
			content := strings.TrimSpace(request.PostFormValue("content"))
			uid := strings.TrimSpace(request.PostFormValue("user"))

			// 检查任务名称
			if rt, ok := utils.CheckTaskName(name); !ok {
				errors["name"] = rt
			}

			// 检查截止日期
			dt := strings.ReplaceAll(deadlineTime, "T", " ")
			if rt, ok := utils.CheckDeadline(dt); !ok {
				errors["deadline_time"] = rt
			}

			// 检查任务描述
			if rt, ok := utils.CheckContent(content); !ok {
				errors["content"] = rt
			}

			if status == "3" {
				completeTime := time.Now().Format(utils.TimeLayout)
				db.Exec(sqlUpdateTaskIncludeTime, name, status, startTime, deadlineTime, completeTime, content, uid, id)
			} else {
				db.Exec(sqlUpdateTask, name, status, startTime, deadlineTime, content, uid, id)
				// sqlUpdateTask            = `update task set name=?, status=?, start_time=?, deadline_time=?, content=? user=? where id=?`

			}
			http.Redirect(writer, request, "/", http.StatusFound)

		}
		tpl := template.Must(template.ParseFiles("views/task/modify.html"))
		tpl.ExecuteTemplate(writer, "modify.html", struct {
			Task   TaskForm
			User   []user.User
			Errors map[string]string
		}{task, users, errors})
	})

	// 删除任务
	http.HandleFunc("/delete/", func(writer http.ResponseWriter, request *http.Request) {
		id := request.FormValue("id")
		db.Exec(sqlDeleteTask, id)
		http.Redirect(writer, request, "/", http.StatusFound)
	})

	http.ListenAndServe(addr, nil)

}
