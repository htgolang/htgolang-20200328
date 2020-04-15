package main

import "fmt"

func main() {
	arr := []int{5, 4, 6, 3, 7, 2, 1, 9, 8,100,59}
	for i := 1; i < len(arr); i++ {
		e := arr[i]
		fmt.Println("第一层循环I:", i, "e:", e)
		j := 0
		for j = i; j > 0 && arr[j-1] > e; j-- {

			arr[j] = arr[j-1]
			fmt.Println("第二层循环j:", j, "arr:", arr)

		}
		arr[j] = e
	}
	fmt.Println("最终结果:", arr)
}
