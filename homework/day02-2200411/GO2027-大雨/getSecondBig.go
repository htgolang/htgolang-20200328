package main

import "fmt"

func sortArr(arr []int) (outarr []int) {
	//arr := []int{5, 4, 6, 3, 7, 2, 1, 9, 8, 100, 59}
	for i := 1; i < len(arr); i++ {
		e := arr[i]

		j := 0
		for j = i; j > 0 && arr[j-1] > e; j-- {

			arr[j] = arr[j-1]

		}
		arr[j] = e
	}
	//fmt.Println("最终结果:", arr)
	return arr
}

func main() {
	arr := []int{1, 10, 5, 10, 6, 7, 10}
	arr = sortArr(arr)
	fmt.Println(arr)
	SecondBig := arr[len(arr)-1]
	SecondBigAbreast := arr[len(arr)-2]
	for i := len(arr) - 2; i > 0; i-- {
		if arr[i] < SecondBig {
			SecondBig = arr[i]
			break
		}

	}
	fmt.Printf("去重后的第二大值:%d\n", SecondBig)
	fmt.Printf("不去重第二大值:%d\n", SecondBigAbreast)
}
