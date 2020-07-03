package main

import "fmt"

func main() {
	// 99乘法表
	for i:=1;i<=9; i++{
		for j := 1; j <= i; j++ {
			if i == 1{
				fmt.Printf("| %d ✖ ️%d = %d |\n|", j, i, j*i)
			}else if i == j {
				if j==9{
					fmt.Printf(" %d ✖ ️%d = %d |\n", j, i, j*i)
				} else{
					fmt.Printf(" %d ✖ ️%d = %d |\n|", j, i, j*i)
				}
			}else{
				fmt.Printf(" %d ✖ ️%d = %d |", j, i, j*i)
			}
		}
	}
}
