package forms

// UserModifyForm 用户修改表单
type UserModifyForm struct {
	ID   int    `form:"id"`
	Name string `form:"name"`
}
