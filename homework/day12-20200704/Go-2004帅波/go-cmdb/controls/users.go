package controls
//ctxuser表示当前登陆用户
import (
	"github.com/astaxie/beego"
	"github.com/strive-after/go-cmdb/base/baseerr"
	"github.com/strive-after/go-cmdb/module"
	"strconv"
	"time"

)

type UserController struct {
	beego.Controller
}


//显示所有用户
func (usercon *UserController) Show() {
	var (
		//当前登陆用户
		useremail  string
		ctxuser module.User
		//所有用户列表
		users []module.User
		//当前页码
		ok bool
		role int
	)
	operation := module.NewOperation(&module.User{})
	errs := baseerr.New()
	//获取当前登录用户
	useremail,_ = usercon.Ctx.GetSecureCookie(Secret,"UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	allusers ,err :=  operation.GetAll(users)
	if err != nil {
		beego.Error("获取失败",err)
		errs.Add("Show","获取失败请联系管理员")
	}
	users ,ok = allusers.([]module.User)
	if !ok {
		beego.Error("转换失败",err)
		errs.Add("Show","获取失败请联系管理员")
	}
	ctxuser.Role = role
	if err = operation.GetId(&ctxuser);err != nil {
		beego.Error(err,"show 获取ctxuser失败")
	}
	if ctxuser.Role != role {
		usercon.SetSession(ctxuser.Email,ctxuser)
		role = ctxuser.Role
	}

	usercon.Data["errors"] = errs
	usercon.Data["conuserrole"] = role
	usercon.Data["UserName"] = ctxuser.Name
	usercon.Data["users"] = users
	usercon.Layout = `layout.html`
	usercon.TplName = `users/alluser.html`
}



//修改用户信息
func (usercon *UserController) ChangeUser() {
	var (
		user  module.User
		id int
		useremail string
		ctxuser module.User
		err error
	)
	operation := module.NewOperation(&module.User{})
	errs := baseerr.New()
	useremail ,_= usercon.Ctx.GetSecureCookie(Secret,"UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	id ,err =strconv.Atoi(usercon.GetString("id"))
	if err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"ChangeUserPost 获取id 失败")
		errs.Add("ChangeUser","修改失败")
	}
	user.ID = uint(id)
	err   = operation.GetId(&user)
	if err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"ChangeUserPost 获取用户 失败")
		errs.Add("ChangeUser","修改失败")
	}
	if usercon.Ctx.Input.IsPost() {
		err = usercon.ParseForm(&user)
		if err != nil {
			beego.Error(err, "当前登录用户", ctxuser.Name, "ChangeUserPost 获取前端传输用户 失败")
			errs.Add("ChangeUser","修改失败")
		}
		user.ID = uint(id)
		if err = operation.Update(&user); err != nil {
			beego.Error(err, "当前登录用户", ctxuser.Name, "ChangeUserPost 更新用户信息失败 失败")
			errs.Add("ChangeUser","修改失败")
		}
		err = operation.GetId(&user)
		if err != nil {
			beego.Error(err, "当前登录用户", ctxuser.Name, "ChangeUserPost 获取用户失败 失败")
			errs.Add("ChangeUser","修改失败")
		}
		//当用户修改的是自己的时候  把session中存放的user信息修改一下 保持最新
		if uint(id) == ctxuser.ID {
			usercon.Ctx.SetSecureCookie(Secret, "UserEmail", user.Email, time.Second*3600)
			usercon.SetSession(user.Email, user)
		}
		if !errs.HasErrors() {
			usercon.Redirect("/user/show", 302)
		}
	}
	usercon.Data["errors"] = errs
	usercon.Data["UserName"] = ctxuser.Name
	usercon.Data["user"] = user
	usercon.Layout = `layout.html`
	usercon.TplName = `users/ChangeUser.html`
}

//删除用户
func (usercon *UserController) Del() {
	var (
		id int
		useremail string
		ctxuser module.User
		err error
	)
	operation := module.NewOperation(&module.User{})
	errs := baseerr.New()
	useremail,_ = usercon.Ctx.GetSecureCookie(Secret,"UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	id ,err =strconv.Atoi(usercon.GetString("id"))
	if err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"DelUserGet 获取前id 失败")
		errs.Add("Del","删除失败")
	}
	//这里暂时不做权限判断  前端做了权限判断因为如果用户不是超管 那么不显示按钮
	//这里需要判断超级管理员不可以删除自己
	if uint(id) == ctxuser.ID {
		beego.Error("用户不可以删除自己")
		errs.Add("Del","删除失败")
		return
	}
	if err = operation.Del(uint(id));err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"DelUserGet 删除用户失败")
		errs.Add("Del","删除失败")
	}
	usercon.Redirect("/user/show",302)
}



