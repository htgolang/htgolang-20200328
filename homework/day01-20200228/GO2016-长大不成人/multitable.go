package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d  ", i, j, i*j)
		}
		fmt.Printf("\n")
	}
}
