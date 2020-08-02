package forms

import (
	"cmdb/models"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/validation"
)

var (
	// UserIllegalChar 用户非法字符
	UserIllegalChar = `\|/'":*? $:[]()+=-~,.<>`
)

// UserForm 用户表单
type UserForm struct {
	ID           int    `form:"id" valid:"Numeric"`                            // 不可以为空;必须是数字;检查数据的有效性;
	StaffID      string `form:"staff_id" valid:"Required;Match(/^SC//d{8}$/)"` // 不能为空;输入的字符必须是SC开头;长度为10;数据唯一;
	Name         string `form:"name" valid:"MaxSize(64)"`                      // 不能包含特殊字符;长度不可以超过64字节;
	NickName     string `form:"nickname" valid:"AlphaNumeric;MaxSize(64)"`     // 不能包含特殊字符;长度不可以超过64字节;数据唯一;
	Gender       int    `form:"gender" valid:"Numeric"`                        // 必须是数字;
	Tel          string `form:"tel" valid:"Tel"`                               // 手机号号码格式检查;数据唯一;
	Email        string `form:"email" valid:"Email"`                           // 邮箱格式检查;数据唯一;
	Addr         string `form:"addr" valid:"MaxSize(1024)"`                    // 长度不可以超过1024字节;
	DepartmentID int    `form:"department_id" valid:"Numeric"`                 // 必须是数字;检查数据的有效性;
	EntryTime    string `form:"entry_time"`
	Title        string `form:"title" valid:"MaxSize(512)"` // 长度不可以超过64字节
	StatusID     int    `form:"status_id" valid:"Numeric"`  // 必须是数字;
}

// Valid 做form表单数据验证
func (u *UserForm) Valid(v *validation.Validation) {
	// 检查工号的唯一性
	if user := models.GetUserByStaffID(u.StaffID); user == nil || user.ID != u.ID {
		v.SetError("StaffID", "staffid is exist")
	}

	// 检查name是否存在特殊字符
	if strings.ContainsAny(u.Name, UserIllegalChar) {
		v.SetError("Name", fmt.Sprintf("name is not allowed to contain %s", UserIllegalChar))
	}

	// 检查别名的唯一性
	if user := models.GetUserByNickname(u.NickName); user == nil || user.ID != u.ID {
		v.SetError("StaffID", "nickname is exist")
	}

	// 检查gender有效性
	if _, ok := models.SexMap[u.Gender]; !ok {
		v.SetError("gender", "gender code illegal")
	}

	// 检查Tel的唯一性
	if user := models.GetUserByTel(u.Tel); user == nil || user.ID != u.ID {
		v.SetError("Tel", "Tel is exist")
	}

	// 检查email的唯一性
	if user := models.GetUserByEmail(u.Email); user == nil || user.ID != u.ID {
		v.SetError("Email", "Email is exist")
	}

	// 检查EntryTime的有效性
	if u.EntryTime != "" {
		if _, err := time.Parse(DateTimeLayout, u.EntryTime); err != nil {
			v.SetError("EntryTime", "datetime format err")
		}
	}

	// 检查StatusID的有效性
	if _, ok := models.UserStatusMap[u.StatusID]; ok {
		v.SetError("StatusID", "status code illegal")
	}
}

// AddUserForm 用户表单
type AddUserForm struct {
	UserForm
	Password      string `form:"password" valid:"Required;MinSize(6)"`
	ConfirmPasswd string `form:"confirm_password" valid:"Required"`
}

// Valid 做form表单数据验证
func (u AddUserForm) Valid(v *validation.Validation) {
	u.UserForm.Valid(v)
	// 比对校验密码
	if u.Password != u.ConfirmPasswd {
		v.SetError("ConfirmPasswd", "password not the same")
	}
}

// UserPassword 用户密码表单
type UserPassword struct {
	OldPassword     string `form:"old_password" valid:"Required"`
	NewPassword     string `form:"new_password" valid:"Required;MinSize(6)"`
	ConfirmPassword string `form:"confirm_password" valid:"Required"`
}

// Valid 做form表单数据验证
func (u UserPassword) Valid(v *validation.Validation) {
	if u.NewPassword != u.ConfirmPassword {
		v.SetError("ConfirmPassword", "password not the same")
	}
}
