package forms

import (
	"cmdb/models"

	"regexp"

	"github.com/astaxie/beego/validation"
)

type PasswordModifyForm struct {
	User        *models.User `form:"_"`
	OldPassword string       `form:"old_password"`
	Password    string       `form:"password"`
	Password2   string       `form:"password2"`
}

func (f *PasswordModifyForm) Valid(valid *validation.Validation) {
	if ok := f.User.ValidPassword(f.OldPassword); !ok {
		valid.AddError("default.default", "旧密码错误")
		return
	}
	passwordRegex := "^[0-9a-zA-Z_.]{6,20}$"
	valid.Match(f.Password, regexp.MustCompile(passwordRegex), "default.default.default").Message("密码格式不正确")

	if valid.HasErrors() {
		return
	} else if f.Password != f.Password2 {
		valid.AddError("default.default", "两次密码不一致")
	} else if f.OldPassword == f.Password {
		valid.AddError("default.defailt", "新旧密码不能一致")
	}
}
