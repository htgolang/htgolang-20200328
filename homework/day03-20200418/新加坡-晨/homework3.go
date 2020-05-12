package main

import (
	"fmt"
	"sort"
	"strconv"
)

var records = [][4]string{
	{"1.1.1.1", "/index.html", "200", "1000"},
	{"1.1.1.4", "/index.html", "200", "100"},
	{"1.1.1.1", "/index.html", "200", "10000"},
	{"1.1.1.2", "/index.html", "300", "4000"},
	{"1.1.1.4", "/index.html", "200", "100"},
	{"1.1.1.6", "/index.html", "304", "20000"},
	{"1.1.1.2", "/index.html", "200", "3000"},
	{"1.1.1.4", "/index.html", "200", "1000"},
	{"1.1.1.8", "/index.html", "200", "31000"},
}

func main() {
	//每个ip出现的次数
	ip := map[string]int{}
	for _,record :=range records{
		ip[record[0]]++
	}
	//fmt.Println(ip)
	//每个状态码出现的次数
	status_code := map[string]int{}
	for _,record := range records{
		status_code[record[2]]++
	}
	//fmt.Println(status_code)

	//每个IP在每个url上产生的流量
	ip_url := map[string]int{}
	for _,record := range records{
		new_url := record[0]+record[1]
		v,_ :=strconv.Atoi(record[3])
		ip_url[new_url] += v
	}
	//fmt.Println(ip_url)

	//对以上三个需求找到top10
	//ip
	if len(sort_map(ip))<10{
		fmt.Println(sort_map(ip))
	}else{
		fmt.Println(sort_map(ip)[:10])
	}

	//状态码
	if len(sort_map(status_code))<10{
		fmt.Println(sort_map(status_code))
	}else{
		fmt.Println(sort_map(status_code)[:10])
	}
	//ip_url流量
	if len(sort_map(ip_url))<10{
		fmt.Println(sort_map(ip_url))
	}else{
		fmt.Println(sort_map(ip_url)[:10])
	}

}

//排序map
type kv struct {
	Key string
	Value int
}
func sort_map(m map[string]int) []kv  {
	var slice []kv
	for k,v := range m {
		slice = append(slice,kv{k,v})
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Value > slice[j].Value
	})
	return slice
}