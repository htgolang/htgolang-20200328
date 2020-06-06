package main

import (
	"fmt"
	"sort"
)

func main() {
	numSlice := []int{1, 2, 4, 5, 5}
	sort.Ints(numSlice)
	var i int
	selection := `
*********************************
你有两种方式获取到第二个最大值：
1：不论是否有相同的数值，取倒数第二个数。
2: 只取切片中第二大的数值。
3:
*********************************
`
	for {
		fmt.Print(selection, "\nplease input your selection: ")
		fmt.Scan(&i)

		switch i {
		case 1:
			fmt.Println(numSlice[len(numSlice)-2])
		case 2:
			for i := 1; i < len(numSlice); i++ {
				if numSlice[len(numSlice)-i-1] < numSlice[len(numSlice)-i] {
					fmt.Println(numSlice[len(numSlice)-i-1])
					break
				}
			}
		default:
			fmt.Println("选择错误，请重新选择")
		}
	}

}
