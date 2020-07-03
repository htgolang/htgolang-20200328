package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
	"unicode/utf8"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbDriver   = "mysql"
	dbUser     = "root"
	dbPasswd   = "123456"
	dbHost     = "192.168.163.200"
	dbPort     = 3306
	dbName     = "test"
	listenAddr = ":8080"
	statusMap  = map[int]string{0: "新建", 1: "正在进行", 2: "暂停", 3: "完成"}
)

const (
	dateTimeLayout = "2006-01-02 15:04:05"
)

type Task struct {
	ID           int
	Name         string
	Status       int
	StartTime    *time.Time
	CompleteTime *time.Time
	DeadlineTime *time.Time
	UserID       int
	Describe     string
}

func (t Task) insertTask(db *sql.DB) error {
	const (
		insertTaskSQL = "INSERT INTO task (`name`, `status`, `userid`, `starttime`, `completetime`, `deadlinetime`, `describe`) VALUE(?, ?, ?, ?, ?, ?, ?)"
	)
	_, err := db.Exec(insertTaskSQL, t.Name, t.Status, t.UserID, t.StartTime, t.CompleteTime, t.DeadlineTime, t.Describe)
	return err
}

func (t Task) updateTask(db *sql.DB) error {
	const (
		updateTaskSQL = "UPDATE task SET `name`=?, `status`=?, `userid`=?, `starttime`=?, `completetime`=?, `deadlinetime`=?, `describe`=? WHERE id=?"
	)
	result, err := db.Exec(updateTaskSQL, t.Name, t.Status, t.UserID, t.StartTime, t.CompleteTime, t.DeadlineTime, t.Describe, t.ID)
	fmt.Println(result.RowsAffected())
	return err
}

func (t Task) deleteTask(db *sql.DB) error {
	const (
		deleteTaskSQL = "DELETE FROM task WHERE id = ?"
	)
	_, err := db.Exec(deleteTaskSQL, t.ID)
	return err
}

func (t *Task) getTaskByID(db *sql.DB) {
	const (
		selectTaskSQL = "SELECT `id`, `name`, `status`, `userid`, `starttime`, `completetime`, `deadlinetime`, `describe` FROM task WHERE `id` = ?"
	)
	row := db.QueryRow(selectTaskSQL, t.ID)
	row.Scan(&t.ID, &t.Name, &t.Status, &t.UserID, &t.StartTime, &t.CompleteTime, &t.DeadlineTime, &t.Describe)
}

type User struct {
	ID           int
	Name         string
	Account      string
	Tel          int64
	Passwd       string
	Address      string
	RegisterTime *time.Time
}

