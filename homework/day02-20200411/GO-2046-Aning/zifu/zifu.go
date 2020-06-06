package main

import "fmt"

func main() {
	article := `Adad the day nics AAAA ddd xxxx!`
	times := map[rune]int{}

	for _, ch := range article {
		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
			times[ch] = times[ch] + 1
		}
	}
	chcount := map[int][]rune{} //初始化
	for k, v := range times {
		if _, ok := chcount[v]; ok {
			chcount[v] = append(chcount[v], k)
		} else {
			chcount[v] = []rune{k}
		}
	}
	fmt.Println(chcount)
}
