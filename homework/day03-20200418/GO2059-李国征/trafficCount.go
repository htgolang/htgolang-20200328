package main

import (
	"fmt"
	"strconv"
)

// 统计出现的流量数量

// 定义数据
var data [][4]string

func Count(num int, data [][4]string) map[string]int {
	// IP数量统计
	// num 数量
	mapCount := map[string]int{}
	for _, i := range data {
		mapCount[i[num]]++
	}
	return mapCount
}

func trafficCount(data [][4]string) map[string]int {
	// 计算流量
	count := map[string]int{}
	for _, i := range data {
		_tmp := fmt.Sprintf("%s%s", i[0], i[1])
		_tmp_int, _ := strconv.Atoi(i[3])
		count[_tmp] = count[_tmp] + _tmp_int
	}
	return count
}

func main() {
	data = [][4]string{
		{"1.1.1.1", "/index.html", "200", "10000"},
		{"1.1.1.2", "/index.html", "200", "10293"},
		{"1.1.1.1", "/index.html", "200", "103450"},
		{"1.1.1.5", "/index.html", "200", "10340"},
		{"1.1.1.7", "/index.html", "200", "10000"},
		{"1.1.1.9", "/index.html", "200", "105300"},
		{"1.1.1.10", "/index.html", "200", "10000"},
		{"1.1.1.1", "/index.html", "200", "106210"},
		{"1.1.1.3", "/index.html", "200", "100350"},
		{"1.1.1.2", "/index.html", "200", "100460"},
		{"1.1.1.6", "/index.html", "200", "10450"},
		{"1.1.1.9", "/index.html", "200", "23630"},
		{"1.1.1.2", "/index.html", "200", "30000"},
		{"1.1.1.5", "/index.html", "200", "10340"},
		{"1.1.1.7", "/index.html", "200", "10000"},
		{"1.1.1.9", "/index.html", "200", "105300"},
		{"1.1.1.10", "/index.html", "200", "10000"},
		{"1.1.1.1", "/index.html", "200", "106210"},
		{"1.1.1.5", "/index.html", "200", "10340"},
		{"1.1.1.7", "/index.html", "200", "10000"},
		{"1.1.1.9", "/index.html", "200", "105300"},
		{"1.1.1.10", "/index.html", "200", "10000"},
		{"1.1.1.1", "/index.html", "200", "106210"},
		{"1.1.1.5", "/index.html", "200", "10340"},
		{"1.1.1.7", "/index.html", "200", "10000"},
		{"1.1.1.9", "/index.html", "200", "105300"},
		{"1.1.1.10", "/index.html", "200", "10000"},
		{"1.1.1.1", "/index.html", "200", "106210"},
		{"1.1.1.3", "/index.html", "200", "100350"},
		{"1.1.1.2", "/index.html", "200", "100460"},
		{"1.1.1.6", "/index.html", "200", "10450"},
		{"1.1.1.9", "/index.html", "200", "23630"},
		{"1.1.1.2", "/index.html", "200", "30000"},
	}
	// sort.Search(data, func(i, j int) bool { return data[i][0] > data[j][0] })
	// IP出现的次数
	mapCount := Count(0, data)
	fmt.Println(mapCount)
	// 状态码出现的次数
	statusCount := Count(2, data)
	fmt.Println(statusCount)
	// 每个IP在每个URL上产生的流量
	countInt := trafficCount(data)
	fmt.Println(countInt)
}
