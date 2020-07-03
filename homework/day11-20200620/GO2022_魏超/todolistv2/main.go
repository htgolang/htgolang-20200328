package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"todolist/models"
	"todolist/utils"
)

var Config utils.TomlConfig

func main() {
	var (
		err error
		dsn string
	)

	Config, err = utils.ParseTomlConfig()
	if err != nil {
		log.Fatal(err)
	}

	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		Config.Mysql.UserName,
		Config.Mysql.Password,
		Config.Mysql.Host,
		Config.Mysql.Port,
		Config.Mysql.DBName,
	)
	err = models.MySQLDB("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// 添加js和css
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("views/css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("views/js/"))))

	err = http.ListenAndServe(net.JoinHostPort(Config.HttpServer.Host, strconv.Itoa(Config.HttpServer.Port)), nil)
	if err != nil {
		log.Fatal(err)
	}
}
