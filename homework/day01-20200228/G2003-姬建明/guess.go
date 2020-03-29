package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	for {
		rand.Seed(time.Now().UnixNano())
		var num int = rand.Intn(100)
		var cai int
		for i := 1; i <= 5; i++ {
			fmt.Print("请输入你猜的数字:")
			fmt.Scan(&cai)
			switch {
			case cai > num:
				fmt.Printf("你猜的数字太大了，剩余次数%d\n", 5-i)
			case cai < num:
				fmt.Printf("你猜的数字太小了，剩余次数%d\n", 5-i)
			case cai == num:
				fmt.Printf("恭喜你猜对了\n")
				os.Exit(0)
			}
			if 5-i == 0 {
				fmt.Println("输入错误次数过多，请重新输入。")
				break
			}
		}
	}
}
