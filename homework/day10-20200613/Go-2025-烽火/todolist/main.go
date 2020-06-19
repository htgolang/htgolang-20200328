package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"todolist/task"
	"todolist/user"

	"github.com/gorilla/sessions"
)

// 命令行的任务管理器v8版本
var Task = new(task.TaskController)
var key = []byte("user-session")
var store = sessions.NewCookieStore(key)

func main() {
	// 设置log
	logfile, _ := os.OpenFile("task.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	defer logfile.Close()
	log.SetOutput(logfile)
	addr := ":9000"

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		session, _ := store.Get(request, "cookie-name")
		auth := session.Values["authenticated"]
		if auth != "ok" {
			http.Redirect(writer, request, "/login/", 302)
			// http.Error(writer, "Forbidden", http.StatusForbidden)
		} else {
			tpl := template.Must(template.ParseFiles("views/index.html"))
			tpl.ExecuteTemplate(writer, "index.html", nil)
		}
	})

	// css
	http.Handle("/css/", http.FileServer(http.Dir("./views/")))
	// image
	http.Handle("/image/", http.FileServer(http.Dir("./views/")))

	// 登陆页
	http.HandleFunc("/login/", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		if request.Method == http.MethodPost {
			username := request.PostForm.Get("username")
			password := request.PostForm.Get("password")
			if username != "" && password != "" {
				rt := user.Validate(username, password)
				if rt == true {
					session, _ := store.Get(request, "cookie-name")
					session.Values["authenticated"] = "ok"
					session.Save(request, writer)

					http.Redirect(writer, request, "/", 301)
				}
			}
		}
		tpl := template.Must(template.ParseFiles("views/login.html"))
		tpl.ExecuteTemplate(writer, "login.html", nil)
	})

	// 查询任务
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		tpl := template.Must(template.ParseFiles("views/list.html"))
		name := request.Form.Get("name")
		tpl.ExecuteTemplate(writer, "list.html", task.ListTask(name))
	})

	// 添加任务
	http.HandleFunc("/add/", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		if request.Method == http.MethodPost {
			name := request.PostForm.Get("name")
			user := request.PostForm.Get("user")
			startTime := request.PostForm.Get("startTime")
			endTime := request.PostForm.Get("endTime")
			status := request.PostForm.Get("status")
			Task.AddTask(name, startTime, endTime, status, user)
			http.Redirect(writer, request, "/list/", 302)
		}
		tpl := template.Must(template.ParseFiles("views/add.html"))
		tpl.ExecuteTemplate(writer, "add.html", nil)
	})

	// 修改任务
	http.HandleFunc("/modify/", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		if request.Method == http.MethodPost {
			name := request.PostForm.Get("name")
			status := request.PostForm.Get("status")
			user := request.PostForm.Get("user")
			Task.Modify(name, status, user)
			http.Redirect(writer, request, "/list/", 302)
		}
		tpl := template.Must(template.ParseFiles("views/modify.html"))
		tpl.ExecuteTemplate(writer, "modify.html", nil)
	})

	// 删除任务
	http.HandleFunc("/delete/", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		if request.Method == http.MethodPost {
			name := request.PostForm.Get("name")
			Task.Delete(name)
			http.Redirect(writer, request, "/list/", 302)
		}
		tpl := template.Must(template.ParseFiles("views/delete.html"))
		tpl.ExecuteTemplate(writer, "delete.html", nil)
	})

	http.ListenAndServe(addr, nil)

}
