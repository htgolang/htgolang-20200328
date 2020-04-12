package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	finddata := 3
	data := binarySearch(values, finddata)
	if data == -1 {
		fmt.Println("Can't find that data")
	} else {
		fmt.Println("Find that data", data)
	}
}

func binarySearch(arrayList []int, searchData int) int {
	var low int = 0
	// 获取最大下标
	var maxKey int = len(arrayList) - 1
	for low <= maxKey {
		// 获取中间的key
		var mid int = low + (maxKey-low)/2
		// 获取中间key对应的值
		var midVaule int = arrayList[mid]
		if midVaule == searchData {
			return midVaule
		} else if midVaule > searchData {
			maxKey = mid - 1
		} else if midVaule < searchData {
			low = mid + 1
		}
	}
	// 如果没找到，返回-1
	return -1
}
