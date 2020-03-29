package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 创建是否继续的函数，以免后面代码重复
func continue_guess() string {
START_GUESS:
	var continue_byte string
	fmt.Print("还要继续玩猜吗？yes, y, Y则继续,其他则退出:")
	if _, err := fmt.Scanln(&continue_byte); err != nil {
		goto START_GUESS
	}
	//返回用户填写的内容
	return continue_byte
}

func main() {
START:
	//设置guess变量
	var guess int
	//设置continue_lable变量
	var continue_lable string
	//设置最大重试次数
	max_time := 5
	rand.Seed(time.Now().Unix())
	number := rand.Int() % 100
	for i := 1; i <= max_time; i++ {
		fmt.Print("请输入你猜的数字:")
		if _, err := fmt.Scanln(&guess); err != nil {
			goto START
		}
		if guess > number {
			fmt.Printf("你输入的数字大了，你还有%d机会\n", max_time-i)
		} else if guess < number {
			fmt.Printf("你输入的数字小了，你还有%d机会\n", max_time-i)
		} else {
			fmt.Println("恭喜，输入正确\n")
		}
	}
	//提示用户是否继续
	continue_lable = continue_guess()
	if continue_lable == "Y" || continue_lable == "y" || continue_lable == "yes" {
		goto START
	} else {
		goto END
	}
END:
}
