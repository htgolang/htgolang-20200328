package controls
//ctxuser表示当前登陆用户
import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/strive-after/go-kubernetes/module"
	"strconv"
	"time"
)

type UserControoler struct {
	beego.Controller
}


//显示所有用户
func (usercon *UserControoler) ShowUserGet() {
	var (
		//当前登陆用户
		useremail  string
		ctxuser module.User
		//所有用户列表
		users []module.User
		//当前页码
	)

	//获取当前登录用户
	useremail = usercon.Ctx.GetCookie("UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	err = module.Db.Model(module.User{}).Find(&users).Error
	if err != nil {
		beego.Error("用户列表获取失败")
		usercon.Redirect(UserErr,302)
		return
	}
	usercon.Data["conuserrole"] = ctxuser.Role
	usercon.Data["UserName"] = ctxuser.Name
	usercon.Data["users"] = users
	usercon.Layout = `layout.html`
	usercon.TplName = `users/alluser.html`
}




//修改用户信息
func (usercon *UserControoler) ChangeUserGet() {
	var (
		//当前登录用户
		//被修改用户
		operation   module.Operation = new(module.User)
		id int
		useremail string
		user module.User
		ctxuser module.User
	)
	//获取修改用户id

	id ,err =strconv.Atoi(usercon.GetString("id"))
	if err != nil {
		beego.Error(err,"当前登录用户",useremail,"ChangeUserGet 获取id 失败")
		usercon.Redirect(UserErr,302)
		return
	}
	user.ID = uint(id)

	err = operation.GetId(&user)
	if err != nil {
		beego.Error(err,"当前登录用户",useremail,"ChangeUserGet 获取user失败")
		usercon.Redirect(UserErr,302)
		return
	}
	useremail = usercon.Ctx.GetCookie("UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	usercon.Data["UserName"] = ctxuser.Name
	usercon.Data["user"] = user
	usercon.Layout = `layout.html`
	usercon.TplName = `users/ChangeUser.html`
}

//修改用户信息
func (usercon *UserControoler) ChangeUserPost() {
	var (
		operation   module.Operation = new(module.User)
		user  module.User
		id int
		useremail string
		ctxuser module.User
	)
	useremail = usercon.Ctx.GetCookie("UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	id ,err =strconv.Atoi(usercon.GetString("id"))
	if err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"ChangeUserPost 获取id 失败")
		usercon.Redirect(UserErr,302)
		return
	}
	err = usercon.ParseForm(&user)
	if err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"ChangeUserPost 获取前端传输用户 失败")
		usercon.Redirect(UserErr,302)
		return
	}
	user.ID = uint(id)
	if err = operation.Update(&user) ;err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"ChangeUserPost 更新用户信息失败 失败")
		usercon.Redirect(UserErr,302)
		return
	}
	err  = operation.GetId(&user)
	if err != nil {
		beego.Error(err, "当前登录用户", ctxuser.Name, "ChangeUserPost 获取用户失败 失败")
		usercon.Redirect(UserErr, 302)
		return
	}
	//当用户修改的是自己的时候  把session中存放的user信息修改一下 保持最新
	if uint(id) == ctxuser.ID {
		usercon.Ctx.SetCookie("UserEmail",user.Email,time.Second*3600)
		usercon.SetSession(user.Email,user)
	}


	usercon.Redirect("/user/show?UserIndex=1",302)
}

//删除用户
func (usercon *UserControoler) Del() {
	var (
		id int
		operation  module.Operation= new(module.User)
		useremail string
		ctxuser module.User
	)
	fmt.Println("delete")
	useremail = usercon.Ctx.GetCookie("UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	id ,err =strconv.Atoi(usercon.GetString("id"))
	if err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"DelUserGet 获取前id 失败")
		usercon.Redirect(UserErr,302)
		return
	}
	//这里暂时不做权限判断  前端做了权限判断因为如果用户不是超管 那么不显示按钮
	//这里需要判断超级管理员不可以删除自己
	if uint(id) == ctxuser.ID {
		beego.Error("用户不可以删除自己")
		usercon.Redirect(UserErr,302)
		return
	}



	if err = operation.Del(uint(id));err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"DelUserGet 删除用户失败")
		usercon.Redirect(UserErr,302)
		return
	}
	usercon.Redirect("/user/show",302)
}



func (usercon *UserControoler) UserInfo() {
	var (
		operation  module.Operation = new(module.User)
		id int
		useremail string
		user module.User
		ctxuser module.User
	)
	useremail = usercon.Ctx.GetCookie("UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	id ,err =strconv.Atoi(usercon.GetString("id"))
	if err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"UserInfo 获取id失败")
		usercon.Redirect(UserErr,302)
		return
	}
	user.ID = uint(id)
	if err = operation.GetId(&user);err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"UserInfo 获取user失败")
		usercon.Redirect(UserErr,302)
		return
	}
	usercon.TplName = "users/lookuser.html"
	usercon.Data["user"] = user
	usercon.Data["UserName"] = ctxuser.Name
	usercon.Layout = `layout.html`
}

