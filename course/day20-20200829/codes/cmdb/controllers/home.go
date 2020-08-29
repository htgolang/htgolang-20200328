package controllers

import (
	"cmdb/base/controllers/auth"
	"fmt"
	"time"

	"cmdb/config"
)

// HomeController 首页控制器
type HomeController struct {
	auth.LayoutController
}

// Index 首页显示方法
func (c *HomeController) Index() {
	v := config.Cache.Get("stime")
	if v != nil {
		vv, ok := v.([]byte)
		fmt.Println(string(vv), ok)
	}
	config.Cache.Put("stime", time.Now().Format("2006-01-02 15:04:05"), time.Minute)
	c.TplName = "home/index.html"
}
