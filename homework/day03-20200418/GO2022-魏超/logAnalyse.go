package main

import (
	"fmt"
	"strconv"
)

func statisticsCount(logs [][4]string, index int) map[string]int {
	var countInfo = map[string]int{}
	for _, logInfo := range logs {
		countInfo[logInfo[index]]++
	}
	return countInfo
}

func statisticsflow(logs [][4]string, indexs ...int) map[string]int {
	var keyCount = map[string]int{}
	for _, logInfo := range logs {
		var key string
		for _, index := range indexs {
			key += " " + logInfo[index]
		}
		if flow, err := strconv.Atoi(logInfo[3]); err == nil {
			keyCount[key] += flow
		}
	}
	return keyCount
}

func insertSort(keyMap map[string]int, isDesc bool) []string {
	keySlice := []string{}
	for key, _ := range keyMap {
		keySlice = append(keySlice, key)
	}
	for i := 1; i < len(keySlice); i++ {
		j := i
		keyValue := keySlice[i]
		for ; j > 0; j-- {
			if (keyMap[keyValue]-keyMap[keySlice[j-1]] > 0) == isDesc {
				keySlice[j] = keySlice[j-1]
			} else {
				break
			}
		}
		keySlice[j] = keyValue
	}
	return keySlice
}

func choiceTop(sortSlice []string, num int) []string {
	if len(sortSlice) >= num {
		return sortSlice[:num]
	} else {
		return sortSlice
	}
}

func main() {
	var (
		infoMap = map[string]int{"ip": 0, "url": 1, "status": 2, "data": 3}
		logs    = [][4]string{
			{"1.1.1.1", "/index.html", "200", "1000"},
			{"1.1.1.2", "/index.html", "200", "10000"},
			{"1.1.1.1", "/index.html", "200", "10000"},
			{"1.1.1.1", "/index1.html", "500", "10000"},
			{"1.1.1.3", "/index1.html", "500", "1000"},
			{"1.1.1.1", "/index1.html", "200", "1000"},
			{"1.1.1.1", "/index1.html", "200", "1000"},
			{"1.1.1.12", "/index1.html", "400", "10000"},
			{"1.1.1.2", "/index.html", "505", "1000"},
			{"1.1.1.2", "/index.html", "200", "50000"},
			{"1.1.1.10", "/index.html", "400", "10000"},
			{"1.1.1.8", "/index.html", "200", "10000"},
			{"1.1.1.3", "/index.html", "504", "11000"},
			{"1.1.1.0", "/inde5x.html", "503", "1000"},
			{"1.1.1.2", "/inde5.html", "200", "50000"},
			{"1.1.1.5", "/inde5x.html", "413", "10000"},
			{"1.1.1.3", "/inde2x.html", "400", "10000"},
			{"1.1.1.4", "/index.html", "200", "11000"},
			{"1.1.1.2", "/index1.html", "505", "10000"},
			{"1.1.1.6", "/index.html", "407", "1000"},
			{"1.1.1.2", "/inde2x.html", "200", "50000"},
			{"1.1.1.7", "/index.html", "401", "10000"},
			{"1.1.1.3", "/index.html", "200", "10000"},
			{"1.1.1.3", "/inde3x.html", "200", "11000"},
			{"1.1.1.2", "/inde3x.html", "407", "1000"},
			{"1.1.1.5", "/index3.html", "200", "50000"},
			{"1.1.1.7", "/index.html", "301", "10000"},
			{"1.1.1.9", "/index.html", "401", "10000"},
			{"1.1.1.11", "/index.html", "200", "11000"},
			{"1.1.1.3", "/index2.html", "200", "11000"},
			{"1.1.1.4", "/index.html", "404", "1000"},
			{"1.1.1.2", "/index.html", "504", "50000"},
			{"1.1.1.3", "/index.html", "200", "10000"},
			{"1.1.1.12", "/index.html", "400", "10000"},
			{"1.1.1.3", "/index.html", "200", "11000"},
			{"1.1.1.2", "/index1.html", "401", "10000"},
		}
	)

	fmt.Println("每个 IP 出现的次数:")
	logIPCount := statisticsCount(logs, infoMap["ip"])
	for key, count := range logIPCount {
		fmt.Printf("%s --- %d\n", key, count)
	}

	fmt.Println("每个状态码出现次数:")
	logCodeStatusCount := statisticsCount(logs, infoMap["status"])
	for key, count := range logCodeStatusCount {
		fmt.Printf("%s --- %d\n", key, count)
	}

	fmt.Println("每个IP在每个URL上产生的流程统计:")
	logIPURLFlowCount := statisticsflow(logs, infoMap["ip"], infoMap["url"])
	for key, count := range logIPURLFlowCount {
		fmt.Printf("%s --- %d\n", key, count)
	}

	fmt.Println("IP 出现次数最多的Top 10")
	for _, key := range choiceTop(insertSort(logIPCount, true), 10) {
		fmt.Printf("%s --- %d\n", key, logIPCount[key])
	}
	fmt.Println("状态码出现次数最多的Top 10")
	for _, key := range choiceTop(insertSort(logCodeStatusCount, true), 10) {
		fmt.Printf("%s --- %d\n", key, logCodeStatusCount[key])
	}
	fmt.Println("每个IP在每个URL上产生的流程最多的Top 10")
	for _, key := range choiceTop(insertSort(logIPURLFlowCount, true), 10) {
		fmt.Printf("%s --- %d\n", key, logIPURLFlowCount[key])
	}
}
