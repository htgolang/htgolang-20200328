package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/models"
	"net/http"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

// DepartmentController 部门信息控制器
type DepartmentController struct {
	auth.AuthorizationController
}

// Index 展示部门信息
func (c *DepartmentController) Index() {
	c.Data["Tasks"] = models.QueryTasks(c.GetString("key"))
	c.TplName = "department/index.html"
}

// Delete 删除部门信息
func (c *DepartmentController) Delete() {
	var (
		departmentID int
		err          error
	)
	departmentID, err = c.GetInt("id")
	if err == nil {
		err = models.DeleteUser(departmentID)
	}
	c.Redirect(c.URLFor("DepartmentController.Index"), http.StatusFound)
}

// Modify 修改部门信息
func (c *DepartmentController) Modify() {
	var (
		errors         map[string]string
		departmentForm forms.DepartmentForm
		valid          = validation.Validation{}
	)
	if c.Ctx.Input.IsGet() {
		if departmentID, err := c.GetInt("id"); err == nil {
			if department := models.GetDepartmentByID(departmentID); department != nil {
				departmentForm.ID = department.ID
				departmentForm.Name = department.Name
				departmentForm.Describe = department.Describe
				departmentForm.Addr = department.Addr
			}
		}
	} else if c.Ctx.Input.IsPost() {
		err := c.ParseForm(departmentForm)
		if err == nil {
			if ok, err := valid.Valid(&departmentForm); err != nil {
				logs.Warn("valid handle error, %s", err)
				errors["submit"] = "server error"
			} else if !ok {
				for _, err := range valid.Errors {
					errors[err.Key] = err.Message
				}
			} else {
				if err := models.UpdateDepartment(departmentForm.ID, orm.Params{
					"name":     departmentForm.Name,
					"describe": departmentForm.Describe,
					"addr":     departmentForm.Addr},
				); err == nil {
					c.Redirect(c.URLFor("DepartmentController.Index"), http.StatusFound)
				} else {
					logs.Error("update data error, %s", err)
					errors["submit"] = "update data error"
				}
			}
		} else {
			logs.Warn("parse form DepartmentForm error, %s", err)
			errors["submit"] = "server error"
		}
	}

	c.Data["Errors"] = errors
	c.Data["Department"] = departmentForm
	c.TplName = "department/edit.html"
}

// Add 添加部门信息
func (c *DepartmentController) Add() {
	var (
		errors         map[string]string
		departmentForm forms.DepartmentForm
	)

	if c.Ctx.Input.IsPost() {
		err := c.ParseForm(departmentForm)
		if err == nil {

			// 表单数据验证
			valid := validation.Validation{}
			if ok, err := valid.Valid(&departmentForm); err != nil {
				logs.Warn("valid handle error, %s", err)
				errors["submit"] = "server error"
			} else if !ok {
				for _, err := range valid.Errors {
					errors[err.Key] = err.Message
				}
			} else {
				department := models.Department{
					Name:     departmentForm.Name,
					Describe: departmentForm.Describe,
					Addr:     departmentForm.Addr,
				}
				err := models.AddDepartment(department)
				if err != nil {
					logs.Error("add data error, %s", err)
					errors["submit"] = "update data error"
				} else {
					c.Redirect(c.URLFor("DepartmentController.Index"), http.StatusFound)
				}
			}

		} else {
			logs.Warn("parse form DepartmentForm error, %s", err)
			errors["submit"] = "server error"
		}
		c.Redirect(c.URLFor("DepartmentController.Index"), http.StatusFound)
	}
	c.Data["Errors"] = errors
	c.Data["Department"] = departmentForm
	c.TplName = "department/add.html"
}
