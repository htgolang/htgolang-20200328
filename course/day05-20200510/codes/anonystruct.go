package main

import "fmt"

func main() {
	// 匿名结构体 ==> 直接初始化给一个变量
	user := struct {
		id   int
		name string
		age  int
	}{1, "kk", 30}

	fmt.Printf("%T\n", user)
	fmt.Printf("%#v\n", user)
	fmt.Println(user.name)
	user.name = "kk"
	fmt.Println(user.name)

	user = struct {
		id   int
		name string
		age  int
	}{1, "kk", 30}

	fmt.Printf("%#v\n", user)

}
