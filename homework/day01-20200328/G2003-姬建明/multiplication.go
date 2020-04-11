package main

import "fmt"

func main() {
	// 打印乘法口诀
	for m := 1; m <= 9; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%d * %d = %2d\t", n, m, n*m)
		}
		fmt.Println()
	}
}
