// 插入排序
package main

import "fmt"

func insertSort(sortSlice []int, isDesc bool) {
	for i := 1; i < len(sortSlice); i++ {
		j := i
		sortValue := sortSlice[i]
		for ; j > 0; j-- {
			if (sortValue-sortSlice[j-1] > 0) == isDesc {
				sortSlice[j] = sortSlice[j-1]
			} else {
				break
			}
		}
		sortSlice[j] = sortValue
	}
}

func main() {
	numSlice := []int{23, 43, 27, 8, 31, 82, 2, 66, 58, 73, 14, 37}
	insertSort(numSlice, false)
	fmt.Println(numSlice)
}
