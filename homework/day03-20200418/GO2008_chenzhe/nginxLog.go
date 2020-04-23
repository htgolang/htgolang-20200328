package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fenge()  {
	fen_ge := strings.Repeat("-",100)
	fmt.Printf("%c[1;37;31m我是邪恶的分割线%s%c[0m\n",0x1B,fen_ge,0x1B)
}
func IpTimes(inLog [][4]string) map[string]int {
	timesMap := map[string]int{}
	for _,j := range inLog{
		timesMap[j[0]]+=1
	}
	for i,j:= range timesMap{
		fmt.Printf("%s出现的次数为%d\n",i,j)
	}
	return timesMap
}
func statusTimes(inLog [][4]string)map[string]int  {
	timesMap := map[string]int{}
	for _,j := range inLog{
		timesMap[j[2]]+=1
	}
	for i,j:= range timesMap{
		fmt.Printf("%s状态码出现的次数为%d\n",i,j)
	}
	return timesMap
}

func urlUsed(inLog [][4]string)map[string]int  {
	urlMap := map[string]int{}
	for _,j := range inLog{
		value,_ :=strconv.Atoi(j[3])
		urlMap[j[0]+j[1]] += value

	}
	for i,j := range urlMap{
		fmt.Printf("%s的流量为%d\n",i,j)
	}
	return urlMap
}

func topTen(inData map[string]int)  {
	//排序
	tmpSlice := []map[string]string{}
	for i,j :=range inData{
		tmpSlice = append(tmpSlice,map[string]string{"name":i,"times":strconv.Itoa(j)})
	}
	leng := len(tmpSlice)
	for i :=0;i<leng;i++{
		for j :=i;j<leng;j++{
			a,_ :=strconv.Atoi(tmpSlice[i]["times"])
			b,_ :=strconv.Atoi(tmpSlice[j]["times"])
			if a<b{
				tmpSlice[i],tmpSlice[j]=tmpSlice[j],tmpSlice[i]
			}
		}
	}
	fmt.Println(tmpSlice)
	}
func main() {
	Log :=   [][4]string {
		{"1.1.1.1", "/index.html", "200", "1000"},
		{"1.1.1.2", "/index.html", "200", "10000"},
		{"1.1.1.1", "/index.html", "200", "10000"},
		{"1.1.1.102", "/index.html", "200", "800"},
		{"1.1.1.14", "/index.html", "200", "60000"},
		{"1.1.1.104", "/index.html", "404", "414"},
	}
	Ip :=IpTimes(Log)
	fenge()
	Status :=statusTimes(Log)
	fenge()
	Url :=urlUsed(Log)
	fenge()
	//执行对应的排序
	topTen(Ip)
	topTen(Status)
	topTen(Url)
}