//查看当前登录用户以及修改当前登录用户信息
func (usercon *UserControoler)  MyInfoGet() {
	var (
		useremail string
		ctxuser module.User
	)
	useremail = usercon.Ctx.GetCookie("UserEmail")
	//因为我们有过滤器函数 如果session为空那么会直接返回登陆页面 所以可以直接用类型转换
	if usercon.GetSession(useremail) == nil {
		beego.Error("MyInfoGet  Session 不存在")
		usercon.Redirect("/login",302)
		return
	}
	ctxuser = usercon.GetSession(useremail).(module.User)
	usercon.TplName = "users/MyInfo.html"
	usercon.Layout = "layout.html"
	usercon.Data["user"] = ctxuser
}
//修改个人资料
func (usercon *UserControoler)  MyInfoPost() {
	var (
		operation  module.Operation = new(module.User)
		user module.User
		useremail string
		ctxuser module.User
	)
	useremail = usercon.Ctx.GetCookie("UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	user.ID = ctxuser.ID
	if err = usercon.ParseForm(&user);err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"MyInfoPost 获取前端传输用户信息失败")
		usercon.Redirect(UserErr,302)
		return
	}

	if err = operation.Update(&user);err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"MyInfoPost 用户信息更新失败")
		usercon.Redirect(UserErr,302)
		return
	}

	err  = operation.GetId(&user)
	if err != nil {
		beego.Error(err, "当前登录用户", ctxuser.Name, "ChangeUserPost 获取用户失败 失败")
		usercon.Redirect(UserErr, 302)
		return
	}

	usercon.Ctx.SetCookie("UserEmail",user.Email,time.Second*3600)
	usercon.SetSession(user.Email,user)
	usercon.Redirect("/user/show?UserIndex=1",302)
}


//修改当前用户密码
func (usercon *UserControoler) MyPassGet() {
	var  (
		useremail string
		ctxuser module.User
	)
	useremail = usercon.Ctx.GetCookie("UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	usercon.Data["Email"] = ctxuser.Email
	usercon.TplName= "users/MyPass.html"
}

func (usercon *UserControoler) MyPassPost() {
	var (
		operation  module.Operation = new(module.User)
		ctxuser module.User
		useremail  string
		err error
	)
	useremail = usercon.Ctx.GetCookie("UserEmail")
	if usercon.GetSession(useremail) == nil {
		beego.Error("MyPassPost  Session 不存在")
		usercon.Redirect("/login",302)
		return
	}
	ctxuser = usercon.GetSession(useremail).(module.User)
	oldpass := usercon.GetString("oldpass")
	newpass := usercon.GetString("newpass")
	if err = operation.ChangePass(ctxuser.ID,oldpass,newpass);err  !=nil {
		beego.Error(err,"当前登陆用户",ctxuser.Name,"MyPassPost  密码更新失败")
		usercon.Redirect("/user/err",302)
	}
	usercon.DelSession(useremail)
	usercon.Ctx.SetCookie("UserEmail",useremail,-1)
	usercon.Redirect("/login",302)
	//修改密码完毕后让用户重新登陆
}

//修改用户密码
func (usercon *UserControoler) UserPassGet() {
	var (
		operation  module.Operation = new(module.User)
		id int
		useremail string
		user  module.User
		ctxuser module.User
	)
	useremail = usercon.Ctx.GetCookie("UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	id ,err =strconv.Atoi(usercon.GetString("id"))
	if err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"UserPassGet 获取id失败")
		usercon.Redirect(UserErr,302)
		return
	}

	user.ID = uint(id)
	if err = operation.GetId(&user);err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"UserPassGet 获取user失败")
		usercon.Redirect(UserErr,302)
		return
	}
	usercon.Data["user"] =  user
	usercon.Data["Username"] = ctxuser.Name
	usercon.TplName = "users/UserPass.html"
	usercon.Layout = "layout.html"
}

func (usercon *UserControoler) UserPassPost() {
	var (
		operation  module.Operation = new(module.User)
		id int
		useremail string
		user module.User
		ctxuser module.User
	)
	useremail = usercon.Ctx.GetCookie("UserEmail")
	ctxuser = usercon.GetSession(useremail).(module.User)
	id ,err =strconv.Atoi(usercon.GetString("id"))
	if err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"UserPassPost 获取id失败")
		usercon.Redirect(UserErr,302)
		return
	}
	user.ID = uint(id)
	if err = operation.GetId(&user);err != nil {
		beego.Error(err,"当前登录用户",ctxuser.Name,"UserPassPost 获取user失败")
		usercon.Redirect(UserErr,302)
		return
	}
	password := usercon.GetString("PassWord")
	user.ChangePass(uint(id)," ",password)
	usercon.Redirect("/user/show?UserIndex=1",302)
}