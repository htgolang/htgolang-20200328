package main

import "fmt"

func main() {
	// 设置初始的排序列表
	values := []int{1, 2, 3, 5, 5}
	var sliceData []int
	// 先做一下排序
	// =sliceData = BubbleSmallsort(values)
	sliceData = BubbleSmallsort(values)
	maxKey := len(sliceData) - 1
	maxData := sliceData[maxKey]
	// 算法1： 打印出第二大元素（排除最大值）
	for maxKey > 0 {
		tmp := sliceData[maxKey-1]
		if tmp == maxData {
			maxKey--
		} else {
			fmt.Println("第二大元素为", tmp)
			break
		}
	}

	maxKey = len(sliceData) - 1
	// 算法2： 打印出第二大元素（可能是最大值）
	fmt.Println("第二大元素为", sliceData[maxKey-1])

}

func BubbleSmallsort(values []int) []int {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return values
}
