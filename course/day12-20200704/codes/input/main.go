package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

const (
	cookieKey = "abc123"
)

type InputForm struct {
	Name     string
	Password string `form:"password"`
}

type InputController struct {
	beego.Controller
}

func (c *InputController) QueryParams() {
	// 方式1
	// c.Ctx.Request.ParseForm()
	// fmt.Println(c.Ctx.Request.Form)

	// 方式2
	// fmt.Println(c.Ctx.Input.Query("name"))

	// 方式3
	// var name string
	// c.Ctx.Input.Bind(&name, "name")
	// fmt.Println(name)

	// 方式4
	// fmt.Println(c.GetString("name"))

	// 方式5
	// var form InputForm
	// fmt.Println(c.ParseForm(&form))
	// fmt.Printf("%#v\n", form)

	// 方式6
	fmt.Println(c.Input())
	c.Ctx.WriteString("")
}

func (c *InputController) Form() {
	// // 方式1
	// c.Ctx.Request.ParseForm()
	// fmt.Println(c.Ctx.Request.Form)

	// 方式2(只能从POst body中获取数据必须用Request.PostForm)
	// c.Ctx.Request.ParseForm()
	// fmt.Println(c.Ctx.Request.PostForm)

	// 方式3
	// fmt.Println(c.GetString("name"))
	var form InputForm
	c.ParseForm(&form)
	fmt.Printf("%#v\n", form)

	// 方式4 c.Ctx.Input.Bind
	// 方式5 c.Input
	// 方式6 c.Ctx.Input.Query

	c.Ctx.WriteString("")
}

func (c *InputController) File() {
	// 1. Request
	// 2. GetFile
	// c.GetFile("name")
	// 3. SaveTo
	c.SaveToFile("img", "./upload/a.jpg")
	c.Ctx.WriteString("")
}

func (c *InputController) Json() {
	var m map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	fmt.Printf("%#v\n", m)
	c.Ctx.WriteString("")
}

func (c *InputController) Cookie() {
	// 获取方式
	// 方式1
	// cookie, err := c.Ctx.Request.Cookie("name")
	// fmt.Println(cookie, err)
	// 方式2
	// fmt.Println(c.Ctx.Input.Cookie("name"))
	// 方式3
	fmt.Println(c.Ctx.GetCookie("name"))
	// 设置
	// 方式1
	// http.SetCookie()
	// 方法2
	c.Ctx.SetCookie("name", "vavvvv")
	c.Ctx.WriteString("")
}

func (c *InputController) SecureCookie() {
	fmt.Println(c.Ctx.GetSecureCookie(cookieKey, "test"))
	c.Ctx.SetSecureCookie(cookieKey, "test", "vvvv")
	// c.GetSecureCookie()
	// c.SetSecureCookie()
}

func (c *InputController) Header() {
	fmt.Println(c.Ctx.Request.Method)
	fmt.Println(c.Ctx.Request.URL)
	fmt.Println(c.Ctx.Request.Header)
	fmt.Println(c.Ctx.Input.URI())
	fmt.Println(c.Ctx.Input.Method())
	fmt.Println(c.Ctx.Input.IP())
	c.Ctx.WriteString("")
}

func main() {
	// 提交数据的方式
	// GET ?queryparam
	// Post ?queryparams
	//		request body
	// 		content-type: application/x-www-form-urlencoded
	//                    application/json js 客户端
	//                    multipart/form-data
	//
	beego.AutoRouter(&InputController{})
	beego.Run()
}
