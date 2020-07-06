package main

import (
	"log"
	"net"
	"net/http"
	"strconv"

	_ "todolist/controller"
	"todolist/global"
	_ "todolist/models"
)

func main() {
	var (
		err error
	)

	log.Printf("start listen %s:%d\n", global.Config.HttpServer.Host, global.Config.HttpServer.Port)
	err = http.ListenAndServe(net.JoinHostPort(global.Config.HttpServer.Host, strconv.Itoa(global.Config.HttpServer.Port)), nil)
	if err != nil {
		log.Fatal(err)
	}
}
