package main

import "fmt"

func main() {
	fmt.Println("打印99乘法表:")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%-2d   ", i, j, i*j)
		}
		fmt.Println()
	}
}
