package forms

// LoginForm 用户登录表单
type LoginForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}
