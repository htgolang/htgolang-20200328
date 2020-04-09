package main

import "fmt"

func main() {
	for aVal := 1; aVal <= 9; aVal++ {
		for bVal := 1; bVal <= aVal; bVal++ {
			fmt.Printf("%dx%d=%02d\t", bVal, aVal, (bVal * aVal))
		}
		fmt.Println("")
	}
	for aVal := 9; aVal >= 1; aVal-- {
		for bVal := 1; bVal <= aVal; bVal++ {
			fmt.Print(bVal, "x", aVal, "=")
			fmt.Printf("%02d  ", aVal*bVal)
		}
		fmt.Println("")
	}
}
