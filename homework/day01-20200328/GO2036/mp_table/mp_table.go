package main

import "fmt"

func main() {
	// 遍历，决定处理第几行
	for i := 1;i <= 9;i++ {

		//遍历，决定这一行有多少列
		for j := 1; j<=i;j++ {
			fmt.Printf("%d * %d = %d\t",j,i,i*j)
		}
		//手动执行回车，在i = j时回车，实际作用就是换行
		fmt.Println()
	}
}
