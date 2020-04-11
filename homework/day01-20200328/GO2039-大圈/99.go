package main

import (
	"fmt"
)

//打印99乘法口诀
func printMultiTable()  {
	for i:=1;i<=9;i++ {
		for k:=1;k<=i;k++ {
			fmt.Printf("%d * %d = %d\t", k,i,i*k)
		}
		fmt.Println()
	}
}

func main() {
	//调用函数
	printMultiTable()
}
