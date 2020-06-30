package main

import (
	"net/http"
	"todolist/models"
	"todolist/user"
)

// 命令行的任务管理器
func main() {

	addr := ":9000"

	// css image
	http.Handle("/static/", http.FileServer(http.Dir("./views/")))

	var task models.Task
	var user user.User
	// homepage
	http.HandleFunc("/", task.Index)

	// task
	http.HandleFunc("/task/add/", task.AddTask)
	http.HandleFunc("/task/delete/", task.DeleteTask)
	http.HandleFunc("/task/modify/", task.ModifyTask)
	http.HandleFunc("/task/query/", task.QueryTask)

	// user
	http.HandleFunc("/user/add/", user.AddUser)
	http.HandleFunc("/user/query/", user.QueryUser)

	http.ListenAndServe(addr, nil)

}
