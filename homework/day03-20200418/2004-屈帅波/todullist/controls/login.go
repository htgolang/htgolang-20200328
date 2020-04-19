package controls

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"todulist/module"
)

type LoginControllers struct {
	beego.Controller
}

type RegisterControllers struct {
	beego.Controller
}
func (login *LoginControllers) Get(){
	login.Redirect("/login",302)
}

//登陆显示页面
func (login *LoginControllers) LoginGet() {
	login.TplName = `login.html`
}
//登陆数据处理
func (login *LoginControllers) LoginPost()  {
	user := module.User{}
	//从前端获取用户输入信息
	if err := login.ParseForm(&user);err != nil {
		beego.Error("获取失败")
		return
	}
	//把信息拿出来
	passwd := fmt.Sprintf("%x",md5.Sum([]byte(user.Password)))
	//根据用户id从数据库查询对应的用户信息作比较
	o := orm.NewOrm()
	if err := o.Read(&user,"Name");err != nil {
		beego.Error("用户不存在")
		return
	}
	if  passwd == user.Password {
		login.Redirect("/operation" ,302)
	}else {
		beego.Error("密码输入错误")
		return
	}

}

//注册用户显示页面
func (reg *RegisterControllers) RegGet() {
	reg.TplName = `reg.html`
}
//用户注册数据处理
func (reg *RegisterControllers) RegPost() {
	user := module.User{}
	//将前端获取的数据直接赋值给user
	err := reg.ParseForm(&user)
	if err != nil {
		beego.Error(err)
		return
	}
	//这里暂时先判断用户不能为空  密码的判断随后再加
	if  user.Name == "" {
		beego.Error("输入不能为空")
		return
	}
	//密码md5加盐
	user.Password = fmt.Sprintf("%x",md5.Sum([]byte(user.Password)))
	o := orm.NewOrm()
	//查询用户是否存在 当查询没问题的时候用户是存在的
	if err = o.Read(&user,"Name");err == nil{
		beego.Error("用户存在")
		return
	}
	if _,err = o.Insert(&user);err != nil {
		beego.Error("存储失败")
		return
	}
	reg.Redirect("/login",302)
}