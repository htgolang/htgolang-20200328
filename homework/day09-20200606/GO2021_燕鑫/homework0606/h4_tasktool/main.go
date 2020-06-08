package main

import (
	. "homework0425/todolist"
	"net/http"
)

func main() {
	go func() {
		TaskMain("-1",nil)
	}()
	httpserver("127.0.0.1:3366")
}

func httpserver(addr string){
	http.HandleFunc("/select/", func(writer http.ResponseWriter, request *http.Request) {
		TaskMain("1",&writer)
	})
	http.ListenAndServe(addr,nil)
}
