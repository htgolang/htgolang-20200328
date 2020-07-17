package controls

import (
	"github.com/astaxie/beego"
	"github.com/strive-after/go-cmdb/base/baseerr"
	"github.com/strive-after/go-cmdb/module"
)

type ManagerController struct {
	beego.Controller
}

func (m *ManagerController) All() {
	errs := baseerr.New()
	managers := []module.Manager{}
	useremail,_ := m.Ctx.GetSecureCookie(Secret,"UserEmail")
	ctxuser := m.GetSession(useremail).(module.User)
	operation := module.NewOperation(&module.Manager{})
	err := operation.GetAll(&managers)
	if err != nil {
		beego.Error("获取失败",err)
		errs.Add("ManagerAll","获取失败请联系管理员")
	}
	m.TplName = "manager/alltask.html"
	m.Layout = `layout.html`
	m.Data["UserName"] = ctxuser.Name
	m.Data["Managers"] = managers
}

//func (m *Manager) Job() {
//
//}