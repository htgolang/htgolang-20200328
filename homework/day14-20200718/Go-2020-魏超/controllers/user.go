package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/models"
	"cmdb/utils"
	"net/http"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

// Usercontroller 用户信息控制器
type Usercontroller struct {
	auth.AuthorizationController
}

// Index 查询用户信息
func (c *Usercontroller) Index() {
	c.Data["Users"] = models.QueryUsers(c.GetString("key"))
	c.TplName = "user/index.html"
}

// Delete 删除用户信息
func (c *Usercontroller) Delete() {
	if userID, err := c.GetInt("uid"); err == nil && c.LoginUser.ID != userID {
		err = models.DeleteUser(userID)
	}
	c.Redirect(c.URLFor("Usercontroller.Index"), http.StatusFound)
}

// Modify 修改用户信息
func (c *Usercontroller) Modify() {
	var (
		userForm forms.UserForm
		errors   map[string]string
	)
	if c.Ctx.Input.IsGet() {
		if userID, err := c.GetInt("pk"); err != nil {
			if user := models.GetUserByID(userID); user != nil {
				userForm.ID = user.ID
				userForm.StaffID = user.StaffID
				userForm.Name = user.Name
				userForm.NickName = user.NickName
				userForm.Gender = user.Gender
				userForm.Tel = user.Tel
				userForm.Email = user.Email
				userForm.Addr = user.Addr
				userForm.DepartmentID = user.DepartmentID
				userForm.Title = user.Title
				userForm.StatusID = user.StatusID
			}
		}
	} else if c.Ctx.Input.IsPost() {
		err := c.ParseForm(userForm)
		if err != nil {
			logs.Warn("parse form UserForm error, %s", err)
			errors["submit"] = "server error"
		} else {
			valid := validation.Validation{}
			if ok, err := valid.Valid(&userForm); err != nil {
				logs.Warn("valid handle error, %s", err)
				errors["submit"] = "server error"
			} else if !ok {
				for _, err := range valid.Errors {
					errors[err.Key] = err.Message
				}
			} else {
				var dimissionTime *time.Time
				if models.UserStatusMap[userForm.StatusID] == "离职" {
					now := time.Now()
					dimissionTime = &now
				}

				if err := models.UpdateUser(userForm.ID, orm.Params{
					"staff_id":       userForm.StaffID,
					"name":           userForm.Name,
					"nickname":       userForm.NickName,
					"gender":         userForm.Gender,
					"tel":            userForm.Tel,
					"email":          userForm.Email,
					"addr":           userForm.Addr,
					"department_id":  userForm.DepartmentID,
					"title":          userForm.Title,
					"dimission_time": dimissionTime,
					"status_id":      userForm.StatusID,
				},
				); err == nil {
					c.Redirect(c.URLFor("Usercontroller.Index"), http.StatusFound)
				} else {
					logs.Error("update data error, %s", err)
					errors["submit"] = "update data error"
				}
			}
		}
	}
	c.Data["Errors"] = errors
	c.Data["User"] = userForm
	c.TplName = "user/edit.html"
}

// Add 添加用户信息
func (c *Usercontroller) Add() {
	var (
		userForm forms.AddUserForm
		errors   map[string]string
	)
	if c.Ctx.Input.IsPost() {
		err := c.ParseForm(userForm)
		if err != nil {
			logs.Warn("parse form AddUserForm error, %s", err)
			errors["submit"] = "server error"
		} else {
			valid := validation.Validation{}
			if ok, err := valid.Valid(&userForm); err != nil {
				logs.Warn("valid handle error, %s", err)
				errors["submit"] = "server error"
			} else if !ok {
				for _, err := range valid.Errors {
					errors[err.Key] = err.Message
				}
			} else {
				user := models.User{
					StaffID:      userForm.StaffID,
					Name:         userForm.Name,
					NickName:     userForm.NickName,
					Gender:       userForm.Gender,
					Tel:          userForm.Tel,
					Email:        userForm.Email,
					Addr:         userForm.Addr,
					DepartmentID: userForm.DepartmentID,
					Title:        userForm.Title,
					StatusID:     2, // 初始添加的用户都为锁定状态
					Password:     utils.HashPasswd(userForm.Password),
				}
				if err := models.AddUser(user); err == nil {
					c.Redirect(c.URLFor("Usercontroller.Index"), http.StatusFound)
				} else {
					logs.Error("update data error, %s", err)
					errors["submit"] = "update data error"
				}
			}
		}
	}
	c.Data["Errors"] = errors
	c.Data["Department"] = userForm
	c.TplName = "user/edit.html"
}

// Password 设置用户密码
func (c *Usercontroller) Password() {
	var (
		passwordForm forms.UserPassword
		errors       map[string]string
	)
	if uid, err := c.GetInt("uid"); err == nil && uid == c.LoginUser.ID {
		if c.Ctx.Input.IsPost() {
			valid := validation.Validation{}
			if ok, err := valid.Valid(&passwordForm); err != nil {
				logs.Warn("valid handle error, %s", err)
				errors["submit"] = "server error"
			} else if !ok {
				for _, err := range valid.Errors {
					errors[err.Key] = err.Message
				}
			} else {
				if utils.ValidPassword(passwordForm.OldPassword, c.LoginUser.Password) {
					if err := models.UpdateUser(uid, orm.Params{
						"password": utils.HashPasswd(passwordForm.NewPassword),
					}); err == nil {
						c.Redirect(c.URLFor("Usercontroller.Modify", "uid", uid), http.StatusFound)
					}
				}
			}
		}
		c.Data["User_id"] = uid
		c.TplName = "user/password.html"
	}
	c.Redirect(c.URLFor("Usercontroller.Index"), http.StatusFound)
}
