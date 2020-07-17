package module

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/strive-after/go-cmdb/util"
)

const  (
	userTableName string = "user"
	managerTableName string = "manager"
	groupTableName string = "group"
)

var (
	ok bool
	err       error
)

func init() {
	//未注册  必须注册才能用 有时候session获取的时候  它会提示  gob: name not registered for interface: "github.com/strive-after/go-kubernetes/module.User"
	//那么就注册一下
	//因为session是gob方式加载结构体等就需要注册
	gob.Register(User{})
	//beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("redisurl")
	beego.AddFuncMap("TimeForMat",util.StopTimeFormat)
	beego.AddFuncMap("Roles",util.Role)
	beego.AddFuncMap("Left",util.IndexLeft)
	beego.AddFuncMap("Right",util.IndexRight)
	err := logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)  //separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	if err != nil {
		beego.Error(err)
		return
	}
}


type Operation interface {
	CRUD
	//返回一个表名  如果做关联查询好用
	TableName() string
	//UserShare
}




type CRUD interface {
	Add(mold interface{}) error
	GetId(mold interface{})  error
	Update(mold interface{}) error
	Del(id uint) error
	Get(mold string,value interface{}) error
	GetAll(mold interface{}) (error)
	UpdateMold(value interface{}) error
	Query(method string,mold interface{}) error
}




func NewOperation(mold interface{}) Operation{
	var operation  Operation
	switch mold.(type) {
	case *User:
		operation = mold.(*User)
		return operation
	case *Manager:
		operation = mold.(*Manager)
		return  operation
	}

	return nil
}