package main

import "fmt"

//二分查找实现
func binarySearch(sliceIn []int, v int) (index int, times int) {
	begin, end := 0, len(sliceIn)-1
	times = 0
	for begin <= end {
		times++
		//其实值求和右移位算出中间值
		bnum := (begin + end) >> 1
		fmt.Printf("第%d次,index:%d\n", times, bnum)
		if sliceIn[bnum] < v {
			begin = bnum + 1
		} else if sliceIn[bnum] > v {
			end = bnum - 1
		} else {
			return bnum, times
		}
	}
	return -1, times
}

//切片排序实现
func sortNum(sliceIn []int) (sliceout []int) {
	for j := 0; j < len(sliceIn)-1; j++ {
		for i := 0; i < len(sliceIn)-1; i++ {
			if sliceIn[i] > sliceIn[i+1] {
				sliceIn[i], sliceIn[i+1] = sliceIn[i+1], sliceIn[i]
			}
		}
	}
	return sliceIn
}

//验证方法
func main() {
	//切片
	e2 := sortNum([]int{1, 13, 14, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	i := 15 //搜索值
	fmt.Printf("排序后:%v\n", e2)
	index, times := binarySearch(e2, i)
	if index == -1 {
		fmt.Printf("经过%d次二分查找,数字不存在", times)
	} else {
		fmt.Printf("经过%d次二分查找,数字存在,排序后的index:%d\n", times, index)
	}

}
