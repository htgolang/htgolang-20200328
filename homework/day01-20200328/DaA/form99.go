package main

import "fmt"

func main() {
	/*
		9*9乘法表
		1*1=1
		2*1=1 2*2=4
		3*1=3 3*2=6 3*3=9

		1.循环主数
		2.循环辅数，辅助小于等于主数
	*/

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-2d ", i, j, i*j)
		}
		fmt.Println()
	}

}
