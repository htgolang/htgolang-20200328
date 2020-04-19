package controls

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"todulist/module"
)

type Taskcontrollers struct {
	beego.Controller
}
//操作页面
func (taskcon *Taskcontrollers) OperTaskGet() {
	taskcon.TplName= `opertask.html`
}

//添加任务
func (taskcon *Taskcontrollers) AddTaskGet() {
	taskcon.TplName=`addtask.html`
}
func (taskcon *Taskcontrollers) AddTaskPost() {
	task := module.Manage{}
	err := taskcon.ParseForm(&task)
	if err != nil {
		beego.Error(err)
		return
	}
	o := orm.NewOrm()
	if _,err =o.Insert(&task);err != nil {
		beego.Error(err)
		return
	}
	taskcon.Redirect("/task",302)
}

//显示所有任务
func (taskcon *Taskcontrollers) ShowTaskGet() {
	tasks := []module.Manage{}
	if _, err := orm.NewOrm().QueryTable("manage").All(&tasks);err != nil{
		beego.Error(err)
		return
	}
	taskcon.TplName= `alltask.html`
	taskcon.Data["tasks"] = tasks
}

//查询任务
func (taskcon *Taskcontrollers) SelTaskGet() {
	taskcon.TplName=`seltaskname.html`
}
func (taskcon *Taskcontrollers) SelTaskPost() {
	taskname := taskcon.GetString("TaskName")
	task := module.Manage{TaskName:taskname}
	o := orm.NewOrm()
	if err := o.Read(&task,"task_name");err != nil {
		beego.Error("任务不存在",err)
		return
	}
	taskcon.Data["task"] = task
	taskcon.TplName= `seltask.html`
}

//删除任务
func (taskcon *Taskcontrollers) DelTaskGet() {
	id, _ := strconv.Atoi(taskcon.GetString("id"))
	task := module.Manage{Id:id}
	if err := orm.NewOrm().Read(&task);err != nil {
		beego.Error(err)
		return
	}
	orm.NewOrm().Delete(&task)
	taskcon.Redirect("/task/show",302)
}



