package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 种子，只需要设置一次
	rand.Seed(time.Now().Unix())
	// rand.Seed(1)

	// [0 - 100)
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int() % 100)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}

}
