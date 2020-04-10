/* 1. 打印乘法口诀
   1 * 1 = 1
   2 * 1 = 2   2 * 2 = 4
   3 * 1 = 3   3 * 2 = 6   3 * 3 = 9
   ...
   9 * 1 = 9 ...                          ...  9 * 9 = 81
   外循环数1-9 内循环数小于外循环数 计算 制表符 回车
   */

package main

import "fmt"

 func main()  {
	for i := 1; i < 10; i++ {
		for a := 1; a <= i ; a++ {
			fmt.Printf("%d * %d = %2d\t" ,i ,a ,i*a)
		}		
		fmt.Println()
	}
}