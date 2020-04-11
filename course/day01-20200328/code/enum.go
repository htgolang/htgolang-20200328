package main

import "fmt"

func main() {
	const (
		Mon = iota // 在常量组内使用 iota 初始化 0, 每次调用+1
		Tuesd
		Wed
		Thur
		Fir
		Sat
		Sun
	)

	fmt.Println(Mon)
}
