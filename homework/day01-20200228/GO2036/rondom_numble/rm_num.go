package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//
	rand.Seed(time.Now().UnixNano())

	//生成100以内的随机整数
	a := rand.Intn(100)
	fmt.Println(a)

	i := 0
	for {
		i++
		var y int
		fmt.Print("请输入100以内的整数：")
		fmt.Scan(&y)
		if y > a {
			fmt.Println("大了")
		} else if y < a {
			fmt.Println("小了")
		} else {
			fmt.Println("恭喜你，答对了！")
			break
		}

		//尝试五次后,如果用户继续玩游戏，则重新循环
		if i == 5 {
			var ch string
			fmt.Print("五次机会已经使用完,是否继续游戏Y|N:")
			fmt.Scan(&ch)
			if ch == "Y" || ch == "y" {
				i = 0
			} else {
				break
			}
		}

	}
}
