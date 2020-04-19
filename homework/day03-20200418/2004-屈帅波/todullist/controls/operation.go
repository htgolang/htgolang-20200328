package controls

import "github.com/astaxie/beego"

type OperationCotrollers struct {
	beego.Controller
}
//总页面
func (oper *OperationCotrollers) AllOperGet() {
	oper.TplName = `alloperation.html`
}
//用户管理列表
func (oper *OperationCotrollers) UserOperGet() {
	oper.TplName = `useroperation.html`
}
//任务操作列表
func (oper *OperationCotrollers) TaskOperGet() {
	oper.TplName = `opertask.html`
}
