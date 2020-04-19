package main

import (
	"fmt"
	"sort"
)

func main() {
	num := map[rune]int{}
	dream := `adsadsagadasasddaefqrwqwetuppiypuertwq`
	//将出现的放入映射中
	for _,v := range dream {
		//当v在a-z  A-Z中表示他不是其他字符
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			num[v]++
		}
	}
	num02 := map[int][]rune{}
	//统一将出现次数放k   出现字母放切片中
	for k,v := range num{
		if _,ok := num02[v];ok {
			num02[v] = append(num02[v],k)
		}else {
			num02[v] = []rune{k}
		}
	}
	list01 := []int{}
	list02 := []rune{}
	//将出现次数放入切片中排序
	for k,_ := range num02{
		list01 = append(list01,k)
	}
	//降序
	sort.Sort(sort.Reverse(sort.IntSlice(list01)))
	//当出现次数的切片长度大于10 就取前10  否则取到最后
	if len(list01) >= 10 {
		//知道出现次数之后那么就可以根据字母的出现次数 将字母取出  按照出现次数顺序放入切片中
		for _,v := range list01[:10]{
			list02 = append(list02,num02[v]...)
		}
	}else {
		for _,v := range list01{
			list02 = append(list02,num02[v]...)
		}
	}
	//从按照出现次数多少的字母切片中  我们打印前10个 如果大于10取至10  否则 就取所有
	if len(list02) >= 10 {
		for _,v := range list02[:10] {
			fmt.Printf("%q出现了%d次\n",v,num[v])
		}
	}else {
		for _,v := range list02 {
			fmt.Printf("%q出现了%d次\n",v,num[v])
		}
	}

}
