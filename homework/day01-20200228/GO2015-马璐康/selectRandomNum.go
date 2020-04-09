package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		var num int
		rand.Seed(time.Now().Unix())
		randNum := rand.Int() % 100
		fmt.Println(randNum)
		var isOK bool
		for i := 1; i <= 5; i++ {
			fmt.Print("请输入你猜的数字:")
			fmt.Scan(&num)
			if randNum > num {
				fmt.Printf("你猜的数字太小，你还有%d次机会，请重新输入：\n", 5-i)
			} else if randNum < num {
				fmt.Printf("你猜的数字太大，你还有%d次机会，请重新输入：\n", 5-i)
			} else {
				fmt.Println("你猜的数字完全正确！")
				isOK = true
				break
			}
		}
		if isOK != true {
			fmt.Println("5次机会用完，退出")
		}
		var txt string
		fmt.Print("请问是否继续猜数字：（y/n）")
		fmt.Scan(&txt)
		if txt != "y" {
			fmt.Println("退出")
			break
		}
	}

}
