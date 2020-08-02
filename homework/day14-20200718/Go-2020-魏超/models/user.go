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
	StaffID       string     `orm:"column(staff_id);unique;description(工号)"`
	Name          string     `orm:"column(name);size(25);description(姓名)"`
	NickName      string     `orm:"column(nickname);size(25);description(别名)"`
	Password      string     `orm:"column(password);size(1024);description(密码)"`
	Gender        int        `orm:"column(gender);size(1);description(性别)"`
	Tel           string     `orm:"column(tel);size(16);unique;description(联系方式)"`
	Email         string     `orm:"column(email);size(522);unique;description(工作邮箱)"`
	Addr          string     `orm:"column(addr);size(1024);description(联系地址)"`
	DepartmentID  int        `orm:"column(department_id);description(部门ID)"`
	Title         string     `orm:"column(title);size(64);description(职位)"`
	EntryTime     *time.Time `orm:"column(entry_time);type(datetime);null;description(入职时间)"`
	DimissionTime *time.Time `orm:"column(dimission_time);type(datetime);null;description(离职时间)"`
	// RoleID        int        `orm:"column(role_id);description(角色ID)"`
	StatusID int        `orm:"column(status_id);size(2);description(状态)"`
	DeleteAt *time.Time `orm:"column(delete_at);type(datetime);null;description(删除时间)"`
}

// TableName 设置表名
func (u *User) TableName() string {
	return "users"
}

// AddUser 添加用户信息
func AddUser(u User) error {
	var (
		err error
	)
	_, err = orm.NewOrm().Insert(&u)
	return err
}

// UpdateUser 更新用户信息
func UpdateUser(pk int, params orm.Params) error {
	var (
		err error
	)
	_, err = orm.NewOrm().QueryTable(new(Task)).Filter("delete_at__isnull", false).Filter("id", pk).Update(params)
	return err
}

// DeleteUser 删除用户
func DeleteUser(id int) error {
	var (
		err error
		now = time.Now()
	)
	_, err = orm.NewOrm().QueryTable(new(User)).Filter("delete_at__isnull", false).Filter("id", id).Update(orm.Params{"delete_at": &now})
	return err
}

// GetUserByID 根据用户ID获取用户数据
func GetUserByID(id int) *User {
	var (
		u   = &User{}
		err error
	)
	err = orm.NewOrm().QueryTable(new(User)).Filter("delete_at__isnull", false).Filter("id", id).One(u)
	if err == nil {
		return u
	}
	return nil
}

// GetUserByStaffID 根据用户StaffID获取用户数据
func GetUserByStaffID(staffid string) *User {
	var (
		user = &User{}
		err  error
	)
	err = orm.NewOrm().QueryTable(new(User)).Filter("delete_at__isnull", false).Filter("staff_id", staffid).One(user)
	if err == nil {
		return user
	}
	return nil
}

// GetUserByNickname 根据用户ID获取用户数据
func GetUserByNickname(nickname string) *User {
	var (
		user = &User{}
		err  error
	)
	err = orm.NewOrm().QueryTable(new(User)).Filter("delete_at__isnull", false).Filter("nickname", nickname).One(user)
	if err == nil {
		return user
	}
	return nil
}

// GetUserByTel 根据用户ID获取用户数据
func GetUserByTel(tel string) *User {
	var (
		user = &User{}
		err  error
	)
	err = orm.NewOrm().QueryTable(new(User)).Filter("delete_at__isnull", false).Filter("tel", tel).One(user)
	if err == nil {
		return user
	}
	return nil
}

// GetUserByEmail 根据用户ID获取用户数据
func GetUserByEmail(email string) *User {
	var (
		user = &User{}
		err  error
	)
	err = orm.NewOrm().QueryTable(new(User)).Filter("delete_at__isnull", false).Filter("email", email).One(user)
	if err == nil {
		return user
	}
	return nil
}

// GetUserByDepartmentID 根据DepartmentID获取用户数据
func GetUserByDepartmentID(id int) []*User {
	var (
		users = []*User{}
		err   error
	)
	_, err = orm.NewOrm().QueryTable(new(User)).Filter("delete_at__isnull", false).Filter("department_id", id).All(&users)
	if err == nil {
		return users
	}
	return nil
}

// QueryUsers 可以根据关键字查询信息
func QueryUsers(key string) []User {
	var (
		users    []User
		queryset orm.QuerySeter
		cond     = orm.NewCondition()
	)
	queryset = orm.NewOrm().QueryTable(&User{}).Filter("delete_at__isnull", false)
	if key != "" {
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
