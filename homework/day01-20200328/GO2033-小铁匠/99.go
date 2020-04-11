package main

import (
	"fmt"
)

func main() {
	//正序
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d\t", j, i, (i * j)) //%2d 占用2个位置，不够用空格占位，%03d:用0占位
		}
		fmt.Println()
	}

	fmt.Println()

	//逆序
	for i := 9; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d\t", j, i, (i * j))
		}
		fmt.Println()
	}

}
