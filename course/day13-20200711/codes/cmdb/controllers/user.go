package controllers

import (
	"fmt"

	"cmdb/base/controllers/auth"
	"cmdb/models"
)

// UserController 用户管理控制器
type UserController struct {
	auth.AuthorizationController
}

// Query 查询用户
func (c *UserController) Query() {
	fmt.Println("Query")
	q := c.GetString("q")

	users := models.QueryUser(q)
	c.Data["users"] = users
	c.Data["q"] = q
	c.TplName = "user/query.html"
}
