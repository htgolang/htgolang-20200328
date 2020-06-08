package main

import "fmt"

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		insertIndex := i - 1
		insertVal := arr[i]

		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}

	return arr
}

func main() {
	arr := []int{7, 2, 5, 8, 22}
	fmt.Println(insertSort(arr))
}
