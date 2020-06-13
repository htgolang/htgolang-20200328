package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg, _ := regexp.Compile("[ab0-9]+")
	fmt.Println(reg.FindAllString("0-a23-b3456", -1))

	// 定义非贪婪模式
	reg, _ = regexp.Compile("(?U)[ab0-9]+")
	fmt.Println(reg.FindAllString("0-a23-b3456", -1))

	// 将非贪婪模式转换为贪婪模式
	reg.Longest()
	fmt.Println(reg.FindAllString("0-a23-b3456", -1))
}
