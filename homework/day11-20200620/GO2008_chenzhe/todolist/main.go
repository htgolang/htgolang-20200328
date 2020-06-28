package main

import (
	"net/http"
	"todolist/models"
)



const (
	listenAdd = ":8888"
)




func main() {

	http.Handle("/static/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/", models.Index)
	//user
	http.HandleFunc("/user", models.ListUser)
	http.HandleFunc("/userdelete",models.DeleteUser)
	http.HandleFunc("/useradd",models.AddUser)
	http.HandleFunc("/useredit",models.EditUser)
	//task
	http.HandleFunc("/task",models.ListTask)
	http.HandleFunc("/taskdelete",models.DeleteTask)
	http.HandleFunc("/taskadd",models.AddTask)
	http.ListenAndServe(listenAdd, nil)
}