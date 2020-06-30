package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
	"todolist/config"
	"todolist/utils"
	"unicode/utf8"
)

var (
	db = utils.MySQL()
)

// func (task *config.Task) StatusText() string {
// 	return config.StatusMap[task.Status]
// }

func TaskQuery(resp http.ResponseWriter, req *http.Request) {
	tasks := make([]config.Task, 0, 20)
	rows, err := db.Query(config.SqlTask)
	if err == nil {
		for rows.Next() {
			var task config.Task
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
			return t.Format(config.DatatimeLayout)
		},
		"status": func(status int) string {
			return config.StatusMap[status]
		},
	}
	tmpl := template.Must(template.New("tpl").Funcs(funcs).ParseFiles("views/tasks.html"))
	tmpl.ExecuteTemplate(resp, "tasks.html", struct {
		Tasks []config.Task
	}{tasks})
}

func TaskAdd(resp http.ResponseWriter, req *http.Request) {
	var (
		task   config.TaskForm
		errors = make(map[string]string)
	)
	//校验
	if req.Method == http.MethodGet {

	} else if req.Method == http.MethodPost {
		name := strings.TrimSpace(req.PostFormValue("name"))
		content := strings.TrimSpace(req.PostFormValue("content"))
		deadlineTime := strings.TrimSpace(req.PostFormValue("deadline_time"))
		// fmt.Println(name, content, deadlineTime)

		task = config.TaskForm{
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
			db.Exec(config.SqlCreate, task.Name, task.Content, task.DeadlineTime)
			http.Redirect(resp, req, "/", http.StatusFound)
		}
	}
	//执行模板
	tmpl := template.Must(template.ParseFiles("views/add.html"))
	tmpl.ExecuteTemplate(resp, "add.html", struct {
		Task   config.TaskForm
		Errors map[string]string
	}{task, errors})
}

func TaskDelete(resp http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	// fmt.Println(id)
	db.Exec(config.SqlDelete, id)
	// fmt.Println(sqlDelete, id)
	http.Redirect(resp, req, "/", http.StatusFound)
}

func TaskEdit(resp http.ResponseWriter, req *http.Request) {
	var (
		task             config.TaskForm
		tempStartTime    string
		tempDeadlineTime string
		errors           = make(map[string]string)
	)
	fmt.Println(req.Method)
	if req.Method == http.MethodGet {
		id := req.FormValue("id")
		// fmt.Println("get:", id)
		row := db.QueryRow(config.SqlGetTask, id)
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
			db.Exec(config.SqlEditTask, name, startTime, deadTime, content, id)
			http.Redirect(resp, req, "/", http.StatusFound)
		}
	}

	tmpl := template.Must(template.ParseFiles("views/edit.html"))
	tmpl.ExecuteTemplate(resp, "edit.html", struct {
		Task   config.TaskForm
		Errors map[string]string
	}{task, errors})
}
