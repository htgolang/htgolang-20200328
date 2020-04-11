package main

import (
	"fmt"
	"math/rand"
	"syscall"
	"time"
)

var (
	coin       int    //硬币次数
	customer   int    //用户输入数值
	sysNumber  int    //生成的数值
	played     int    //游玩的次数
	userChoice string //用户的选择
)

func main() {
	//随机数种子只需要设置一次
	rand.Seed(time.Now().Unix())
	//生成随机整数赋值
	sysNumber = rand.Intn(100)
	played = 0 //初始化游玩次数
	for coin = 5; coin >= 1; coin-- {
		played++
		fmt.Printf("正确数字是%d\n", sysNumber)
		fmt.Println("请猜个数字")
		fmt.Scan(&customer)
		switch {
		case customer > sysNumber:
			fmt.Printf("猜大了,你还剩%d次\n", coin-1)
		case customer < sysNumber:
			fmt.Printf("猜小了,你还剩%d次\n", coin-1)
		case customer == sysNumber:
			fmt.Printf("经过%d次,你猜对啦，还玩吗？\n", played)
			choice()
		}
		if coin == 1 {
			fmt.Println("你已经没有机会了,你还想玩吗？")
			choice()
		}
	}
}

func choice() {
	fmt.Scan(&userChoice)
	switch userChoice {
	case "yes", "Y", "y":
		main()
	case "no", "N", "n":
		fmt.Println("下回见~")
		syscall.Exit(0)
	default:
		syscall.Exit(0)
	}
}

func replay() {
	main()
}
