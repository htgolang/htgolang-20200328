package main

import "fmt"

func main() {
	// 设置初始的排序列表
	values := []int{3, 11, 9, 28, 30, 99, 81, 56}
	// 打印初始的列表
	fmt.Println(values)
	// 进行从小到大的冒泡排序算法，并打印结果
	BubbleSmallsort(values)
	// 进行从大到小的冒泡排序算法，并打印结果
	BubbleLargesort(values)
}

func BubbleSmallsort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	fmt.Println(values)
}

func BubbleLargesort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	fmt.Println(values)
}
