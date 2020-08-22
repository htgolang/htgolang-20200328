package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/services"
	"html/template"

	"github.com/astaxie/beego/validation"
)

// PasswordController 用户修改密码控制器
type PasswordController struct {
	auth.LayoutController
}

// Modify 修改用户密码
func (c *PasswordController) Modify() {
	form := &forms.PasswordModifyForm{User: c.LoginUser}
	errs := errors.New()
	text := ""
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			// 验证
			valid := &validation.Validation{}
			if success, err := valid.Valid(form); err != nil {
				errs.Add("default", err.Error())
			} else if !success {
				errs.AddValidation(valid)
			} else {
				services.UserService.ModifyPassword(c.LoginUser.ID, form.Password)
				text = "修改密码成功"
			}
		}
	}
	c.TplName = "password/modify.html"
	c.Data["errors"] = errs
	c.Data["text"] = text
	c.Data["title"] = "修改用户密码"
	c.Data["xsrf_input"] = template.HTML(c.XSRFFormHTML())
}
