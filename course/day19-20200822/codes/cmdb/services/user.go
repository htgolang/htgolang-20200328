package services

import (
	"github.com/astaxie/beego/orm"

	"cmdb/forms"
	"cmdb/models"
	"cmdb/utils"
)

type userService struct {
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
func (s *userService) Modify(form *forms.UserModifyForm) {
	if user := s.GetByPk(form.ID); user != nil {
		user.Name = form.Name
		ormer := orm.NewOrm()
		ormer.Update(user, "Name")
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
