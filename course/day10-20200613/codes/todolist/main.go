package main

import (
	"html/template"
	"net/http"

	"todolist/models"
)

func main() {
	addr := ":9999"
	http.HandleFunc("/list/", func(response http.ResponseWriter, request *http.Request) {
		tpl := template.Must(template.ParseFiles("views/list.html"))
		tpl.ExecuteTemplate(response, "list.html", models.GetTasks())
	})

	http.HandleFunc("/add/", func(response http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			name := request.PostFormValue("name")
			models.AddTask(name)
			http.Redirect(response, request, "/list/", 302)
		}
		tpl := template.Must(template.ParseFiles("views/add.html"))
		tpl.ExecuteTemplate(response, "add.html", nil)
	})

	http.ListenAndServe(addr, nil)
}
