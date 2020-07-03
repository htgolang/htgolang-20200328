package controls

import "github.com/astaxie/beego"

const (
	UserErr  = "/user/err"
	RegErr = "/RegErr"
)

var (
	err error
)


type ErrControls struct {
	beego.Controller
}

func (err *ErrControls)  LoginErr() {
	err.TplName = "error/login404.html"
}

func (err *ErrControls)  RegErr() {
	err.TplName = "error/reg404.html"
}

func (err *ErrControls)  UserErr() {
	err.TplName = "error/user404.html"
}

func (err *ErrControls)  TaskErr() {
	err.TplName = "error/task404.html"
}

func (err *ErrControls) Error404() {
	err.TplName = "error/404.html"
}

func (err *ErrControls) RoleErr() {
	err.TplName = "error/role404.html"
}
