package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg, err := regexp.Compile("^132\\d{8}$")
	fmt.Println(err, reg)
	// 匹配 Match
	fmt.Println(reg.MatchString("132xxx"))
	fmt.Println(reg.MatchString("13212312312"))
	// 替换 Replace 132????????
	reg, err = regexp.Compile("132\\d{8}")
	fmt.Println(reg.ReplaceAllString("我的电话是132xxx请记录下", "132????????"))
	fmt.Println(reg.ReplaceAllString("我的电话是13212312312请记录下", "132????????"))
	// 查找 Find
	fmt.Println(reg.FindAllString("我的电话是13212312312,13212312313,15812312313", -1))
	// 分割
	reg, err = regexp.Compile("[:;\\t,]")
	fmt.Println(reg.Split("13212312312,13212312313,15812312313:xxx;zzzz\tyyyyy", -1))
}
