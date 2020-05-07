package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//随机数
	var number int
	//输入的数字
	var numuser int
	//统计次数int
	var count int = 1
	rand.Seed(time.Now().Unix()) //功能种子
	number = rand.Intn(100)
	//fmt.Println(number)

	for count < 6 { //判断次数
		fmt.Println("please input your num: ")
		fmt.Scan(&numuser)
		if count == 5 {
			fmt.Println("太笨了，还要继续猜吗?Y/N: ")
			var again string
			fmt.Scan(&again)
			if again == "y" || again == "Y" {
				number = rand.Intn(100)
				fmt.Println(number)
				count = 1
			}
		}
		if numuser == number {
			fmt.Printf("恭喜你%d次猜中！", count)
			break
		} else if numuser > number {
			fmt.Printf("输入的数字太大了，你还有%d次机会猜测\n", 5-count)
			count += 1
		} else {
			fmt.Printf("输入的数字太小了，你还有%d次机会猜测\n", 5-count)
			count += 1
		}
	}
}
