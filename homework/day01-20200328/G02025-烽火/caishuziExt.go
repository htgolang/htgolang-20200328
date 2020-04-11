package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		count   int
		randNum int
		userNum int
		userChoice  string
		flag    bool
	)
	rand.Seed(time.Now().Unix())
	randNum = rand.Intn(100)
	fmt.Println(randNum)

	for !flag {
		fmt.Println("请输入您猜测的数字：")
		fmt.Scan(&userNum)
		if count < 4 {
			if userNum == randNum {
				fmt.Printf("经过%d次终于猜对了，太聪明了!\n", count+1)
				fmt.Println("是否还要继续?[y/n]")
				fmt.Scan(&userChoice)
				switch userChoice {
				case "Y", "y", "yes":
					randNum = rand.Intn(100)
					fmt.Println(randNum)
					count = 0
				default:
					flag = true
				}
			} else if userNum > randNum {
				count++
				fmt.Println("猜的太大了")
			} else {
				count++
				fmt.Println("猜的太小了")
			}
		} else {
			fmt.Println("5次用完还没猜对，游戏结束! 是否还要继续?[y/n]")
			fmt.Scan(&userChoice)
			switch userChoice {
			case "Y", "y", "yes":
				randNum = rand.Intn(100)
				fmt.Println(randNum)
				count = 0
			default:
				flag = true
			}
		}
	}
}
