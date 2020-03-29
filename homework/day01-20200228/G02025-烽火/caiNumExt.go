package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randNum int
	rand.Seed(time.Now().Unix())
	randNum = rand.Intn(100)
	fmt.Printf("rand num is: %d \n", randNum)
	var userNum int
	var userChoice string
	var flag bool = true

	for flag {
		fmt.Println("请输入您猜测的数字：")
		fmt.Scan(&userNum)
		if userNum == randNum {
			fmt.Println("猜对了，太聪明了!")
		} else {
			fmt.Println("猜错了")
		}
		fmt.Println("是否继续:")
		fmt.Scan(&userChoice)

		switch userChoice {
		case "yes", "Y", "y":
			randNum = rand.Intn(100)
			fmt.Printf("new rand num is: %d \n", randNum)
		default:
			fmt.Println("game over!")
			flag = false

		}
	}
}
