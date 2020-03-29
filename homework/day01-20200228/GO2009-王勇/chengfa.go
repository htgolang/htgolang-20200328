package main

import "fmt"

func main() {
	//第一版
	for a := 1; a <= 9; a++ {
		for b := 1; b <= 9; b++ {
			fmt.Print(a, "*", b, "=", a*b)
		}
	}

	//fmt.Println(a, "*", b, "=", a*b)
	fmt.Println()
	//第二版
	for row := 1; row <= 9; row++ {
		for column := 1; column <= row; column++ {
			fmt.Printf("%3d * %d = %2d", column, row, column*row)
		}
		fmt.Println()
	}

}
