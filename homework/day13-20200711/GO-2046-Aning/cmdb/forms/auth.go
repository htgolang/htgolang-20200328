package forms

type LoginForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}
