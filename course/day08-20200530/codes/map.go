package main

import (
	"fmt"
	"sync"
)

func main() {
	var smap sync.Map

	smap.Store("1", 1)
	fmt.Println(smap.Load("1"))
	smap.Delete("1")
	fmt.Println(smap.Load("1"))

	smap.Store("1", 1)
	smap.Store("2", 2)
	smap.Store("3", 3)

	smap.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}
