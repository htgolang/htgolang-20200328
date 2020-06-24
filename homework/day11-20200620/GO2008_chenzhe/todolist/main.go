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
	http.HandleFunc("/user", models.ListUser)
	http.HandleFunc("/userdelete",models.DeleteUser)
	http.HandleFunc("/useradd",models.AddUser)
	http.HandleFunc("/useredit",models.EditUser)
	http.ListenAndServe(listenAdd, nil)
}