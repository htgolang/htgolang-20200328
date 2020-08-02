package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/services"

	"github.com/astaxie/beego/validation"
)

type PasswordController struct {
	auth.AuthController
}

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
	c.Data["xsrf_token"] = c.XSRFToken()

	// form := &forms.PasswordModifyForm{}
	// errs := errors.New()
	// text := ""

	// if c.Ctx.Input.IsPost() {
	// 	if err := c.ParseForm(form); err == nil {
	// 		// //验证
	// 		// if ok := c.LoginUser.ValidatePassword(form.OldPassword); !ok {
	// 		// 	errs.Add("default", "旧密码错误")
	// 		// } else {
	// 		valid := &validation.Validation{}
	// 		if success, err := valid.Valid(form); err != nil {
	// 			errs.Add("default", err.Error())
	// 		} else if !success {
	// 			errs.AddValidation(valid)
	// 		} else {
	// 			services.UserService.ModifyPassword(c.LoginUser.ID, form.Password)
	// 			text = "修改成功"
	// 		}
	// 		// 第二版
	// 		// passwordRegex := "^[0-9a-zA-Z_.\\$\\!#%^&\\(\\)\\+]{6,20}$"
	// 		// valid := &validation.Validation{}
	// 		// valid.Match(form.Password, regexp.MustCompile(passwordRegex), "default.default.default").Message("密码格式不正确")

	// 		// if valid.HasErrors() {
	// 		// 	for key, errors := range valid.ErrorsMap {
	// 		// 		for _, err := range errors {
	// 		// 			errs.Add(key, err.Message)
	// 		// 		}
	// 		// 	}
	// 		// } else if form.Password != form.Password2 {
	// 		// 	errs.Add("default", "两次密码不一致")
	// 		// } else if form.OldPassword == form.Password {
	// 		// 	errs.Add("default", "新旧密码不能一致")
	// 		// } else {
	// 		// 	models.ModifyUserPassword(c.LoginUser.ID, form.Password)
	// 		// 	text = "修改成功"
	// 		// }

	// 		// 第一版
	// 		// passwordRegex := "^[0-9a-zA-Z_.\\$\\!#%^&\\(\\)\\+]{6,20}$"
	// 		// if isMatch, _ := regexp.MatchString(passwordRegex, form.Password); !isMatch {
	// 		// 	errs.Add("default", "密码组成错误")
	// 		// } else if form.Password != form.Password2 {
	// 		// 	errs.Add("default", "两次密码不一致")
	// 		// } else if form.OldPassword == form.Password {
	// 		// 	errs.Add("default", "新旧密码不能一致")
	// 		// } else {
	// 		// 	models.ModifyUserPassword(c.LoginUser.ID, form.Password)
	// 		// 	text = "修改成功"
	// 		// }
	// 		// }
	// 	}
	// }

	// c.TplName = "password/modify.html"
	// c.Data["errors"] = errs
	// c.Data["text"] = text

}
