package main

import (
	"net/http"
	"todolist/config"
	"todolist/controller"
)

func main() {

	http.HandleFunc("/", controller.TaskQuery)

	http.HandleFunc("/add/", controller.TaskAdd)

	http.HandleFunc("/delete/", controller.TaskDelete)

	http.HandleFunc("/edit/", controller.TaskEdit)

	http.ListenAndServe(config.ListenAdd, nil)
}
