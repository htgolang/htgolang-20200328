package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//生成100以内的随机数并且赋值给变量result
	rand.Seed(time.Now().Unix())
GAMESTART:
	var result int = rand.Intn(100)
	var userInput int
	var regame string
	fmt.Println("游戏开始，你有5次机会请输入1-100以内的一个随机整数：")
	for i := 1; i <= 5; i++ {
		fmt.Scan(&userInput)
		switch {
		case userInput > result:
			fmt.Printf("你猜的答案是:%d 大于真实值！你已经使用了:%d 次机会，还剩:%d 次机会\n", userInput, i, 5-i)
		case userInput < result:
			fmt.Printf("你猜的答案是:%d 小于真实值！你已经使用了:%d 次机会，还剩:%d 次机会\n", userInput, i, 5-i)
		case userInput == result:
			fmt.Printf("你经过猜测%d 次后猜对了，太聪明了！", i)
			fmt.Println("请问是否要重新开始游戏(Y|N)")
			fmt.Scan(&regame)
		}
		if i == 5 && userInput != result {
			fmt.Printf("你已经用尽5次机会，并未猜出正确答案，太笨了！正确答案为:%d \n", result)
			fmt.Println("请问是否要重新开始游戏(Y|N)")
			fmt.Scan(&regame)
		}

	}
	if regame == "Y" || regame == "y" {
		goto GAMESTART
	} else {
		fmt.Println("游戏结束！")
	}
}
