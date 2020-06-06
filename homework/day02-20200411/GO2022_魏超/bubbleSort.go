// 冒泡排序
package main

import "fmt"

// isDesc=true 等于为降序,否则为升序
func bubbleSort(sortSlice []int, isDesc bool) {
	for i := 0; i < len(sortSlice); i++ {
		for j := 0; j < len(sortSlice)-i-1; j++ {
			if (sortSlice[j]-sortSlice[j+1] < 0) == isDesc {
				sortSlice[j], sortSlice[j+1] = sortSlice[j+1], sortSlice[j]
			}
		}
	}
}

func main() {
	numSlice := []int{23, 43, 27, 8, 31, 82, 2, 66, 58, 73, 14, 37}
	bubbleSort(numSlice, false)
	fmt.Println(numSlice)
}
