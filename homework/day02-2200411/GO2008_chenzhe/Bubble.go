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

func bubbleSort(arr []int) []int  {
	arrLen := len(arr)
	for i := 0;i < arrLen ;i++  {
		for j := i;j < arrLen ;j++  {
			if arr[i] < arr[j]{
				arr[i],arr[j]=arr[j],arr[i]
			}
		}
	}
	return arr
}

func main()  {
	var numListBefore []int = ranNumber()
	fmt.Println(numListBefore)
	var numListAfter []int = bubbleSort(numListBefore)
	fmt.Println(numListAfter)
}
