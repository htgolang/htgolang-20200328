package routers

// 先到内置
// 第三方
// 当前项目

import (
	"github.com/astaxie/beego"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"cmdb/controllers"
	v1 "cmdb/controllers/api/v1"
	"cmdb/filters"
)

func init() {

	beego.InsertFilter("/*", beego.BeforeExec, filters.BeforeExecute)
	beego.InsertFilter("/*", beego.AfterExec, filters.AfterExecute, false)

	beego.Handler("/metrics", promhttp.Handler())

	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.HomeController{}, "*:Index")

	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.PasswordController{})

	// prometheus
	beego.AutoRouter(&controllers.NodeController{})
	beego.AutoRouter(&controllers.JobController{})
	beego.AutoRouter(&controllers.TargetController{})
	beego.AutoRouter(&controllers.AlertController{})

	// k8s
	beego.AutoRouter(&controllers.DeploymentController{})

	// /v1/
	v1 := beego.NewNamespace("/v1", beego.NSAutoRouter(&v1.PrometheusController{}))
	beego.AddNamespace(v1)
}
