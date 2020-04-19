package main

import (
	"fmt"
	"strconv"
)

//{"1.1.1.1", "/index.html", "200", "1000"},
//{"1.1.1.2", "/index.html", "200", "10000"},
//{"1.1.1.1", "/index.html", "200", "10000"}

func IP(b map[int][]string,n int) map[string]int{
	num := map[string]int{}
	for _,v := range b{
		num[v[n]]++
	}
	return  num
}

func URL(b map[int][]string) map[string]int{
	num := map[string]int{}
	for _,v := range b{
		if _,ok := num[v[0]+v[1]];ok {
			sum ,_ := strconv.Atoi(v[3])
			num[v[0]+v[1]] += sum
		}else {
			num[v[0]+v[1]],_ = strconv.Atoi(v[3])
		}
	}
	return  num
}


func main() {
	logs := map[int][]string{}
	logs[0] = append(logs[0],"1.1.1.1","/index.html","200","1000")
	logs[1] = append(logs[1],"1.1.1.2","/index.html","200","10000")
	logs[2] = append(logs[2],"1.1.1.1", "/index.html", "200", "10000")
	nums := IP(logs,0)
	//ip出现次数
	fmt.Println(nums)
	//状态码出现次数
	nums = IP(logs,2)
	fmt.Println(nums)
	nums = URL(logs)
	fmt.Println(nums)
}