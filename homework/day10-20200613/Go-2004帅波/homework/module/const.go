package module

import (
	"github.com/astaxie/beego/orm"
	"sync"
)

const  (
	userTableName string = "user"
	groupTableName string = "group"
)

var (
	globalOrm orm.Ormer
	once      sync.Once
	err       error
)

func Ormer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
