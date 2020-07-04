package main

import (
	"fmt"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Index() {

}

func main() {
	fmt.Println(beego.AppConfig.String("mysql::MYSQL_HOST"))
	fmt.Println(beego.AppConfig.Int("mysql::MYSQL_PORT"))
	fmt.Println(beego.AppConfig.DefaultBool("mysql::MYSQL_PARSETIME", true))

	beego.AutoRouter(&HomeController{})
	beego.Run()
}
