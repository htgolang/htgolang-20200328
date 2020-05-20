package main

import (
	"fmt"
	"sort"
	"strconv"
)

//下标为0元素是IP  次数加一
func ip(logs [][4]string) map[string]int {
	timeip := map[string]int{}
	for _, v := range logs {
		timeip[v[0]]++
	}
	return timeip
}

//下标为2的元素是状态码  次数加一
func status(logs [][4]string) map[string]int {
	timestatus := map[string]int{}
	for _, v := range logs {
		timestatus[v[2]]++
	}
	return timestatus
}

//输入内容 打印函数
func print(inputt map[string]int, typee string) {
	var typeee string = ""
	if typee == "" {
		typee = typeee
	}
	fmt.Println(typee, "结果如下：")
	for k, v := range inputt {
		fmt.Println(k, ":", v)
	}
}

//排序函数  输入map 根据
func sortt(orderr map[string]int, op string, limitt int) {
	slicemap := make([]string, 0, len(orderr))
	for kk, _ := range orderr {
		slicemap = append(slicemap, kk)
	}
	//排序
	if op == ">" {
		fmt.Println("降序排列")
		sort.Slice(slicemap, func(m, n int) bool { return orderr[slicemap[m]] > orderr[slicemap[n]] })
	} else if op == "<" {
		fmt.Println("升序排列")
		sort.Slice(slicemap, func(m, n int) bool { return orderr[slicemap[m]] < orderr[slicemap[n]] })
	}
	// for _, vv := range slicemap {
	// 	fmt.Printf("%s : %d\n", vv, orderr[vv])
	// }

	//获取前几的切片 遍历输出
	laterslice := slicemap[:limitt]
	fmt.Println("输出前", limitt, "个")
	if limitt > len(laterslice) {
		fmt.Println("输入错误，请推出")
		return
	}
	for i := 0; i < len(laterslice); i++ {
		fmt.Println(laterslice[i])
	}
}
func ipurl(mergee [][4]string) map[string]int {
	ipursum := map[string]int{}
	for _, vv := range mergee {
		m, _ := strconv.Atoi(vv[3])
		ipursum[vv[0]+vv[1]] = ipursum[vv[0]+vv[1]] + m
	}
	// for kk, vv := range ipursum {
	// 	fmt.Println(kk, vv)
	// }
	return ipursum
}
func main() {
	//定义一个切片  4个元素 对应日志
	logs := [][4]string{
		{"1.1.1.6", "/index.html", "500", "1000"},
		{"1.1.1.6", "/index.html", "500", "1001"},
		{"1.1.1.3", "/index.html", "200", "1002"},
		{"1.1.1.3", "/index.html", "400", "1003"},
		{"1.1.1.5", "/index.html", "400", "1004"},
		{"1.1.1.6", "/index.html", "400", "1005"},
		{"1.1.1.4", "/index.html", "200", "1003"},
		{"1.1.1.5", "/index.html", "200", "1004"},
		{"1.1.1.6", "/index.html", "404", "1005"},
		{"1.1.1.1", "/index.html", "200", "1000"},
		{"1.1.1.2", "/index.html", "500", "1001"},
		{"1.1.1.3", "/index.html", "200", "1002"},
	}

	timeip := ip(logs)
	print(timeip, "IP")

	timestatus := status(logs)
	print(timestatus, "status")

	sortt(timeip, ">", 5)
	sortt(timestatus, ">", 2)

	resultt := ipurl(logs)
	print(resultt, "")

}
