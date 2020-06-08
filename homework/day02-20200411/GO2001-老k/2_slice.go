package main

import "fmt"

//切片，获取切片中第二个最大元素(相同大小的元素) 1 2 4 5 5 并排第一 4 不并排 5
func main() {
	arr := []int{1, 2, 4, 5, 5}
	sliceData := BubbleSort(arr)
	maxKey := len(sliceData) - 1
	maxData := sliceData[maxKey]
	fmt.Println(maxKey)
	fmt.Println(maxData)
	for {
		result := sliceData[maxKey]
		if result == maxData {
			maxKey--
		} else {
			fmt.Println("第二大的值：", result)
			break
		}
	}

}
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
