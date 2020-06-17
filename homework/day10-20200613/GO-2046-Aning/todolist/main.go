package main

import (
	"net/http"
	"strconv"
	"text/template"
	"todolist/models"
)

func main() {
	addr := ":9999"
	//展示列表
	http.HandleFunc("/list/", func(response http.ResponseWriter, request *http.Request) {
		tpl := template.Must(template.ParseFiles("views/list.html"))
		tpl.ExecuteTemplate(response, "list.html", models.GetTasks())
	})

	//新增
	http.HandleFunc("/add/", func(response http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			name := request.PostFormValue("name")
			models.AddTask(name)
			http.Redirect(response, request, "/list/", 302)
		}
		tpl := template.Must(template.ParseFiles("views/add.html"))
		tpl.ExecuteTemplate(response, "add.html", nil)
	})

	//删除
	http.HandleFunc("/del/", func(response http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			id, _ := strconv.Atoi(request.PostFormValue("id"))
			models.DelTask(id)
			http.Redirect(response, request, "/list/", 302)
		}
		tpl := template.Must(template.ParseFiles("views/del.html"))
		tpl.ExecuteTemplate(response, "del.html", nil)
	})

	//监听服务
	http.ListenAndServe(addr, nil)
}