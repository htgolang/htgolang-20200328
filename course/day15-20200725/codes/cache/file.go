package main

import (
	"encoding/gob"
	"fmt"
	"time"

	"github.com/astaxie/beego/cache"
)

type User struct {
	ID   int
	Name string
}

func main() {
	gob.Register(new(User))

	cache, _ := cache.NewCache("file", `{"CachePath" : "cache", "FileSuffix":".cache", "EmbedExpiry" : "60", "DirectoryLevel" : "3"}`)

	if cache.IsExist("name") {
		fmt.Println(cache.Get("name"))
	} else {
		fmt.Println("set")
		cache.Put("name", time.Now().Format("15:04:05"), time.Minute)
	}

	if cache.IsExist("user") {
		fmt.Println(cache.Get("user"))
	} else {
		fmt.Println("set")
		user := &User{1, "kk"}
		cache.Put("user", user, time.Minute)
	}

}
