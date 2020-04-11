package main

import "fmt"

func main() {
	// 再控制台打印1..10
	for index := 1; index <= 10; index++ {
		fmt.Println(index)
	}

	// 计算 1+2+3 ... 100
	total := 0

	for index := 1; index <= 100; index++ {
		total += index
	}
	fmt.Println(total)

}
