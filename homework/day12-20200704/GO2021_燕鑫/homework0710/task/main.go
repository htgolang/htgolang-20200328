package main

import (
	"github.com/astaxie/beego"
	_ "task/router"
)

func main() {
	beego.Run()
}

//func httpserver(addr string) {
//	db, err := gorm.Open("mysql", tools.DBURL)
//	if err != nil {
//		os.Exit(36)
//	}
//	//db.LogMode(true)
//	defer db.Close()
//	h := handler.NewServiceHandler(db)
//	mux := httprouter.New()
//	mux.GET("/login/", h.Login)
//	mux.POST("/login/", h.Login)
//	mux.POST("/signup/",h.SignUp)
//	mux.GET("/listtask/:user/:id/", h.ListTasksById)
//	mux.GET("/listtask/:user/", h.ListTasksById)
//	mux.GET("/updatetask/:user/:id/", h.UpdateTaskById)
//	mux.POST("/updatetask/:user/:id/", h.UpdateTaskById)
//	mux.GET("/deletetask/:user/:id/", h.DeleteTaskById)
//	mux.GET("/createtask/:user/", h.CreateTask)
//	mux.POST("/createtask/:user/", h.CreateTask)
//	mux.GET("/logout/", h.Logout)
//	server := http.Server{Addr: addr, Handler: mux}
//	_ = server.ListenAndServe()
//}
