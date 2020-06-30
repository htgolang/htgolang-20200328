package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	_ "github.com/go-sql-driver/mysql"
)

//连接数据库信息
const (
	dbDriver = "mysql"
	dbUser   = "gostudy"
	dbPasswd = "123456q!"
	dbName   = "gostudy"
	dbHost   = "120.79.60.117"
	dbPort   = 3306
)

//web服务端口
const (
	listenAdd = ":8989"
)

//查询数据库SQL
const (
	sqlTask     = "select task.id,task.name,task.status,task.start_time,task.complete_time,task.deadline_time,user.name as user  from task left join user on task.user=user.id"
	sqlCreate   = "insert into task(name,content,deadline_time) value(?,?,?)"
	sqlDelete   = "delete from task where id = ?"
	sqlEditTask = "update task set name=?, start_time=?, deadline_time=?, content=? where id =?"
	sqlGetTask  = "select id, name , content, start_time, complete_time, deadline_time,user from task where id = ?"
)

//时间格式
const (
	datatimeLayout = "2006-01-02 15:04:05"
)

//任务状态
var (
	statusMap = map[int]string{
		0: "新建",
		1: "正在进行",
		2: "暂停",
		3: "完成",
	}
)

//实体化task
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

//返回任务状态对应的content
func (task *Task) StatusText() string {
	return statusMap[task.Status]
}

//返回到页面的task
type TaskForm struct {
	ID           int
	Name         string
	Status       int
	StartTime    string
	DeadlineTime string
	Content      string
	User         int
}

func main() {

	//连接数据库库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=PRC", dbUser, dbPasswd, dbHost, dbPort, dbName)
	db, err := sql.Open(dbDriver, dsn)

	if err != nil {
		log.Fatal(err)
	}
	//判断数据库连通性
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	//显示query
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		tasks := make([]Task, 0, 20)
		rows, err := db.Query(sqlTask)
		if err == nil {
			for rows.Next() {
				var task Task
				//				var status int转化  &task.Status
				err := rows.Scan(&task.ID, &task.Name, &task.Status, &task.StartTime, &task.CompleteTime, &task.DeadlineTime, &task.User)
				// task.Status = statusMap[status]
				if err == nil {
					tasks = append(tasks, task)
				} else {
					fmt.Println(err)
				}
			}
		}
		// fmt.Println(tasks)
		funcs := template.FuncMap{
			"datetime": func(t *time.Time) string {
				if t == nil {
					return "--"
				}
				return t.Format(datatimeLayout)
			},
			"status": func(status int) string {
				return statusMap[status]
			},
		}
		tmpl := template.Must(template.New("tpl").Funcs(funcs).ParseFiles("views/tasks.html"))
		tmpl.ExecuteTemplate(resp, "tasks.html", struct {
			Tasks []Task
		}{tasks})
	})

	//增加
	http.HandleFunc("/add/", func(resp http.ResponseWriter, req *http.Request) {
		var (
			task   TaskForm
			errors = make(map[string]string)
		)
		//校验
		if req.Method == http.MethodGet {

		} else if req.Method == http.MethodPost {
			name := strings.TrimSpace(req.PostFormValue("name"))
			content := strings.TrimSpace(req.PostFormValue("content"))
			deadlineTime := strings.TrimSpace(req.PostFormValue("deadline_time"))
			fmt.Println(name, content, deadlineTime)

			task = TaskForm{
				Name:         name,
				Content:      content,
				DeadlineTime: deadlineTime,
			}

			//name长度
			nameLength := utf8.RuneCountInString(task.Name)
			if nameLength == 0 {
				errors["name"] = "name not allow null"
			} else if nameLength > 32 {
				errors["name"] = "name not allow more then 32"
			}
			//描述长度校验
			contentLength := utf8.RuneCountInString(task.Content)
			if contentLength > 512 {
				errors["content"] = "content not allow more then 512"
			}
			//时间校验
			if _, err := time.Parse("2006-01-02", deadlineTime); err != nil {
				errors["deadline_time"] = "deadline_time not allow  null"
			}
			//通过提交
			if len(errors) == 0 {
				db.Exec(sqlCreate, task.Name, task.Content, task.DeadlineTime)
				http.Redirect(resp, req, "/", http.StatusFound)
			}
		}
		//执行模板
		tmpl := template.Must(template.ParseFiles("views/add.html"))
		tmpl.ExecuteTemplate(resp, "add.html", struct {
			Task   TaskForm
			Errors map[string]string
		}{task, errors})
	})

	//删除
	http.HandleFunc("/delete/", func(resp http.ResponseWriter, req *http.Request) {
		id := req.FormValue("id")
		fmt.Println(id)
		db.Exec(sqlDelete, id)
		// fmt.Println(sqlDelete, id)
		http.Redirect(resp, req, "/", http.StatusFound)
	})

	//编辑
	http.HandleFunc("/edit/", func(resp http.ResponseWriter, req *http.Request) {
		var (
			task             TaskForm
			tempStartTime    string
			tempDeadlineTime string
			errors           = make(map[string]string)
		)
		fmt.Println(req.Method)
		if req.Method == http.MethodGet {
			id := req.FormValue("id")
			fmt.Println("get:", id)
			row := db.QueryRow(sqlGetTask, id)
			err := row.Scan(&task.ID, &task.Name, &task.Status, &tempStartTime, &tempDeadlineTime, &task.Content, &task.User)
			if err != nil {
				fmt.Println(err)
			}
			if tempStartTime != "" {
				startTime := strings.ReplaceAll(tempStartTime, ":00Z", "")
				task.StartTime = startTime
			}
			deadTime := strings.ReplaceAll(tempDeadlineTime, ":00Z", "")
			task.DeadlineTime = deadTime
		} else if req.Method == http.MethodPost {
			name := strings.TrimSpace(req.FormValue("name"))
			content := strings.TrimSpace(req.FormValue("content"))
			startTime := strings.TrimSpace(req.FormValue("start_time"))
			deadTime := strings.TrimSpace(req.FormValue("deadline_time"))
			id, _ := strconv.Atoi(req.PostFormValue("id"))

			//检测name
			nameLength := utf8.RuneCountInString(name)
			if nameLength == 0 {
				errors["name"] = "name is not null"
			}
			if nameLength > 32 {
				errors["name"] = "name is not allow more then 32"
			}
			//检测任务描述
			contentLen := utf8.RuneCountInString(content)
			if contentLen == 0 {
				errors["content"] = "content is not null"
			} else if contentLen > 512 {
				errors["content"] = "content is not allow more then 32"
			}

			//检测时间
			if _, err := time.Parse("2006-01-02", startTime); err != nil {
				errors["start_time"] = "start_time not allow  null"
			}
			if _, err := time.Parse("2006-01-02", deadTime); err != nil {
				errors["deadline_time"] = "deadline_time not allow  null"
			}

			if len(errors) == 0 {
				db.Exec(sqlEditTask, name, startTime, deadTime, content, id)
				http.Redirect(resp, req, "/", http.StatusFound)
			}
		}

		tmpl := template.Must(template.ParseFiles("views/edit.html"))
		tmpl.ExecuteTemplate(resp, "edit.html", struct {
			Task   TaskForm
			Errors map[string]string
		}{task, errors})
	})

	//启动监听
	http.ListenAndServe(listenAdd, nil)
}
