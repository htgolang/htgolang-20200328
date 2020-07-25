package forms

import "github.com/astaxie/beego/validation"

// AuthForm 登录表单
type AuthForm struct {
	Name     string `form:"name" valid:"Required"`
	Password string `form:"password" valid:"Required"`
}

// Valid 做form表单自定义数据验证
func (a AuthForm) Valid(v *validation.Validation) {
}
