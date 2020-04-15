package main

import "fmt"

func binarySea(inSlice []int,destInt int) int  {
	sliceLen := len(inSlice)
	for i:=0;i < sliceLen-1 ;i++  {
		if inSlice[i] > inSlice[i+1]{
			fmt.Println("切片不是从小到大排序的")
			return -1
		}
	}
	//low和high用于判断跳出循环使用
	low := 0
	high := sliceLen
	for i := sliceLen/2;;{
		if inSlice[i] > destInt{
			high = i
			i = (i+low)/2
		}else if inSlice[i] < destInt{
			low = i
			i = (i+high)/2
		}else {
			return i
		}
		if low == high{
			return -1
		}
	}
}

func main() {
	srcSlice  := []int{1,2,6,9,12,16,20}
	fmt.Println(binarySea(srcSlice,5))
}