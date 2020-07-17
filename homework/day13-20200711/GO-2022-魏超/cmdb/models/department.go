package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Department))
}

// Department 部门对象
type Department struct {
	ID       int    `orm:"column(id);pk;"`
	Name     string `orm:"column(name);size(522)description(部门名称)"`
	Describe string `orm:"column(describe);size(1024);description(部门职能描述信息)"`
	Addr     string `orm:"column(addr);size(1024);description(部门地址)"`
}

// TableName 设置表名
func (d Department) TableName() string {
	return "department"
}

// AddDepartment 添加部门信息
func (d Department) AddDepartment() error {
	var err error
	_, err = orm.NewOrm().Insert(&d)
	return err
}

// DeleteDepartment 删除部门信息，条件没有员工数据此部门，才可以删除部门信息
func (d Department) DeleteDepartment() error {
	var (
		// userCount int64
		err error
	)
	// 交给 controllers 操作
	// userCount, _ = orm.NewOrm().QueryTable(&User{}).Filter("department_id", d.ID).Count()
	// if userCount != 0 {
	// 	return errors.New("Department exist employee info")
	// }
	_, err = orm.NewOrm().Delete(&d, "id")
	return err
}

// UpdateDepartment 更细部门信息
func (d Department) UpdateDepartment() error {
	var (
		err error
	)
	_, err = orm.NewOrm().Update(&d)
	return err
}

// GetDepartmentByID 通过部门ID获取部门信息
func (d *Department) GetDepartmentByID() error {
	return orm.NewOrm().Read(d, "id")
}

// QueryDepartments 根据部门名称和地址查询部门信息
func QueryDepartments(key string) []Department {
	var (
		departments []Department
		queryset    orm.QuerySeter
		cond        *orm.Condition
	)
	queryset = orm.NewOrm().QueryTable(&Department{})
	if key != "" {
		cond = orm.NewCondition()
		cond = cond.Or("name__icontains", key)
		cond = cond.Or("addr__icontains", key)
		queryset = queryset.SetCond(cond)
	}
	queryset.All(&departments)
	return departments
}
