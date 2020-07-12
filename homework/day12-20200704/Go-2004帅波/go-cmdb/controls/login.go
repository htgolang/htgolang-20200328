package controls

import (
	"github.com/astaxie/beego"
	"github.com/strive-after/go-cmdb/base/baseerr"
	"github.com/strive-after/go-cmdb/module"
	"net/http"
	"time"
)
var (
	Secret string = "CMDB"
)
type AuthController struct {
	beego.Controller
}


type Operation struct {
	beego.Controller
}



//登陆数据处理
func (login *AuthController) Login()  {
	var (
		user module.User
		err error
		email string
	)
	errs := baseerr.New()
	//如果是get直接加载页面
	//如果是post 做数据 处理
	email = login.GetString("email")
	if login.Ctx.Input.IsPost() {
		if err := login.ParseForm(&user);err != nil {
			errs.Add("Login","登陆失败")
			beego.Error("登陆失败")
			email = user.Email
		}
		err = module.CheckPassword(user.Password,&user)
		if err != nil {
			errs.Add("Login","登陆失败")
			beego.Error("登陆失败，2",errs)
			email = user.Email
		}

		//如果记住用户名那么cookie保存时间为3600s
		err = user.Get("email",&user)
		if err != nil {
			errs.Add("Login","登陆失败,查询错误")
			beego.Error(err)
			email = user.Email
		}
		if !errs.HasErrors() {
			login.Ctx.SetSecureCookie(Secret,"UserEmail",user.Email,time.Second*3600)
			login.SetSession(user.Email,user)
			login.Redirect("/",http.StatusFound)
			return
		}
	}
	login.Data["Email"] = email
	login.Data["errors"] = errs
	login.TplName = `login.html`
}


//删除session

func (login *AuthController) Out() {
	email ,_:= login.Ctx.GetSecureCookie(Secret,"UserEmail")
	login.DelSession(email)
	login.Ctx.SetSecureCookie(Secret,"UserEmail",email,-1)
	login.Redirect("/auth/login?email="+email,302)
}


//用户注册数据处理
func (reg *AuthController) Reg() {
	var (
		inputuser module.User
		user  module.Operation   = new(module.User)
	)
	errs :=  baseerr.New()
	if reg.Ctx.Input.IsPost() {
		//将前端获取的数据直接赋值给user
		err := reg.ParseForm(&inputuser)
		if err != nil {
			beego.Error(err)
			errs.Add("Reg","注册失败")
		}
		err  = user.Add(&inputuser)
		if err != nil {
			beego.Error(err)
			errs.Add("Reg","注册失败")
		}
		if !errs.HasErrors() {
			reg.Redirect("/auth/login",302)
			return
		}
	}
	reg.TplName = `reg.html`
	reg.Data["errors"] = errs
}


func (login *Operation)  Get() {
	userEmail,ok  := login.Ctx.GetSecureCookie(Secret,"UserEmail")
	if !ok {
		beego.Error("获取cookie失败")
		return
	}
	user := login.GetSession(userEmail).(module.User)
	login.Data["UserName"] = user.Name
	login.TplName = `operation.html`
	login.Layout = `layout.html`
}


