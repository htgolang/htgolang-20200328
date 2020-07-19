package main

import (
	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"

	_ "cmdb/routers"
)

func main() {
	beego.Run()
}
