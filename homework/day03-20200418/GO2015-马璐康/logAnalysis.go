package main

import (
	"fmt"
	"strconv"
)

func countNums(logFile [][4]string, num int) map[string]int {
	state := map[string]int{}
	for _, logData := range logFile {
		state[logData[num]]++
	}
	return state
}

func statisticsTraffic(logFile [][4]string, nums ...int) map[string]int {
	state := map[string]int{}
	for _, logData := range logFile {
		//fmt.Println(logData)
		//fmt.Println(nums)
		var num string
		for _, index := range nums {
			//	fmt.Println(index)
			num += " " + logData[index]
		}
		if traffic, err := strconv.Atoi(logData[3]); err == nil {
			state[num] += traffic
		}
	}
	return state
}

func insertSort(logMap map[string]int) []string {
	state := []string{}
	for k, _ := range logMap {
		state = append(state, k)
	}

	for i := 1; i < len(state); i++ {
		val := state[i]
		j := i - 1
		for j >= 0 && logMap[state[j]] < logMap[val] {
			state[j+1] = state[j]
			j--
		}
		state[j+1] = val
	}
	return state
}

func sortTop10(sortSlice []string, num int) []string {
	if len(sortSlice) < num {
		return sortSlice
	} else {
		return sortSlice[:num]
	}
}

func main() {
	logState := map[string]int{"IP": 0, "URL": 1, "StatusCode": 2, "Info": 3}
	logFile := [][4]string{
		// ip   url  状态码  字节大小(B)
		{"1.1.1.1", "/index1.html", "200", "1000"},
		{"1.1.1.3", "/index2.html", "302", "10000"},
		{"1.1.1.3", "/index3.html", "404", "100"},
		{"1.1.1.4", "/index4.html", "504", "1000"},
		{"1.1.1.5", "/index5.html", "503", "10000"},
		{"1.1.1.1", "/index1.html", "200", "100"},
		{"1.1.1.3", "/index3.html", "302", "1000"},
		{"1.1.1.3", "/index3.html", "503", "10000"},
		{"1.1.1.4", "/index3.html", "504", "100"},
		{"1.1.1.5", "/index2.html", "200", "1000"},
		{"1.1.1.2", "/index2.html", "302", "10000"},
		{"1.1.1.1", "/index1.html", "200", "100"},
		{"1.1.1.2", "/index2.html", "302", "10000"},
		{"1.1.1.3", "/index3.html", "302", "1000"},
		{"1.1.1.2", "/index3.html", "503", "10000"},
		{"1.1.1.4", "/index3.html", "504", "100"},
		{"1.1.1.4", "/index3.html", "504", "100"},
		{"1.1.1.6", "/index3.html", "504", "100"},
		{"1.1.1.7", "/index2.html", "302", "10000"},
		{"1.1.1.8", "/index3.html", "302", "1000"},
		{"1.1.1.9", "/index3.html", "503", "10000"},
		{"1.1.1.10", "/index3.html", "504", "100"},
		{"1.1.1.11", "/index3.html", "504", "100"},
	}
	fmt.Println("1:每个IP出现的次数：")
	logIpCount := countNums(logFile, logState["IP"])
	//fmt.Println(logIpCount)
	for IP, nums := range logIpCount {
		fmt.Println(IP, ":", nums)
	}

	fmt.Println("2:每个状态码出现的次数：")
	logStatusCodeCount := countNums(logFile, logState["StatusCode"])
	for statusCode, nums := range logStatusCodeCount {
		fmt.Println(statusCode, ":", nums)
	}
	fmt.Println("3:每个IP在每个URL上产生的流量：")
	logStatisticsTrafficCount := statisticsTraffic(logFile, logState["IP"], logState["URL"])
	for k, trafficData := range logStatisticsTrafficCount {
		fmt.Println(k, ":", trafficData)
	}

	fmt.Println("4: ip出现的次数top10排序")
	logIPTop10 := sortTop10(insertSort(logIpCount), 10)
	for _, IP := range logIPTop10 {
		fmt.Println(IP, logIpCount[IP])
	}

	fmt.Println("5: 每个状态码出现的次数top10排序：")
	logStatusCodeTop10 := sortTop10(insertSort(logStatusCodeCount), 10)
	for _, statusCode := range logStatusCodeTop10 {
		fmt.Println(statusCode, logStatusCodeCount[statusCode])
	}
	fmt.Println("6: 每个IP在每个URL上产生的流量top10排序：")
	logStatisticsTrafficTop10 := sortTop10(insertSort(logStatisticsTrafficCount), 10)
	for _, traffic := range logStatisticsTrafficTop10 {
		fmt.Println(traffic, logStatisticsTrafficCount[traffic])
	}
}
