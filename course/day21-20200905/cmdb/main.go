package main

import (
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"

	"cmdb/cmds"
	_ "cmdb/routers"
)

func main() {
	cmds.Execute()
}
