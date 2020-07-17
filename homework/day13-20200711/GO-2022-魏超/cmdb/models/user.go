package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	// SexMap 性别map
	SexMap = map[int]string{0: "女", 1: "男"}
	// UserStatusMap 用户状态map
	UserStatusMap = map[int]string{0: "在职", 1: "离职", 2: "锁定"}
)

func init() {
	orm.RegisterModel(new(User))
}

// User 用户对象
type User struct {
	ID            int        `orm:"column(id);pk,auto"`
	StaffID       int        `orm:"column(staff_id);description(工号)"`
	Name          string     `orm:"column(name);size(25);description(姓名)"`
	Nickname      string     `orm:"column(nickname);size(25);description(别名)"`
	Password      string     `orm:"column(password);size(1024);description(密码)"`
	Gender        int        `orm:"column(gender);size(1);description(性别)"`
	Tel           string     `orm:"column(tel);size(16);description(联系方式)"`
	Email         string     `orm:"column(email);size(522);description(工作邮箱)"`
	Addr          string     `orm:"column(addr);size(1024);description(联系地址)"`
	DepartmentID  int        `orm:"column(department_id);description(部门ID)"`
	Title         string     `orm:"column(title);size(64);description(职位)"`
	EntryTime     *time.Time `orm:"column(entry_time);type(datetime);description(入职时间);null"`
	DimissionTime *time.Time `orm:"column(dimission_time);type(datetime);description(离职时间);null"`
	RoleID        int        `orm:"column(role_id);description(角色ID)"`
	Status        int        `orm:"column(status);size(2);description(状态)"`
}

// TableName 设置表名
func (u *User) TableName() string {
	return "users"
}

// UpdateUser 更新用户信息
func (u User) UpdateUser() error {
	var (
		err error
	)
	_, err = orm.NewOrm().Update(&u)
	return err
}

// DeleteUser 删除用户
func (u User) DeleteUser() error {
	var (
		err error
	)
	_, err = orm.NewOrm().Delete(&u, "id")
	return err
}

// GetUserByID 根据用户ID获取用户数据
func (u *User) GetUserByID() error {
	return orm.NewOrm().Read(u, "id")
}

// GetUserByNickname 根据用户ID获取用户数据
func (u *User) GetUserByNickname() error {
	return orm.NewOrm().Read(u, "nickname")
}

// GetUserByTel 根据用户ID获取用户数据
func (u *User) GetUserByTel() error {
	return orm.NewOrm().Read(u, "tel")
}

// GetUserByEmail 根据用户ID获取用户数据
func (u *User) GetUserByEmail() error {
	return orm.NewOrm().Read(u, "email")
}

// QueryUsers 可以根据关键字查询信息
func QueryUsers(key string) []User {
	var (
		users    []User
		queryset orm.QuerySeter
		cond     *orm.Condition
	)
	queryset = orm.NewOrm().QueryTable(&User{})
	if key != "" {
		cond = orm.NewCondition()
		cond = cond.Or("staff_id__icontains", key)
		cond = cond.Or("name__icontains", key)
		cond = cond.Or("nickname__icontains", key)
		cond = cond.Or("tel__icontains", key)
		cond = cond.Or("email__icontains", key)
		cond = cond.Or("addr__icontains", key)
		cond = cond.Or("title__icontains", key)
		queryset.SetCond(cond)
	}
	queryset.All(&users)
	return users
}
