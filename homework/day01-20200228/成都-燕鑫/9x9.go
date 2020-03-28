package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 9; i++ {
		for m := 1; m <= i; m++ {
			fmt.Printf("%d*%d=%-2d ", m, i, m*i)
		}
		fmt.Print("\n")
	}
}
