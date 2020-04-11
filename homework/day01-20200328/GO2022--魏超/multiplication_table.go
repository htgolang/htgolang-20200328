
/*
九九乘法表
*/
package main

import (
	"fmt"
)

const COUNT int = 10

func main() {
	for i := 1; i < COUNT; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%2d * %2d = %2d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
