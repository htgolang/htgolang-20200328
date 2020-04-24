package main

import (
	"fmt"
	"strconv"
)

/*
{ // ip url 状态码 字节大小(B)
{"1.1.1.1", "/index.html", "200", "1000"},
{"1.1.1.2", "/index.html", "200", "10000"},
{"1.1.1.1", "/index.html", "200", "10000"}, }

 */

func IpStat(log map[int][]string,num int) map[string]int {
	IpNum := map[string]int{}
	for _,value :=  range log {
		IpNum[value[num]]++

	}
	return IpNum
}


func Traffic(log map[int][]string) map[string]int{
	num := map[string]int{}
	for _,v := range log{
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
	//日志以切片的方式插入
	logs[0] = append(logs[0],"1.1.1.1","/index.html","200","1000")
	logs[1] = append(logs[1],"1.1.1.2","/index.html","200","10000")
	logs[2] = append(logs[2],"1.1.1.1", "/index.html", "200", "10000")
	fmt.Println(logs)
	//IP访问次数统计
	num := IpStat(logs,0)
	fmt.Println(num)
	//状态码统计
	num = IpStat(logs,2)
	fmt.Println(num)

	//ip访问URL的总流量统计
	fmt.Println(Traffic(logs))
}
