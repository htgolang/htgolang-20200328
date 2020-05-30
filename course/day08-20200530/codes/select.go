package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 5)
		ch2 <- 2
	}()

	select {
	case e := <-ch1:
		fmt.Println("ch1:", e)
	case e := <-ch2:
		fmt.Println("ch2:", e)
	default:
		fmt.Println("default")
	}

}
