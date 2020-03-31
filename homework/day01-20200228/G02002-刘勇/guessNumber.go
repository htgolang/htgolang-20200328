package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for {
		// 随机数种子
		rand.Seed(time.Now().Unix())
		// 赋值随机数
		num := rand.Int() % 101
		// 验证随机数
		// fmt.Println(num)

		var guess int
		fmt.Println("猜数字游戏开始,五次机会,数字范围0-100")

		for times := 1; times <= 5; times++ {
			//打印随机数,以测试猜中后的循环是否正常
			// fmt.Println(num)

			fmt.Print("请输入数字: ")
			fmt.Scan(&guess)

			// switch {
			// case guess == num:
			// 	fmt.Println("真厉害,第", times, "就猜中了")
			// 	// break
			// case guess > num:
			// 	fmt.Println("太大了")
			// default:
			// 	fmt.Println("太小了")
			// }
			if guess == num {
				fmt.Println("真厉害,第", times, "次就猜中了")
				break
			} else if guess > num {
				fmt.Println("太大了")
			} else {
				fmt.Println("太小了")
			}

			//五次猜不中,提示语
			if times == 5 && guess != num {
				fmt.Println("5次机会用完了,游戏失败")
			}

		}
		var game string
		fmt.Println("是否重新开始游戏?(输入Y重新开始,任意键退出):")
		fmt.Scan(&game)
		if game == "Y" || game == "y" || game == "yes" || game == "YES" || game == "Yes" {
			continue
		} else {
			break
		}

	}

}
