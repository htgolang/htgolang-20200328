package main

import "fmt"

func secondNumV1(inSlice []int) int {
	// 并列取相同次序
	for i := 1 ; i < len(inSlice) ; i++{
		tmp := inSlice[i]
		for j := 0;j < i ;j++  {
			if tmp > inSlice[j]{
				copy(inSlice[j+1:i+1],inSlice[j:i+1])
				inSlice[j]=tmp
				break
			}else if tmp == inSlice[j]{
				copy(inSlice[i:],inSlice[i+1:])
				inSlice=inSlice[:len(inSlice)-1]
				i--
				break
			}
		}
	}
	return inSlice[1]
}
func secondNumV2(inSlice []int) int {
	// 并列分前后顺序
	sliceLen := len(inSlice)
	for i := 1 ; i < sliceLen ; i++{
		tmp := inSlice[i]
		for j := 0;j < i ;j++  {
			if tmp > inSlice[j]{
				copy(inSlice[j+1:i+1],inSlice[j:i+1])
				inSlice[j]=tmp
				break
			}
		}
	}
	return inSlice[1]
}

func main() {
	src_slice := []int{1,2,4,5,5,5,5,6,7,7}
	// 因为函数传递切片是地址传递，会改变原切片，而v1函数修改了切片的值，故先执行v2
	fmt.Println(secondNumV2(src_slice))
	fmt.Println(secondNumV1(src_slice))

}
