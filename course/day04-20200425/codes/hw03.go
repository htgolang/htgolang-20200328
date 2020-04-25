package main

import (
	"fmt"
	"strconv"
)

func main() {
	lines := [][4]string{
		{"1.1.1.1", "/index.html", "200", "1000"},
		{"1.1.1.2", "/index.html", "200", "10000"},
		{"1.1.1.1", "/index.html", "200", "10000"},
	}

	ip := map[string]int{}
	status := map[int]int{}

	traffic := map[[2]string]int{}

	for _, line := range lines {
		ip[line[0]]++

		code, _ := strconv.Atoi(line[2])
		status[code]++

		key := [2]string{line[0], line[1]}
		tr, _ := strconv.Atoi(line[3])

		traffic[key] += tr
	}

	fmt.Println(ip)
	fmt.Println(status)
	fmt.Println(traffic)
}
