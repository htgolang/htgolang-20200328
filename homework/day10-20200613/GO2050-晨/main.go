package main

import (

	"html/template"
	"net/http"
	"strconv"
	"todolist/models"
)

func main()  {
	http.HandleFunc("/list/", func(resp http.ResponseWriter, req *http.Request) {
		tpl := template.Must(template.ParseFiles("views/list.html"))
		tpl.ExecuteTemplate(resp,"list.html",models.Taskslist)  //传入数据
	})

	http.HandleFunc("/add/", func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost{
			taskname := req.PostFormValue("taskname")
			status := req.PostFormValue("status")
			models.AddTask(taskname,status)
			http.Redirect(resp,req,"/list/",302)
		}
		tpl := template.Must(template.ParseFiles("views/add.html"))
		tpl.ExecuteTemplate(resp,"add.html",nil)
	})

	http.HandleFunc("/query/", func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			word := req.PostFormValue("query")
			newtask :=models.Search(word)
			tpl := template.Must(template.ParseFiles("views/query_result.html"))
			tpl.ExecuteTemplate(resp,"query_result.html",newtask)
		}
		tpl := template.Must(template.ParseFiles("views/query.html"))
		tpl.ExecuteTemplate(resp,"query.html",nil)
	})

	http.HandleFunc("/delete/", func(resp http.ResponseWriter, req *http.Request) {
		id ,_:= strconv.Atoi(req.FormValue("id"))
		models.Delete(id)
		http.Redirect(resp,req,"/list/",302)
	})

	http.HandleFunc("/modify/", func(resp http.ResponseWriter, req *http.Request) {
		id ,_:= strconv.Atoi(req.FormValue("id"))
		task := models.Retrieve(id)
		if req.Method == http.MethodPost {

			taskname := req.PostFormValue("taskname")
			taskstatus := req.PostFormValue("taskstatus")
			models.Modify(id,taskname,taskstatus)
			http.Redirect(resp,req,"/list/",302)

		}
		tpl := template.Must(template.ParseFiles("views/modify.html"))
		tpl.ExecuteTemplate(resp,"modify.html",task)


	})



	http.ListenAndServe(":8080",nil)
}
