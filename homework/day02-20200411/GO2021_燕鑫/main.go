package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Homework 1:")
	testDubble()
	fmt.Println("Homework 3:")
	testInsertion()
	fmt.Println("Homework 2:")
	a:=[]int{1,2,4,5,5}
	findSecondLargestDenseRank(a)
	findSecondLargestRowNumber(a)
	fmt.Println("Homework 4:")
	testBinarySearch()

}

func testDubble() {
	fmt.Println("Dubble Sort:")
	a := []int{7, 6, 4, 3, 5, 2, 1, 7, 8, 9, 20, 10, 15, 13, 3}
	fmt.Println("Before sorting: ", a)
	a = bubble(a, false)
	fmt.Println("After sorting: ", a)
}

func testInsertion() {
	fmt.Println("Insertion Sort:")
	a := []int{7, 6, 4, 3, 5, 2, 1, 7, 8, 9, 20, 10, 15, 13, 3}
	fmt.Println("Before sorting: ", a)
	a = insertion(a, false)
	fmt.Println("After sorting: ", a)
}

func bubble(sli []int, desc bool) []int {
	switch desc {
	case true:
		for i := 0; i < len(sli)-1; i++ {
			for j := 0; j < len(sli)-i-1; j++ {
				if sli[j] < sli[j+1] {
					sli[j], sli[j+1] = sli[j+1], sli[j]
				}
			}
		}
	case false:
		for i := 0; i < len(sli)-1; i++ {
			for j := 0; j < len(sli)-i-1; j++ {
				if sli[j] > sli[j+1] {
					sli[j], sli[j+1] = sli[j+1], sli[j]
				}
			}
		}
	}

	return sli
}

func insertion(sli []int, desc bool) []int {
	switch desc {
	case true:
		for i := 1; i < len(sli); i++ {
			insertValue := sli[i]
			for j := i - 1; j >= 0; j-- {
				if insertValue > sli[j] {
					sli[j+1], sli[j] = sli[j], sli[j+1]
				} else {
					break
				}
			}
		}
	case false:
		for i := 1; i < len(sli); i++ {
			insertValue := sli[i]
			for j := i - 1; j >= 0; j-- {
				if insertValue < sli[j] {
					sli[j+1], sli[j] = sli[j], sli[j+1]
				} else {
					break
				}
			}
		}
	}

	return sli
}

//This function allows that two numbers have same rank
func findSecondLargestDenseRank(sli []int) {
	sli = bubble(sli, true)
	for i := 0; i < len(sli)-1; i++ {
		if sli[i] != sli[i+1] {
			fmt.Println("The second lagest number of specified array is: ", sli[i+1])
			return
		}
	}
}

//There isn't any number which has the same rank with others
func findSecondLargestRowNumber(sli []int) {
	sli = bubble(sli, true)
	fmt.Println("The second lagest number of specified array is: ", sli[1])
}

func testBinarySearch() {
	//The arrary must be monotone increasing.
	sli := []int{1, 2, 3, 5, 7, 10, 12}
	for {
		fmt.Print("Enter num: ")
		b := ""
		fmt.Scan(&b)
		if b=="q"{
			break
		}
		n, _ := strconv.Atoi(b)
		if index := binarySearch(sli, n); index == -1 {
			fmt.Println("The given number isn't in this array!")
		} else {
			fmt.Printf("The index of given number is %d!\n", index)
		}
	}
}


func binarySearch(sli []int, value int) (index int) {
	for m, n := 0, len(sli)-1; m <= n; {
		mid := (m+n)/2
		if value < sli[mid] {
			n--
		} else if value > sli[m] {
			m++
		} else {
			return m
		}
	}
	return -1
}
