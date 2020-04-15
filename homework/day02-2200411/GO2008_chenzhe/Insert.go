package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ranNumber() []int {
	rand.Seed(time.Now().Unix())
	var intArry []int = make([]int,20)
	for i,_:= range intArry{
		intArry[i]=rand.Intn(1000)
	}
	return intArry
}

func insertSort(arr []int) []int{
	arrLen := len(arr)
	var tmp int
	for i := 1;i < arrLen  ;i++  {
		tmp = arr[i]
		for j := 0;j < i ;j++{
			if tmp > arr[j]{
				copy(arr[j+1:i+1],arr[j:i+1])
				arr[j]=tmp
				break
			}
		}
	}
	return arr
}


func main() {
	var numListBefore []int = ranNumber()
	fmt.Println(numListBefore)
	var numListAfter []int = insertSort(numListBefore)
	fmt.Println(numListAfter)
}
