package main

import (
	"errors"
	"fmt"
)

func quickSort(sortSlice []int, isDesc bool) {
	var (
		tmp   = 0
		start = tmp + 1
		end   = len(sortSlice) - 1
	)
	if start > end {
		return
	}
	for start <= end {
		for end >= start {
			if (sortSlice[end] - sortSlice[tmp]) > 0 == isDesc {
				sortSlice[tmp], sortSlice[end] = sortSlice[end], sortSlice[tmp]
				tmp = end
				end--
				break
			}
			end--
		}
		for start <= end {
			if (sortSlice[start] - sortSlice[tmp]) < 0 == isDesc {
				sortSlice[tmp], sortSlice[start] = sortSlice[start], sortSlice[tmp]
				tmp = start
				start++
				break
			}
			start++
		}
	}
	quickSort(sortSlice[:tmp], isDesc)
	quickSort(sortSlice[tmp+1:], isDesc)
}

func binarySearch(sortSlice []int, num int) bool {
	start := 0
	end := len(sortSlice) - 1
	for start <= end {
		switch index := (start + end) / 2; {
		case sortSlice[index] < num:
			start = index + 1
		case sortSlice[index] > num:
			end = index - 1
		default:
			return true
		}
	}
	return false
}

func secondMax(sortSlice []int) (int, error) {
	if len(sortSlice) >= 2 {
		return sortSlice[len(sortSlice)-2], nil
	} else if len(sortSlice) == 1 {
		return -1, errors.New("只存在一个数值,无法判断第二大数值.")
	} else {
		return -1, errors.New("数据为空,无法判断第二大数值.")
	}
}

func diffSecondMax(sortSlice []int) (int, error) {
	tmpSlice := []int{}
	for _, num := range sortSlice {
		if binarySearch(tmpSlice, num) == false {
			tmpSlice = append(tmpSlice, num)
		}
	}
	if len(tmpSlice) >= 2 {
		return tmpSlice[len(tmpSlice)-2], nil
	} else if len(tmpSlice) == 1 {
		return -1, errors.New("去重后只存在一个数值,无法判断第二大数值.")
	} else {
		return -1, errors.New("数据为空,无法判断第二大数值.")
	}

}

func main() {
	numSlice := []int{23, 43, 27, 8, 31, 82, 2, 66, 58, 73, 14, 82, 82, 37}

	fmt.Println(numSlice)
	quickSort(numSlice, false)

	if num, err := secondMax(numSlice); err == nil {
		fmt.Println("不去重，第二大的数值:", num)
	} else {
		fmt.Println(err)
	}

	if num, err := diffSecondMax(numSlice); err == nil {
		fmt.Println("去重后，第二大的数值:", num)
	} else {
		fmt.Println(err)
	}
}
