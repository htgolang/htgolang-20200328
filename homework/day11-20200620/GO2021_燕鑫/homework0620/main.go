package main

import (
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"homework0620/handler"
	"homework0620/tools"
	"net/http"
	"os"
)

func main() {
	httpserver("127.0.0.1:3366")
}

func httpserver(addr string) {
	db, err := gorm.Open("mysql", tools.DBURL)
	if err != nil {
		os.Exit(36)
	}
	//db.LogMode(true)
	defer db.Close()
	h := handler.NewServiceHandler(db)
	mux := httprouter.New()
	mux.GET("/login/", h.Login)
	mux.POST("/login/", h.Login)
	mux.POST("/signup/",h.SignUp)
	mux.GET("/listtask/:user/:id/", h.ListTasksById)
	mux.GET("/listtask/:user/", h.ListTasksById)
	mux.GET("/updatetask/:user/:id/", h.UpdateTaskById)
	mux.POST("/updatetask/:user/:id/", h.UpdateTaskById)
	mux.GET("/deletetask/:user/:id/", h.DeleteTaskById)
	mux.GET("/createtask/:user/", h.CreateTask)
	mux.POST("/createtask/:user/", h.CreateTask)
	mux.GET("/logout/", h.Logout)
	server := http.Server{Addr: addr, Handler: mux}
	_ = server.ListenAndServe()
}
