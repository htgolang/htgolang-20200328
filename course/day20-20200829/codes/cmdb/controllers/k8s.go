package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/services"
	"net/http"

	"github.com/astaxie/beego"
)

type k8sController struct {
	auth.LayoutController
}

func (c *k8sController) Prepare() {
	c.LayoutController.Prepare()
	c.Data["nav"] = "k8s"
	c.Data["subnav"] = c.GetNav()
}

type DeploymentController struct {
	k8sController
}

func (c *DeploymentController) Query() {
	c.Data["deployments"] = services.DeploymentService.Query()
	c.TplName = "k8s/deployment/query.html"
}

func (c *DeploymentController) Delete() {
	name := c.GetString("name")
	namespace := c.GetString("namespace", "default")
	// 数据检查&权限
	services.DeploymentService.Delete(name, namespace)
	c.Redirect(beego.URLFor("DeploymentController.Query"), http.StatusFound)
}

func (c *DeploymentController) Create() {
	form := &forms.DeploymentCreateForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			// 验证
			services.DeploymentService.Create(form)
			c.Redirect(beego.URLFor("DeploymentController.Query"), http.StatusFound)
		}
	}
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["namespaces"] = services.NamespaceService.Query()
	c.TplName = "k8s/deployment/create.html"
}

func (c *DeploymentController) Modify() {
	// Get 查询值显示
	// POST 更新
	form := &forms.JobModifyForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			// 验证
			services.JobService.Modify(form)
			c.Redirect(beego.URLFor("DeploymentController.Query"), http.StatusFound)
		}
	} else {
		if pk, err := c.GetInt("pk"); err == nil {
			job := services.JobService.GetByPk(pk)
			form.ID = job.ID
			form.Key = job.Key
			form.Remark = job.Remark
			form.Node = job.Node.ID
		}
	}

	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["nodes"] = services.NodeService.Query("")
	c.TplName = "k8s/deployment/modify.html"
}
