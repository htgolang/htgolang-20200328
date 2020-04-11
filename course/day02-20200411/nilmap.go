package main

import "fmt"

func main() {
	var nilMap map[string]string

	fmt.Println(len(nilMap))
	fmt.Println(nilMap["kk"])

	nilMap["kk"] = "xx"
	fmt.Println(nilMap)
}