func getUser(db *sql.DB) ([]User, error) {
	const (
		sqlUser = "select id, account, name, tel, passwd, registertime, address from user"
	)

	var (
		userRows *sql.Rows
		err      error
	)
	users := make([]User, 0, 20)

	// 获取user数据
	userRows, err = db.Query(sqlUser)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for userRows.Next() {
		var user User
		err = userRows.Scan(&user.ID, &user.Account, &user.Name, &user.Tel, &user.Passwd, &user.RegisterTime, &user.Address)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func getTask(db *sql.DB) ([]Task, error) {
	const (
		sqlTasks = "select `id`, `name`, `userid`, `status`, `deadlinetime`, `starttime`, `completetime`, `describe` from task;"
	)

	var (
		taskRows *sql.Rows
		err      error
	)
	tasks := make([]Task, 0, 20)

	// 获取task数据
	taskRows, err = db.Query(sqlTasks)
	if err != nil {
		return nil, err
	}
	for taskRows.Next() {
		var task Task
		err = taskRows.Scan(&task.ID, &task.Name, &task.UserID, &task.Status, &task.DeadlineTime, &task.StartTime, &task.CompleteTime, &task.Describe)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func main() {
	var (
		dsn string
		err error
		db  *sql.DB

		tasks []Task
		users []User
	)
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=PRC&parseTime=true", dbUser, dbPasswd, dbHost, dbPort, dbName)

	// 打开数据库连接池
	db, err = sql.Open(dbDriver, dsn)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 添加js和css
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("views/css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("views/js/"))))

	// 路由
	http.HandleFunc("/", func(response http.ResponseWriter, requests *http.Request) {
		// 获取tasks
		tasks, err = getTask(db)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 获取users
		users, err = getUser(db)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(tasks)
		funcMap := template.FuncMap{
			"datetime": func(t *time.Time) string {
				if t == nil {
					return "--"
				} else {
					return t.Format(dateTimeLayout)
				}
			},
			"status": func(statusCode int) string {
				return statusMap[statusCode]
			},
			"headleuser": func(userID int) string {
				for _, user := range users {
					if user.ID == userID {
						return user.Name
					}
				}
				return "--"
			},
		}
		tmpl := template.Must(template.New("task.html").Funcs(funcMap).ParseFiles("views/task.html"))
		err = tmpl.Execute(response, struct {
			Tasks []Task
		}{
			Tasks: tasks,
		})
		if err != nil {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/add_task/", func(response http.ResponseWriter, request *http.Request) {
		var (
			task   Task
			errors = make(map[string]string)
		)

		// 获取users
		users, err = getUser(db)
		if err != nil {
			fmt.Println(err)
			return
		}
		if request.Method == http.MethodGet {
			// 加载模板
		} else if request.Method == http.MethodPost {
			// 数据验证
			// 成功  --- 跳转到列表
			// 失败  --- 当前页面，数据回显，错误提醒
			var deadlineTime *time.Time
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

			task = Task{
				Name:         taskName,
				DeadlineTime: deadlineTime,
				UserID:       userID,
				Describe:     describe,
				Status:       0,
			}
			// 当验证无错误，将数据写入数据库中，并跳转回显
			if len(errors) == 0 {
				task.insertTask(db)
				http.Redirect(response, request, "/", http.StatusFound)
			}

		}

		funcMap := template.FuncMap{
			"datetime": func(t *time.Time) string {
				if t == nil {
					return ""
				} else {
					return t.Format(dateTimeLayout)
				}
			},
		}

		tmpl := template.Must(template.New("add_task.html").Funcs(funcMap).ParseFiles("views/add_task.html"))
		err = tmpl.Execute(response, struct {
			Task   Task
			Users  []User
			Errors map[string]string
		}{
			Task:   task,
			Users:  users,
			Errors: errors,
		})
	})

	http.HandleFunc("/delete_task/", func(response http.ResponseWriter, request *http.Request) {
		task_id, _ := strconv.Atoi(request.FormValue("id"))
		task := Task{ID: task_id}
		task.deleteTask(db)
		http.Redirect(response, request, "/", http.StatusFound)
	})

	http.HandleFunc("/edit_task/", func(response http.ResponseWriter, request *http.Request) {
		var (
			task   Task
			errors = make(map[string]string)
		)

		// 获取users
		users, err = getUser(db)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 获取任务id
		task_id, _ := strconv.Atoi(request.FormValue("id"))
		task = Task{ID: task_id}

		//生成任务数据
		task.getTaskByID(db)

		if request.Method == http.MethodGet {

		} else if request.Method == http.MethodPost && statusMap[task.Status] == "完成" {
			errors["subime"] = "任务已经完成无法修改"
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
			if task.UserID == 0 {
				errors["userid"] = "必须选择责任人"
			}

			if len(errors) == 0 {
				if statusMap[task.Status] == "正在进行" {
					now := time.Now()
					task.StartTime = &now
				}
				if statusMap[task.Status] == "完成" {
					now := time.Now()
					task.CompleteTime = &now
				}
				fmt.Println(task)
				fmt.Println(task.updateTask(db))
				http.Redirect(response, request, "/", http.StatusFound)
			}
		}

		tmpl := template.Must(template.New("edit_task.html").ParseFiles("views/edit_task.html"))
		err = tmpl.Execute(response, struct {
			Task   Task
			Users  []User
			Status map[int]string
			Errors map[string]string
		}{
			Task:   task,
			Users:  users,
			Status: statusMap,
			Errors: errors,
		})
	})

	http.ListenAndServe(listenAddr, nil)
}
