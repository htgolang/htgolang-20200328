package main

import (
	"encoding/xml"

	"github.com/astaxie/beego"
)

type OutputController struct {
	beego.Controller
}

func (c *OutputController) CString() {
	c.Ctx.WriteString("Context WriteString")
}

func (c *OutputController) OutputBody() {
	c.Ctx.Output.Body([]byte("Output Body"))
}

func (c *OutputController) Tpl() {
	c.TplName = "output.html"
}

func (c *OutputController) Json() {
	c.Data["json"] = map[string]string{"a": "xxx", "b": "yyyy"}
	c.ServeJSON()
}

type User struct {
	Name string
	Addr string
}

func (c *OutputController) Xml() {
	c.Data["xml"] = struct {
		XMLName xml.Name `xml:"root"`
		User    User     `xml:"user`
	}{User: User{Name: "kk", Addr: "127.0.0.1"}}
	c.ServeXML()
}

func (c *OutputController) Yaml() {
	c.Data["yaml"] = map[string]string{"a": "xxx", "b": "yyyy"}
	c.ServeYAML()
}

func (c *OutputController) Redir() {
	c.Redirect("http://www.baidu.com", 302)
}

func (c *OutputController) StopRun() {
	c.StopRun()
}

func (c *OutputController) Ab() {
	c.Abort("404")
}

func main() {
	beego.AutoRouter(&OutputController{})
	beego.Run()
}
