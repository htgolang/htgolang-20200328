package main

import "fmt"

func bubbling(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func main() {
	arr := []int{6, 4, 8, 5, 10}
	fmt.Println(bubbling(arr))
}
