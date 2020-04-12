package main

import "fmt"

func binary_search(alist []int, item int) bool {
	start := 0
	end := len(alist) - 1
	for start <= end {
		mindpoint := (start + end) / 2
		if alist[mindpoint] == item {
			return true
		} else if item < alist[mindpoint] {
			end = mindpoint - 1
		} else {
			start = mindpoint + 1
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(binary_search(arr, 4))
}
