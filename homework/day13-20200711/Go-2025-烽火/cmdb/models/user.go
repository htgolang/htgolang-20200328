package models

import (
	"cmdb/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	ID         int        `form:"id" orm:"column(id)"`
	StaffID    string     `form:"staff_id" orm:"column(staff_id);size(32)"`
	Name       string     `form:"name" orm:"size(64)"`
	Nickname   string     `form:"nickname" orm:"size(64)"`
	Password   string     `form:"password" orm:"size(1024)"`
	Gender     int        `form:"gender" orm:""`
	Tel        string     `form:"tel" orm:"size(32)"`
	Addr       string     `form:"addr" orm:"size(128)"`
	Email      string     `form:"email" orm:"size(64)"`
	Department string     `form:"department" orm:"size(128)"`
	Status     int        `form:"status" orm:""`
	CreatedAt  *time.Time `form:"created_at" orm:"auto_now_add"`
	UpdatedAt  *time.Time `form:"updated_at" orm:"auto_now"`
	DeletedAt  *time.Time `form:"deleted_at" orm:"null"`
	Flag       int        `form:"flag" orm:"default(0)"`
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
	// fmt.Printf("%#v\n", user)
	sid, _ := strconv.Atoi(strings.ReplaceAll(user.StaffID, "T", ""))
	sid++
	return fmt.Sprintf("T%05d", sid)
}

func QueryUser(key string) []*User {
	var users = []*User{}
	if key != "" {
		qs := orm.NewOrm().QueryTable(&User{})
		cond := orm.NewCondition()
		cond = cond.Or("id__iexact", key).Or("name__icontains", key).Or("tel__icontains", key).Or("addr__icontains", key).Or("email__icontains", key)
		qs.SetCond(cond).All(&users)
	}
	return users
}

func GetUserById(id string) *User {
	uid, _ := strconv.Atoi(id)
	user := &User{ID: uid}
	ormer := orm.NewOrm()
	ormer.Read(user)
	return user
}

func ModifyUser(user *User) {
	ormer := orm.NewOrm()
	ormer.Update(user)
}

func DeleteUser(id string) {
	uid, _ := strconv.Atoi(id)
	user := &User{ID: uid}
	ormer := orm.NewOrm()
	ormer.Delete(user)
}
