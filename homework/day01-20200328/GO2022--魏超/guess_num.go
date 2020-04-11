/*
猜数字
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randomNum int
	var num int
	var isContinue string
LABELOUTER:
	for {
		rand.Seed(time.Now().Unix())
		randomNum = rand.Intn(100)
	LABELINNER:
		for i := 1; i <= 5; i++ {
			fmt.Print("请输入你猜测的值:")
			fmt.Scan(&num)
			switch {
			case num == randomNum:
				fmt.Println("太聪明了!!!")
				break LABELINNER
			case num > randomNum:
				fmt.Println("输入的值太大了.")
			case num < randomNum:
				fmt.Println("输入的值太小了.")
			}
			if i == 5 {
				fmt.Printf("你太笨了，游戏服结束，应该的值:%d\n", randomNum)
			}
		}
		fmt.Print("想继续继续下一轮请输入[yes/y/Y]:")
		fmt.Scan(&isContinue)
		switch isContinue {
		case "yes", "y", "Y":
			fmt.Println("开始进行下一轮")
		default:
			break LABELOUTER
		}
	}
}
