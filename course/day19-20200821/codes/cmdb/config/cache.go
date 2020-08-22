package config

import (
	"github.com/astaxie/beego/cache"
)

var Cache cache.Cache

func Init(adapter, config string) {
	var err error
	Cache, err = cache.NewCache(adapter, config)
	if err != nil {
		panic(err)
	}
}
