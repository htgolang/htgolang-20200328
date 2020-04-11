package main

import "fmt"

func main() {
	var score float32

	fmt.Print("请输入分数:")
	fmt.Scan(&score)
	fmt.Println("你输入的分数是：", score)

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 60:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}
