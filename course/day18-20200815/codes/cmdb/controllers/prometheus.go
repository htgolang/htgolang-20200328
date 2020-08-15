package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/services"
	"net/http"

	"github.com/astaxie/beego"
)

type prometheusController struct {
	auth.LayoutController
}

func (c *prometheusController) Prepare() {
	c.LayoutController.Prepare()
	c.Data["nav"] = "prometheus"
	c.Data["subnav"] = c.GetNav()
}

type NodeController struct {
	prometheusController
}

func (c *NodeController) Query() {
	beego.ReadFromRequest(&c.Controller)

	q := c.GetString("q")
	c.Data["nodes"] = services.NodeService.Query(q)
	c.Data["q"] = q

	c.TplName = "prometheus/node/query.html"
}

func (c *NodeController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil {
		services.NodeService.Delete(pk)
	}
	c.Redirect(beego.URLFor("NodeController.Query"), http.StatusFound)
}

type JobController struct {
	prometheusController
}

func (c *JobController) Query() {
	beego.ReadFromRequest(&c.Controller)

	q := c.GetString("q")
	c.Data["jobs"] = services.JobService.Query(q)
	c.Data["q"] = q

	c.TplName = "prometheus/job/query.html"
}

func (c *JobController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil {
		services.JobService.Delete(pk)
	}
	c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
}

func (c *JobController) Create() {
	form := &forms.JobCreateForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			// 验证
			services.JobService.Create(form)
			c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
		}
	}
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["nodes"] = services.NodeService.Query("")
	c.TplName = "prometheus/job/create.html"
}

func (c *JobController) Modify() {
	// Get 查询值显示
	// POST 更新
	form := &forms.JobModifyForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			// 验证
			services.JobService.Modify(form)
			c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
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
	c.TplName = "prometheus/job/modify.html"
}

type TargetController struct {
	prometheusController
}

func (c *TargetController) Query() {
	beego.ReadFromRequest(&c.Controller)

	q := c.GetString("q")
	c.Data["targets"] = services.TargetService.Query(q)
	c.Data["q"] = q

	c.TplName = "prometheus/target/query.html"
}

func (c *TargetController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil {
		services.TargetService.Delete(pk)
	}
	c.Redirect(beego.URLFor("TargetController.Query"), http.StatusFound)
}

func (c *TargetController) Create() {
	form := &forms.TargetCreateForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			// 验证
			services.TargetService.Create(form)
			c.Redirect(beego.URLFor("TargetController.Query"), http.StatusFound)
		}
	}
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["jobs"] = services.JobService.Query("")
	c.TplName = "prometheus/target/create.html"
}

func (c *TargetController) Modify() {
	// Get 查询值显示
	// POST 更新
	form := &forms.TargetModifyForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			// 验证
			services.TargetService.Modify(form)
			c.Redirect(beego.URLFor("TargetController.Query"), http.StatusFound)
		}
	} else {
		if pk, err := c.GetInt("pk"); err == nil {
			if target := services.TargetService.GetByPk(pk); target != nil {
				form = forms.NewTargetModifyForm(target)
			}
		}
	}

	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["jobs"] = services.JobService.Query("")
	c.TplName = "prometheus/target/modify.html"
}
