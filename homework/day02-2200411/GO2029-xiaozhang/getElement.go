package main

import "fmt"

/*
切片， 获取切片中第二个最大元素（相同大小的元素） 1 2 4 5 5 并排第一为4，不并排为5

1 2 4 5 5
要计算最大的第二个数字，有两种
1. 5 5 相同，算并列第一（去重），所以4就是第二个最大的数字
2. 5 5 相同，但不考虑并列情况（不去重）， 5就是第二个最大的数字
这两种功能都实现
*/

var slice = []int{2, 1, 5, 4, 5}
var newSlice []int
var myMap = make(map[int]int)

// 冒泡算法排序，同时返回不去重时，第二大的数字
func bubbleAlgorithmV1(slice []int) int {
	for j := 0; j < len(slice); j++ {
		for k := 0; k < len(slice)-1-j; k++ {
			if slice[k] > slice[k+1] {
				tmp := slice[k]
				slice[k] = slice[k+1]
				slice[k+1] = tmp
			}
		}
	}
	return slice[len(slice)-2]
}

// 获取去重后的第二大的值。先去重再排序，再取倒数第二个值
func secMaxNum(slice []int) int {
	// 去重
	for v, k := range slice {
		myMap[k] = v
	}
	for k := range myMap {
		newSlice = append(newSlice, k)
	}
	// 排序并获取第二大的值
	return bubbleAlgorithmV1(newSlice)
}

func main() {
	// 当最大值不算并列（存在并列）时，那么第二个最大值依然是最大值
	fmt.Printf("算并列之后最大值是：%d\n", bubbleAlgorithmV1(slice))
	// 当最大值算并列（存在并列）是，那么第二个最大值就是仅小于最大值的那个值
	fmt.Printf("不算并列后的最大值是：%d\n", secMaxNum(slice))
}