func (usercon *UserController) Info() {
	var (
		id int
		useremail string
		user module.User
		ctxuser module.User
		err error
	)
	operation := module.NewOperation(&module.User{})
	errs := baseerr.New()
	useremail,_ = usercon.Ctx.GetSecureCookie(Secret,"UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	id ,err =strconv.Atoi(usercon.GetString("id"))
	if err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"UserInfo 获取id失败")
		errs.Add("Info","获取信息失败")
	}
	user.ID = uint(id)
	if err = operation.GetId(&user);err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"UserInfo 获取user失败")
		errs.Add("Info","获取信息失败")
	}
	usercon.Data["errors"] = errs
	usercon.TplName = "users/info.html"
	usercon.Data["user"] = user
	usercon.Data["UserName"] = ctxuser.Name
	usercon.Layout = `layout.html`
}

//查看当前登录用户以及修改当前登录用户信息
//修改个人资料
func (usercon *UserController)  MyInfo() {
	var (
		user module.User
		useremail string
		ctxuser module.User
		err error
	)
	operation := module.NewOperation(&module.User{})
	errs := baseerr.New()
	useremail ,_ = usercon.Ctx.GetSecureCookie(Secret,"UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	if usercon.Ctx.Input.IsPost() {
		user.ID = ctxuser.ID
		if err = usercon.ParseForm(&user); err != nil {
			beego.Error(err, "当前登录用户", ctxuser.Name, "MyInfoPost 获取前端传输用户信息失败")
			errs.Add("MyInfo","获取当前用户信息失败")
		}

		if err = operation.Update(&user); err != nil {
			beego.Error(err, "当前登录用户", ctxuser.Name, "MyInfoPost 用户信息更新失败")
			errs.Add("MyInfo","获取当前用户信息失败")
		}

		err = operation.GetId(&user)
		if err != nil {
			beego.Error(err, "当前登录用户", ctxuser.Name, "ChangeUserPost 获取用户失败 失败")
			errs.Add("MyInfo","获取当前用户信息失败")
		}
		if errs.HasErrors() {
			usercon.Ctx.SetSecureCookie(Secret, "UserEmail", user.Email, time.Second*3600)
			usercon.SetSession(user.Email, user)
			usercon.Redirect("/user/show", 302)
		}
	}
	usercon.Data["errors"] = errs
	usercon.TplName = "users/MyInfo.html"
	usercon.Layout = "layout.html"
	usercon.Data["user"] = ctxuser
	usercon.Data["UserName"] = ctxuser.Name
}


//修改当前登陆用户密码
func (usercon *UserController) MyPass() {
	var (
		ctxuser module.User
		useremail  string
		err error
	)
	//operation := module.NewOperation(&module.User{})
	errs := baseerr.New()
	useremail, _  = usercon.Ctx.GetSecureCookie(Secret,"UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	if usercon.Ctx.Input.IsPost() {
		oldpass := usercon.GetString("oldpass")
		newpass := usercon.GetString("newpass")
		err = module.ChangePasswordHash(oldpass,newpass,&ctxuser)
		if err != nil {
			beego.Error(err,"当前登陆用户",ctxuser.Name,"MyPassPost  密码更新失败")
			errs.Add("MyPass","密码修改失败")
		}
		//if err = operation.ChangePass(ctxuser.ID,oldpass,newpass);err  !=nil {
		//	beego.Error(err,"当前登陆用户",ctxuser.Name,"MyPassPost  密码更新失败")
		//	errs.Add("MyPass","密码修改失败")
		//}
		//如果错误切片为0  errs.HasErrors()返回false
		if !errs.HasErrors() {
			//修改密码完毕后让用户重新登陆
			usercon.DelSession(useremail)
			usercon.Ctx.SetSecureCookie(Secret, "UserEmail", useremail, -1)
			usercon.Redirect("/auth/login?email="+ctxuser.Email, 302)
		}

	}
	usercon.Data["errors"] = errs
	usercon.Data["Email"] = ctxuser.Email
	usercon.TplName= "users/MyPass.html"
}



func (usercon *UserController) UserPass() {
	var (
		id int
		useremail string
		user module.User
		ctxuser module.User
		err error
	)
	operation := module.NewOperation(&module.User{})
	errs := baseerr.New()
	useremail, _  = usercon.Ctx.GetSecureCookie(Secret,"UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	id ,err =strconv.Atoi(usercon.GetString("id"))
	if err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"UserPassPost 获取id失败")
		errs.Add("UserPass","密码修改失败")
	}
	user.ID = uint(id)
	if err = operation.GetId(&user);err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"UserPassPost 获取user失败")
		errs.Add("UserPass","密码修改失败")
	}

	if usercon.Ctx.Input.IsPost() {
		password := usercon.GetString("PassWord")
		err = module.ChangePasswordHash("",password,&user)
		if err != nil {
			beego.Error("修改失败",err)
			errs.Add("UserPass","密码修改失败")
		}

		if user.ID == ctxuser.ID {
			usercon.DelSession(useremail)
			usercon.Ctx.SetSecureCookie(Secret,"UserEmail",useremail,-1)
			usercon.Redirect("/auth/login?email="+ctxuser.Email,302)
		}
		if !errs.HasErrors() {
			usercon.Redirect("/user/show", 302)
		}

	}
	usercon.Data["errors"] = errs
	usercon.Data["user"] =  user
	usercon.Data["UserName"] = ctxuser.Name
	usercon.TplName = "users/UserPass.html"
	usercon.Layout = "layout.html"
}