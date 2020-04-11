package main

import "fmt"

func main() {
	// 老公

	// 如果有卖西瓜的 买一个包子， 否则买是个包子

	var y string
	fmt.Print("有卖西瓜的吗:")
	fmt.Scan(&y)
	// y = "yes"

	if y == "yes" {
		fmt.Println("买一个包子")
	} else {
		fmt.Println("买是十个包子")
	}
}
