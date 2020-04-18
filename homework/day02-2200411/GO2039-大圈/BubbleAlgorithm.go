package main

import "fmt"

//冒泡算法的实现

/*
用户输入5个数字，使用冒泡算法对用户输入的数字进行排序
 */
var num int
var sel []int
func userInput() {
	for i:=0; i<5;i++ {
		fmt.Printf("请输入一个数字：")
		fmt.Scan(&num)
		sel = append(sel,num)
	}
	fmt.Printf("切片：%#v\n",sel)
	fmt.Println("=========开始排序===========")
}

func bubbleAlgorithm() {
	//根据元素个数确定外层循环的次数
	for j:=0;j<len(sel)-1;j++ {
		//内层循环的次数是变化的，根据外层第几次循环而改变
		for k:=0;k<len(sel)-j-1;k++ {
			if sel[k] > sel[k+1]  {
				tmp := sel[k]
				sel[k] = sel[k+1]
				sel[k+1] = tmp
			}
		}
		fmt.Printf("第%d次外层循环结束时的排序结果：%v\n", j+1, sel)
	}
	fmt.Println("=========排序完成后的最终结果===========")
	fmt.Println(sel)
}

func main() {
	userInput()
	bubbleAlgorithm()
}