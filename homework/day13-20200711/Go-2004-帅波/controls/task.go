package controls
//
//import (
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/orm"
//	"strconv"
//	"todulist/module"
//)
//
//type Taskcontrollers struct {
//	beego.Controller
//}
//
//
////添加任务
//func (taskcon *Taskcontrollers) AddTaskGet() {
//	username := taskcon.Ctx.GetCookie("UserName")
//	user := taskcon.GetSession(username)
//	beego.Info(user)
//	taskcon.Data["UserName"] = username
//	taskcon.Layout = "layout.html"
//	taskcon.TplName=`manager/addtask.html`
//}
//
//func (taskcon *Taskcontrollers) AddTaskPost() {
//	task := module.Manage{}
//	user := module.User{Name:taskcon.GetSession("UserName").(string)}
//	err := taskcon.ParseForm(&task)
//	if err != nil {
//		beego.Error(err)
//		taskcon.Redirect("/task/err",302)
//	}
//	o := orm.NewOrm()
//	o.Read(&user,"Name")
//	task.User = &user
//	if _,err =o.Insert(&task);err != nil {
//		beego.Error(err)
//		taskcon.Redirect("/task/err",302)
//	}
//	taskcon.Redirect("/task",302)
//}
//
////显示所有任务
//func (taskcon *Taskcontrollers) ShowTaskGet() {
//	username := taskcon.GetSession("UserName")
//	taskcon.Data["UserName"] = username
//	taskcon.Layout = "layout.html"
//	tasks := []module.Manage{}
//
//	Index,err := strconv.Atoi(taskcon.GetString("taskIndex"))
//	module.ErrInfo("获取当前页码失败",err)
//
//	//一页2行数据
//	pagesize := 1
//	//获取页数 向上取整
//	//起始量  因为从0开始所以-1  数据量起始点是0
//	//总共的页数= 总数据数量/ 单页数据量
//	count,start,pagenum,pagesize := module.Page("manage",Index,pagesize)
//	if _, err := orm.NewOrm().QueryTable("manage").RelatedSel("User").Limit(pagesize,start).All(&tasks);err != nil{
//		beego.Error(err)
//		taskcon.Redirect("/task/err",302)
//	}
//	taskcon.TplName= `manager/alltask.html`
//	taskcon.Data["tasks"] = tasks
//	taskcon.Data["taskIndex"] = Index
//	taskcon.Data["pagenum"] = pagenum
//	taskcon.Data["count"] = count
//
//}
//
////查询任务
//func (taskcon *Taskcontrollers) SelTaskGet() {
//	username := taskcon.GetSession("UserName")
//	taskcon.Data["UserName"] = username
//	taskcon.Layout = "layout.html"
//	taskcon.TplName=`manager/seltaskname.html`
//}
//
//func (taskcon *Taskcontrollers) SelTaskPost() {
//	taskname := taskcon.GetString("TaskName")
//	task := module.Manage{TaskName:taskname}
//	o := orm.NewOrm()
//	if err := o.Read(&task,"task_name");err != nil {
//		beego.Error("任务不存在",err)
//		taskcon.Redirect("/task/err",302)
//	}
//	taskcon.Data["task"] = task
//	taskcon.TplName= `manager/seltask.html`
//}
//
////暂停任务
//func (taskcon *Taskcontrollers) PauseTaskGet() {
//	StatusTask(taskcon,2)
//}
//
////取消任务
//func (taskcon *Taskcontrollers) CancelTaskGet() {
//	StatusTask(taskcon,4)
//}
//
////重试任务
//func (taskcon *Taskcontrollers) RetryTaskGet() {
//	StatusTask(taskcon,1)
//}
//
//
//
//
////删除任务
//func (taskcon *Taskcontrollers) DelTaskGet() {
//	beego.Error(1)
//	o := orm.NewOrm()
//	id, _ := strconv.Atoi(taskcon.GetString("id"))
//	task := module.Manage{Id:id}
//	if err := orm.NewOrm().QueryTable("Manage").Filter("Id",id).RelatedSel("User").One(&task);err != nil {
//		beego.Error(err)
//		taskcon.Redirect("/task/err",302)
//	}
//	//做判断 判断任务归属人的权限是否小于 删除任务人的权限 如果小于可以删除  或者如果是自己扇自己任务可以删除
//	//其他的提是
//	contextname := taskcon.GetSession("UserName").(string)
//	contextuser := module.User{Name:contextname}
//	err := o.Read(&contextuser,"Name")
//	if !module.ErrInfo("查询用户失败",err) {taskcon.Redirect("/task/err",302);return}
//	taskUser := task.User
//	//这里判断用户删除自己的或者是  删除比他权限小的可以直接删除 否则跳转权限报错页面
//	if contextuser.Role  > taskUser.Role  || taskUser.Id == contextuser.Id {
//		_,err = o.Delete(&task)
//		if !module.ErrInfo("删除失败",err) {taskcon.Redirect("/task/err",302);return}
//		taskcon.Redirect("/task/show",302)
//		return
//	}else {
//		taskcon.Redirect("/role/err",302)
//		return
//	}
//}
//
//
//
////根据我们需要的任务状态 做数据库修改
//func StatusTask(taskcon *Taskcontrollers,status int) {
//	o := orm.NewOrm()
//	id, _ := strconv.Atoi(taskcon.GetString("id"))
//
//	task := module.Manage{Id:id}
//	if err := o.QueryTable("Manage").Filter("Id",id).RelatedSel("User").One(&task);err != nil {
//		beego.Error(err)
//		taskcon.Redirect("/task/err",302)
//	}
//
//	//做判断 判断任务归属人的权限是否小于 删除任务人的权限 如果小于可以删除  或者如果是自己扇自己任务可以删除
//	//其他的提是
//	contextname := taskcon.GetSession("UserName").(string)
//	contextuser := module.User{Name:contextname}
//	err := o.Read(&contextuser,"Name")
//	if !module.ErrInfo("查询用户失败",err) {taskcon.Redirect("/task/err",302);return}
//	taskUser := task.User
//	//这里判断用户删除自己的或者是  删除比他权限小的可以直接删除 否则跳转权限报错页面
//
//	if contextuser.Role  > taskUser.Role  || taskUser.Id == contextuser.Id {
//		task.TaskStatus = status
//		_,err = o.Update(&task)
//		if !module.ErrInfo("删除失败",err) {taskcon.Redirect("/task/err",302);return}
//		taskcon.Redirect("/task/show?taskIndex=1",302)
//		return
//	}else {
//		taskcon.Redirect("/role/err",302)
//		return
//	}
//}
