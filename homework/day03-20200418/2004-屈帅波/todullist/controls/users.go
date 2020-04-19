package controls

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"todulist/module"
)

type UserControoler struct {
	beego.Controller
}
//增加用户
func (usercon *UserControoler) AddUserGet() {
	usercon.TplName = `adduser.html`
}

func (usercon *UserControoler) AddUserPost() {
	user := module.User{}
	err := usercon.ParseForm(&user)
	if err != nil {
		beego.Error("输入有误")
		return
	}
	o := orm.NewOrm()
	if err = o.Read(&user,"Name");err == nil{
		beego.Error("用户存在")
		return
	}
	user.Password = fmt.Sprintf("%x",md5.Sum([]byte(user.Password)))
	if _,err = o.Insert(&user);err != nil {
		beego.Error("插入失败")
		return
	}
	usercon.Redirect("/user",302)

}
//显示所有用户
func (usercon *UserControoler) ShowUserGet() {
	o := orm.NewOrm()
	users := []module.User{}
	o.QueryTable("user").All(&users)
	usercon.TplName = `alluser.html`
	usercon.Data["users"] = users
}
//查询条件
func (usercon *UserControoler) SelUserGet() {
	usercon.TplName = `seluser.html`
}
func (usercon *UserControoler) SelUserPost() {
	user := module.User{}
	user.Name = usercon.GetString("Name")
	o := orm.NewOrm()
	err := o.Read(&user,"Name")
	if err != nil {
		beego.Error("用户不存在",err)
		return
	}
	usercon.Data["user"] = user
	usercon.TplName = `seluser.html`
}

//修改用户信息
func (usercon *UserControoler) ChangeUserGet() {
	id ,_:=strconv.Atoi(usercon.GetString("id"))
	user := module.User{Id:id}
	o := orm.NewOrm()
	if err := o.Read(&user);err != nil {
		beego.Error(err)
		return
	}
	usercon.TplName = `changeuser.html`
	usercon.Data["user"] = user
}
func (usercon *UserControoler) ChangeUserPost() {
	id, _ := strconv.Atoi(usercon.GetString("id"))
	user := module.User{Id:id}
	o := orm.NewOrm()
	err := o.Read(&user)
	if err != nil {
		beego.Error(err)
		return
	}
	user.Password = fmt.Sprintf("%x",md5.Sum([]byte(usercon.GetString("Password"))))
	user.Name = usercon.GetString("Name")
	if _,err = o.Update(&user);err != nil{
		beego.Error(err)
		return
	}
	usercon.Redirect("/user/show",302)
}


func (usercon *UserControoler) DelUserGet() {
	id, _ := strconv.Atoi(usercon.GetString("id"))
	user := module.User{Id:id}
	o := orm.NewOrm()
	if _, err := o.Delete(&user);err != nil {
		beego.Error(err)
		return
	}
	usercon.Redirect("/user/show",302)
}

