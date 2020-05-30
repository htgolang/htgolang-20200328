package main

import (
	"fmt"
	"sync"
)

type Connection struct {
	id   int
	addr string
}

func main() {
	addr := "192.168.1.100"
	index := 0

	pool := sync.Pool{
		New: func() interface{} {
			index++
			fmt.Println("new:", index)
			return &Connection{index, addr}
		},
	}

	c := pool.Get()
	fmt.Println(c)
	pool.Put(c)

	c2 := pool.Get()
	fmt.Println(c2)

	c3 := pool.Get()
	fmt.Println(c3)
}
