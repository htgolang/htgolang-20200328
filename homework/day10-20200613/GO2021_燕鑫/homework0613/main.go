package main

import (
	"github.com/julienschmidt/httprouter"
	"homework0613/handler"
	"homework0613/tools"
	"net/http"
)

func main() {
	httpserver("127.0.0.1:3366")
}

func httpserver(addr string) {
	dburl := tools.DBFILE
	h:=handler.NewServiceHandler(dburl)
	mux := httprouter.New()
	mux.GET("/listtask/:user/:id/", h.ListTasksById)
	mux.GET("/listtask/:user/", h.ListTasksById)
	mux.GET("/updatetask/:user/:id/",h.UpdateTaskById)
	mux.POST("/updatetask/:user/:id/",h.UpdateTaskById)
	server := http.Server{Addr: addr, Handler: mux}
	_ = server.ListenAndServe()
}


