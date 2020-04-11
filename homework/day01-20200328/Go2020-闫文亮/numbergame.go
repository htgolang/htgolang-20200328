package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().Unix())
		x := rand.Int() % 100
		fmt.Println(x)
		i := 1
        for i <= 5 {
			var number int
			fmt.Print("输入猜测的数字：")
			fmt.Scan(&number)

			if number > x {
				fmt.Print("你猜的数字太大了\n")
			} else if number < x {
				fmt.Print("你猜的数字太小了\n")
			} else {
				fmt.Printf("正确,经过%d次猜对了",i)
				break
			}
			if i == 5 {
				fmt.Print("太笨了")
				var y string
				fmt.Print("还玩吗：")
				fmt.Scan(&y)
				if y == "yes" {
					//重新生成随机数，并重新开始循环i=0
					x = rand.Int() % 100
					fmt.Println(x)
					i = 0
				}
			}else{
				fmt.Print("再见")
			}
            i++

		}
}
