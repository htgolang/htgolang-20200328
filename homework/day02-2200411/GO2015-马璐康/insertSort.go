package main

import "fmt"

func main() {

	arr := []int{36, 75, 6, 9, 36, 54, 76, 909, 102, 28}

	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1

		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = key

	}
	fmt.Println("排序后：", arr)
}
