package forms

// 用户登陆表单
type LoginForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}
