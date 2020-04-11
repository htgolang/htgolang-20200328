package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 1
	var randNum int
	rand.Seed(time.Now().Unix())
	randNum = rand.Intn(100)
	fmt.Println(randNum)
	var userNum int

	for {
		fmt.Println("请输入您猜测的数字：")
		fmt.Scan(&userNum)
		if count < 5 {
			if userNum == randNum {
				fmt.Printf("经过%d次终于猜对了，太聪明了!\n", count)
				break
			} else if userNum > randNum {
				count++
				fmt.Println("猜的太大了")
			} else {
				count++
				fmt.Println("猜的太小了")
			}
		} else {
			fmt.Println("5次用完还没猜对，游戏结束")
			break
		}
	}

}
