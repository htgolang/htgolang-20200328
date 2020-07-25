package forms

import (
	"cmdb/models"

	"github.com/astaxie/beego/validation"
)

// DepartmentForm 部门表单
type DepartmentForm struct {
	ID       int    `form:"id" valid:"Numeric"`                  // 必须是数字;
	Name     string `form:"name" valid:"Required;MaxSize(64)"`   // 不可以为空;长度检查不可以超过64字节;名称唯一
	Describe string `form:"describe" valid:"MaxSize(1024)"`      // 长度不可以超过1024字节;
	Addr     string `form:"addr" valid:"Required;MaxSize(1024)"` // 不可以为空;长度检查不可以超过1024字节;
}

// Valid 做form表单自定义数据验证
func (d *DepartmentForm) Valid(v *validation.Validation) {
	// 检查name的唯一性
	if department := models.GetDepartmentsByName(d.Name); department != nil || department.ID == d.ID {
		v.SetError("Name", "department name is exist")
	}
}
