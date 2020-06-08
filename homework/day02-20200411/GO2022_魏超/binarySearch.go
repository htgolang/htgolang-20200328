// 二分查找法
package main

import (
	"fmt"
)

func selectSort(sortSlice []int, isDesc bool) {
	for i := 0; i < len(sortSlice); i++ {
		min := i
		for j := i + 1; j < len(sortSlice); j++ {
			if (sortSlice[j] - sortSlice[min]) > 0 == isDesc {
				min = j
			}
		}
		sortSlice[i], sortSlice[min] = sortSlice[min], sortSlice[i]
	}
}

func binarySearch(sortSlice []int, num int) bool {
	start := 0
	end := len(sortSlice) - 1
	for start <= end {
		switch index := (start + end) / 2; {
		case sortSlice[index] < num:
			start = index + 1
		case sortSlice[index] > num:
			end = index - 1
		default:
			return true
		}
	}
	return false
}

func main() {
	var (
		num      int
		numSlice = []int{23, 43, 27, 8, 31, 82, 2, 66, 58, 73, 14, 37, 22}
	)

	fmt.Print("请输入你要查询的数字:")
	fmt.Scan(&num)
	selectSort(numSlice, false)
	fmt.Println(numSlice)
	if binarySearch(numSlice, num) {
		fmt.Println("您输入的数字存在.")
	} else {
		fmt.Println("您输入的数字不存在.")
	}
}
