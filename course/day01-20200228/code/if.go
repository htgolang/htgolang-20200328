package main

import "fmt"

func main() {
	// 老婆
	// 买10个包子
	// 如果有卖西瓜的，买一个西瓜

	fmt.Println("买十个包子")

	var y string
	fmt.Print("有没有卖西瓜的:")
	fmt.Scan(&y)
	fmt.Println("你输入的是:", y)

	if y == "yes" {
		fmt.Println("买一个西瓜")
	}
}
