package main

import "fmt"

func main() {
	// 老公

	// 如果有卖西瓜的 买一个包子， 否则买是个包子

	var y string
	fmt.Print("有卖西瓜的吗:")
	fmt.Scan(&y)
	// y = "yes"

	switch y {
	case "yes", "y", "Y":
		fmt.Println("买一个包子")
	case "no", "n", "N":
		fmt.Println("买十个包子")
	default:
		fmt.Println("输入错误")
	}
}
