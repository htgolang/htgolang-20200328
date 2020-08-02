package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/cache"
)

func main() {
	cache, _ := cache.NewCache("memory", `{"interval" : 60}`)
	// 获取值
	// 放入cache
	// 删除Cache
	// 判断Cache是否存在

	fmt.Println(cache.Get("name"))

	cache.Put("name", "kk", 1000*time.Second)
	fmt.Println(cache.Get("name"))
	time.Sleep(12 * time.Second)
	fmt.Println(cache.Get("name"))

	cache.Put("name", "kk", 10*time.Second)
	fmt.Println(cache.Get("name"))
	fmt.Println(cache.IsExist("name"))

	cache.Delete("name")
	fmt.Println(cache.Get("name"))
	fmt.Println(cache.IsExist("name"))
}
