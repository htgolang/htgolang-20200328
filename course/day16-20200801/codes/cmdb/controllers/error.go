package controllers

import (
	"cmdb/base/controllers/base"
)

// ErrorController 错误处理控制器
type ErrorController struct {
	base.BaseController
}

// Error开头的方法
// Error404 => 处理404错误的逻辑
func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
}

// Error404 => 处理404错误的逻辑
func (c *ErrorController) ErrorNotPermission() {
	c.TplName = "error/not_permission.html"
}
