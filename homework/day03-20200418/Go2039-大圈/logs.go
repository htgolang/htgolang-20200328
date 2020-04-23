package main

import (
	"fmt"
	"strconv"
)

/*
3. 统计
    a.每个IP出现次数
    b.每个状态码出现次数
    c.每个IP在每个URL上产生的流量 IP, URL key
    d. a, b, c top 10
 */
//定义一个切片(切片内的元素是4个元素的数组), 存放日志
var logInfo = [][4]string {
	{"192.168.9.1", "/index.html","200","100"},
	{"192.168.9.2", "/index.php","301","190"},
	{"192.168.9.3", "/index.js","400","10"},
	{"192.168.9.41", "/index.html","200","1000"},
	{"192.168.9.22", "/index.html","502","100"},
	{"192.168.9.2", "/index.html","504","101"},
	{"192.168.9.31", "/index.js","405","10"},
	{"192.168.9.40", "/index.html","500","1000"},
	{"192.168.9.2", "/index.html","504","100"},
	{"192.168.9.12", "/index.html","200","100"},
	{"192.168.9.3", "/index.js","401","10"},
	{"192.168.9.41", "/index.html","200","1000"},
	{"192.168.9.3", "/in.js","200","10"},
	{"192.168.9.2", "/index.html","302","101"},
	{"192.168.9.25", "/index.html","200","103"},
	{"192.168.9.20", "/index.html","403","0"},
	{"192.168.8.20", "/test.html","403","1"},
	{"192.168.8.21", "/deny.html","404","2"},
}

//统计每个IP出现的次数
func getCountIp(logs [][4]string) map[string]int {
	mapIp := make(map[string]int)
	for _, v := range logInfo {
		mapIp[v[0]]++
	}
	for k,v := range mapIp{
		fmt.Printf("IP：%s count: %d\n",k,v)
	}
	return mapIp
}

//统计每个状态码出现的次数
func getCountStatus(logs [][4]string) map[string]int {
	mapStatus := make(map[string]int)
	for _, v := range logInfo {
		mapStatus[v[2]]++
	}
	for k,v := range mapStatus{
		fmt.Printf("codeStatus：%s count: %d\n",k,v)
	}
	return mapStatus
}
//统计每个IP+URL的流量总和
func getTrafficStatistics(logs [][4]string) map[string]int {
	mapTS := make(map[string]int)
	for _, v:= range logInfo {
		i, _ := strconv.Atoi(v[3])
		mapTS[v[0]+v[1]] = mapTS[v[0]+v[1]] + i
	}
	for k,v := range mapTS {
		fmt.Println(k,v)
	}
	return mapTS
}

//定义一个空切片(二维切片)
var sliceIp [][]string
func mapIpToSlice(m map[string]int)  {
	//将mapIp转换成一个二维切片，切片内的每个元素就是一个键值对
	for k,v := range m{
		var s1 []string
		s1 = append(s1,k)
		s1 = append(s1,strconv.Itoa(v))
		sliceIp = append(sliceIp,s1)
	}
}

var sliceStatus [][]string
func mapStatusToSlice(m map[string]int)  {
	//将mapStatus转换成一个二维切片，切片内的每个元素就是一个键值对
	for k,v := range m{
		var s1 []string
		s1 = append(s1,k)
		s1 = append(s1,strconv.Itoa(v))
		sliceStatus = append(sliceStatus,s1)
	}
}

var sliceIpUrl [][]string
func mapIpUrlToSlice(m map[string]int)  {
	//将mapIpUrl转换成一个二维切片，切片内的每个元素就是一个键值对
	for k,v := range m{
		var s1 []string
		s1 = append(s1,k)
		s1 = append(s1,strconv.Itoa(v))
		sliceIpUrl = append(sliceIpUrl,s1)
	}
}

func mySort(s [][]string)  {
	//对二维切片按key进行排序,以防每次从map得到切片时顺序不一致
	for j:=0;j<len(s)-1;j++ {
		for k:=0;k<len(s)-j-1;k++ {
			if s[k][0] < s[k+1][0] {
				tmp := s[k+1]
				s[k+1] = s[k]
				s[k] = tmp
			}
		}
	}

	//对二维切片按value进行排序,降序
	for j:=0;j<len(s)-1;j++ {
		for k:=0;k<len(s)-j-1;k++ {
			a, _ := strconv.Atoi(s[k][1])
			b, _ := strconv.Atoi(s[k+1][1])
			if a < b {
				tmp := s[k+1]
				s[k+1] = s[k]
				s[k] = tmp
			}
		}
	}
}

//然后获取前10位IP地址
func getTopTen(s [][]string)  {
	for i,v := range s {
		if i < 10 && len(s) >9 {
			fmt.Printf("%s： %s\n",string(v[0]),v[1])
		}else {
			break
		}
	}
}

func main() {
	fmt.Println("===================IP次数统计===========================")
	myMapIP := getCountIp(logInfo)
	mapIpToSlice(myMapIP)
	mySort(sliceIp)
	fmt.Println("=====================前10IP=========================")
	getTopTen(sliceIp)

	fmt.Println("===================状态码次数统计===========================")
	myMapStatus := getCountStatus(logInfo)
	mapStatusToSlice(myMapStatus)
	mySort(sliceStatus)
	fmt.Println("====================前10状态码==========================")
	getTopTen(sliceStatus)

	fmt.Println("===================IP+URL流量统计===========================")
	myMapIpUrl := getTrafficStatistics(logInfo)
	mapIpUrlToSlice(myMapIpUrl)
	mySort(sliceIpUrl)
	fmt.Println("====================前10 IP+URL流量==========================")
	getTopTen(sliceIpUrl)
}



