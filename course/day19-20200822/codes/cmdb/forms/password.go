package forms

import (
	"regexp"

	"github.com/astaxie/beego/validation"

	"cmdb/models"
)

// PasswordModifyForm 修改用户密码
type PasswordModifyForm struct {
	User        *models.User `form:"-"`
	OldPassword string       `form:"old_password"`
	Password    string       `form:"password"`
	Password2   string       `form:"password2"`
}

// Valid 数据检查
func (f *PasswordModifyForm) Valid(validation *validation.Validation) {
	if ok := f.User.ValidPassword(f.OldPassword); !ok {
		validation.AddError("default.default", "旧密码错误")
		return
	}

	// 验证密码范围数字，大小写英文字母、特殊字符(_.$!#%^&*()+)
	passwordRegex := "^[0-9a-zA-Z_.\\$\\!#%^&\\*\\(\\)\\+]{6,20}$"
	validation.Match(f.Password, regexp.MustCompile(passwordRegex), "default.default.default").Message("密码格式不正确")
	if validation.HasErrors() {
		return
	} else if f.Password != f.Password2 {
		validation.AddError("default.default", "两次密码不一致")
	} else if f.OldPassword == f.Password {
		validation.AddError("default.default", "新旧密码不能一致")
	}
}
