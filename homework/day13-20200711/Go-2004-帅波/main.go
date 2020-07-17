package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"

	_ "github.com/strive-after/go-cmdb/module"
	_ "github.com/strive-after/go-cmdb/route"
)

func main() {
	beego.Run()
}
