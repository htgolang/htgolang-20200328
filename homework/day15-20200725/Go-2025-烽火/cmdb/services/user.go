package services

import (
	"cmdb/forms"
	"cmdb/models"
	"cmdb/utils"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type userService struct {
}

//获取所有用户信息
func (s *userService) GetAllAccounts() []*models.User {
	var accounts []*models.User

	ormer := orm.NewOrm()
	_, err := ormer.QueryTable(&models.User{}).All(&accounts, "ID", "Name")
	if err != nil {
		log.Fatal(err.Error())
	}
	return accounts
}

// GetByPk 通过用户ID获取用户信息
func (s *userService) GetByPk(pk int) *models.User {
	user := &models.User{ID: pk}
	ormer := orm.NewOrm()
	if err := ormer.Read(user); err == nil {
		return user
	}
	return nil
}

// GetByName 通过用户名获取用户
func (s *userService) GetByName(name string) *models.User {
	user := &models.User{Name: name}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

func getNewStaffID() string {
	user := new(models.User)
	ormer := orm.NewOrm()
	ormer.QueryTable(user).OrderBy("-StaffID").One(user)
	sid, _ := strconv.Atoi(strings.ReplaceAll(user.StaffID, "T", ""))
	sid++
	return fmt.Sprintf("T%05d", sid)
}

// Add 添加用户
func (s *userService) Add(form *forms.FormUser) error {
	user := &models.User{
		StaffID: getNewStaffID(),
	}
	user.Name = form.Name
	user.Nickname = form.Nickname
	user.Password = utils.GeneratePassword(form.Password)
	user.Gender = form.Gender
	user.Tel = form.Tel
	user.Addr = form.Addr
	user.Email = form.Email
	user.Department = form.Department
	ormer := orm.NewOrm()
	_, err := ormer.Insert(user)
	return err
}

// Query 查询用户
func (s *userService) Query(q string) []*models.User {
	var users []*models.User
	queryset := orm.NewOrm().QueryTable(&models.User{})
	if q != "" {
		cond := orm.NewCondition()
		cond = cond.Or("name__icontains", q)
		cond = cond.Or("nickname__icontains", q)
		cond = cond.Or("tel__icontains", q)
		cond = cond.Or("addr__icontains", q)
		cond = cond.Or("email__icontains", q)
		cond = cond.Or("department__icontains", q)
		queryset = queryset.SetCond(cond)
	}
	queryset.All(&users)
	return users
}

// Modify 修改用户信息
func (s *userService) Modify(form *forms.FormUser) {
	if user := s.GetByPk(form.ID); user != nil {
		user.Name = form.Name
		user.Nickname = form.Nickname
		user.Gender = form.Gender
		user.Tel = form.Tel
		user.Addr = form.Addr
		user.Email = form.Email
		user.Department = form.Department
		ormer := orm.NewOrm()
		ormer.Update(user, "Name", "Nickname", "Gender", "Tel", "Addr", "Email", "Department")
	}
}

// Delete 删除用户 Delete
func (s *userService) Delete(pk int) {
	ormer := orm.NewOrm()
	ormer.Delete(&models.User{ID: pk})
}

// ModifyPassword 修改用户密码
func (s *userService) ModifyPassword(pk int, password string) {
	if user := s.GetByPk(pk); user != nil {
		user.Password = utils.GeneratePassword(password)
		ormer := orm.NewOrm()
		ormer.Update(user, "Password")
	}
}

// UserService 用户操作服务
var UserService = new(userService)
