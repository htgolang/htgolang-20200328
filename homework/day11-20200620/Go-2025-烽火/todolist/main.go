package main

import (
	"net/http"
	"todolist/models"
)

// 命令行的任务管理器
func main() {

	addr := ":9000"

	// css image
	http.Handle("/static/", http.FileServer(http.Dir("./views/")))

	var task models.Task
	// homepage
	http.HandleFunc("/", task.Index)

	http.HandleFunc("/add/", task.AddTask)
	http.HandleFunc("/delete/", task.DeleteTask)
	http.HandleFunc("/modify/", task.ModifyTask)
	http.HandleFunc("/query/", task.QueryTask)

	http.ListenAndServe(addr, nil)

}
