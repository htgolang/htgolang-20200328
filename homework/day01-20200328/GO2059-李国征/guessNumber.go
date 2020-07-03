package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 猜数
RELOADSTART:
	// 写个随机数
	// 生成个种子
	rand.Seed(time.Now().UnixNano())
	number := rand.Int() % 100
	// 用户输入数字
	var input int
	var s string
	var opportunity int = 5
	fmt.Printf("请猜数(只能输入整型):")

	// 开始循环
	for i:=1; i<=opportunity;i++{
		fmt.Scan(&input)
		switch {
		case input > number:
			if opportunity-i !=0 {
				fmt.Printf("输入的数值太大了！请重新输入， 你还有%d次猜测机会！\n 请继续猜数：", opportunity-i)
			}
		case input < number:
			if opportunity-i != 0{
				fmt.Printf("输入的数值太小了！请重新输入， 你还有%d次猜测机会！\n 请继续猜数：", opportunity-i)
			}
		case input == number:
			if opportunity -i != 0{
				fmt.Printf("真聪明！ 这你都能猜的对！正确数值:%d\n", number)
			}
		}
	}
	fmt.Println("你也太笨了， 5次机会白白浪费掉了！ 请重新开始游戏吧！ctrl + C 直接退出！")
	fmt.Printf("是否继续(Y/N):")
	fmt.Scan(&s)
	if s == "y" || s == "y"{
		goto RELOADSTART
	}

}
