package models

import (
	"time"

	"cmdb/utils"

	"github.com/astaxie/beego/orm"
)

type User struct {
	ID         int        `orm:"column(id)"`
	StaffID    string     `orm:"column(staff_id);size(32)"`
	Name       string     `orm:"size(64)"`
	Nickname   string     `orm:"size(64)"`
	Password   string     `orm:"size(1024)"`
	Gender     int        `orm:""`
	Tel        string     `orm:"size(32)"`
	Addr       string     `orm:"size(128)"`
	Email      string     `orm:"size(64)"`
	Department string     `orm:"size(128)"`
	Status     int        `orm:""`
	CreatedAt  *time.Time `orm:"auto_now_add"`
	UpdatedAt  *time.Time `orm:"auto_now"`
	DeletedAt  *time.Time `orm:"null"`
	IsAdmin    int        `orm:"column(isAdmin);default(0)"`
}

func init() {
	orm.RegisterModel(new(User))
	// orm.RunSyncdb("default", true, true)
}

func (u *User) ValidPassword(password string) bool {
	return utils.CheckPassword(password, u.Password)
}

// GenderText 性别显示
func (u *User) GenderText() string {
	if u.Gender == 1 {
		return "女"
	}
	return "男"
}

// StatusText 状态显示
func (u *User) StatusText() string {
	switch u.Status {
	case 0:
		return "正常"
	case 1:
		return "锁定"
	case 2:
		return "离职"
	}
	return "未知"
}

func GetNameById(id int) string {
	user := &User{ID: id}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "ID"); err == nil {
		return user.Name
	}
	return "未知用户"
}
