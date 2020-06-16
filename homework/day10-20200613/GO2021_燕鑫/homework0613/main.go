package main

import (
	"github.com/julienschmidt/httprouter"
	"homework0613/handler"
	"net/http"
)

func main() {
	httpserver("127.0.0.1:3366")
}

func httpserver(addr string) {
	mux := httprouter.New()
	mux.GET("/listtask/:user/:id/", handler.ListTasksById)
	mux.GET("/listtask/:user/", handler.ListTasksById)
	server := http.Server{Addr: addr, Handler: mux}
	_ = server.ListenAndServe()
}


