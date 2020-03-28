package main

import "fmt"

func main() {
	name := ""
	fmt.Print("请输入你的名字:")

	fmt.Scan(&name)

	fmt.Println("你输入的名字是：", name)

	age := 0
	fmt.Print("请输入你的年龄:")

	fmt.Scan(&age)

	fmt.Println("你输入的年龄是：", age)

	msg := ""
	fmt.Print("请输入你的msg:")

	fmt.Scan(&msg)

	fmt.Println("你输入的msg是：", msg)

}
