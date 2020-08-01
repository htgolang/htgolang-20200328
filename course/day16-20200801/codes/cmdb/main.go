package main

import (
	_ "github.com/go-sql-driver/mysql"

	"cmdb/cmds"
	_ "cmdb/routers"
)

func main() {
	cmds.Execute()
}
