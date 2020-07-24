package models

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"cmdb/forms"
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
	Flag       int        `orm:"default(0)"`
}

func init() {
	orm.RegisterModel(new(User))
	// orm.RunSyncdb("default", true, true)
}

func (u *User) ValidatePassword(password string) bool {
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

func GetAllAccounts() []*User {
	user := new(User)
	var accounts []*User

	ormer := orm.NewOrm()
	_, err := ormer.QueryTable(user).All(&accounts, "ID", "Name")
	if err != nil {
		log.Fatal(err.Error())
	}
	return accounts
}
func GetUserByPk(pk int) *User {
	user := &User{ID: pk}
	ormer := orm.NewOrm()
	if err := ormer.Read(user); err == nil {
		return user
	}
	return nil
}

func GetUserByName(name string) *User {
	user := &User{Name: name}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}

	return nil
}

func GetUserNameById(id int) string {
	user := &User{ID: id}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "ID"); err == nil {
		return user.Name
	}
	return "未知用户"

}

func GenerateStaffID() string {
	user := new(User)
	ormer := orm.NewOrm()
	ormer.QueryTable(user).OrderBy("-StaffID").One(user)
	sid, _ := strconv.Atoi(strings.ReplaceAll(user.StaffID, "T", ""))
	sid++
	return fmt.Sprintf("T%05d", sid)
}

func QueryUser(key string) []*User {
	var users []*User
	qs := orm.NewOrm().QueryTable(&User{})

	if key != "" && key != "all" {
		cond := orm.NewCondition()
		cond = cond.Or("id__iexact", key).Or("name__icontains", key).Or("tel__icontains", key).Or("addr__icontains", key).Or("email__icontains", key)
		qs.SetCond(cond).Filter("deleted_at__isnull", true).All(&users)
	} else {
		qs.Filter("deleted_at__isnull", true).All(&users)
	}
	return users
}

func ModifyUser(form *forms.FormUser) {
	if user := GetUserByPk(form.ID); user != nil {
		user.Name = form.Name
		user.Nickname = form.Nickname
		user.Gender = form.Gender
		user.Tel = form.Tel
		user.Addr = form.Addr
		user.Email = form.Email
		user.Department = form.Department
		ormer := orm.NewOrm()
		ormer.Update(user)
	}
}

func DeleteUser(id int) {
	now := time.Now()
	user := &User{ID: id}
	ormer := orm.NewOrm()
	if ormer.Read(user) == nil {
		user.Flag = 1
		user.DeletedAt = &now
		// _, err := ormer.Update(user)
		ormer.Update(user)
	}
}
