package main

import "fmt"

func main() {
	numbers := []int{9, 3, 10, 2, 5, 4}
	InsertSort(numbers)

	fmt.Println(numbers)
}

func InsertSort(values []int) {
	length := len(values)
	if length <= 1 {
		return
	}

	for i := 1; i < length; i++ {
		tmp := values[i] // 从第二个数开始，从左向右依次取数
		key := i - 1     // 下标从0开始，依次从左向右

		// 每次取到的数都跟左侧的数做比较，如果左侧的数比取的数大，就将左侧的数右移一位，直至左侧没有数字比取的数大为止
		for key >= 0 && tmp < values[key] {
			values[key+1] = values[key]
			key--
			//fmt.Println(values)
		}

		// 将取到的数插入到不小于左侧数的位置
		if key+1 != i {
			values[key+1] = tmp
		}
		//fmt.Println(values)
	}
}
