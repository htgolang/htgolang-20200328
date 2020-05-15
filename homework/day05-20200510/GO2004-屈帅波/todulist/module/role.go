package module

import (
	"github.com/astaxie/beego/orm"
)

//权限判断
func AuthRole(username string,role int) bool {
	o := orm.NewOrm()
	user := User{Name:username}
	o.Read(&user,"Name")
	//当用户的权限小于 做此操作的权限那么返回false
	if user.Role < role {
		return false
	}
	//用户权限大于做此操作权限那么返回true
	return  true
}
