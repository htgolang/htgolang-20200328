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
	Name     string `orm:"column(name);size(522);unique;description(部门名称)"`
	Describe string `orm:"column(describe);size(1024);description(部门职能描述信息)"`
	Addr     string `orm:"column(addr);size(1024);description(部门地址)"`
}

// TableName 设置表名
func (d Department) TableName() string {
	return "department"
}

// AddDepartment 添加部门信息
func AddDepartment(d Department) error {
	var err error
	_, err = orm.NewOrm().Insert(&d)
	return err
}

// DeleteDepartment 删除部门信息，条件没有员工数据此部门，才可以删除部门信息(硬删除)
func DeleteDepartment(id int) error {
	var (
		// userCount int64
		err error
	)
	_, err = orm.NewOrm().QueryTable(new(Department)).Filter("id", id).Delete()
	return err
}

// UpdateDepartment 更新部门信息
func UpdateDepartment(pk int, params orm.Params) error {
	var (
		err error
	)
	_, err = orm.NewOrm().QueryTable(new(Department)).Filter("id", pk).Update(params)
	return err
}

// GetDepartmentByID 通过部门ID获取部门信息
func GetDepartmentByID(id int) *Department {
	var (
		department = &Department{}
		err        error
	)
	err = orm.NewOrm().QueryTable(new(Department)).Filter("id", id).One(department)
	if err == nil {
		return department
	}
	return nil
}

// QueryDepartments 根据部门名称和地址查询部门信息
func QueryDepartments(key string) []*Department {
	var (
		departments []*Department
		queryset    orm.QuerySeter
		cond        *orm.Condition
	)
	queryset = orm.NewOrm().QueryTable(new(Department))
	if key != "" {
		cond = orm.NewCondition()
		cond = cond.Or("name__icontains", key)
		cond = cond.Or("addr__icontains", key)
		queryset = queryset.SetCond(cond)
	}
	queryset.All(&departments)
	return departments
}

// GetDepartmentsByName 根据部门名称查询是否存在
func GetDepartmentsByName(name string) *Department {
	var (
		department = &Department{}
		err        error
	)
	err = orm.NewOrm().QueryTable(new(Department)).Filter("name", name).One(department)
	if err == nil {
		return department
	}
	return nil

}
