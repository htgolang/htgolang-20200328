package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	for {
		rand.Seed(time.Now().Unix())
		var (
			num = rand.Intn(100)			
			flag bool = true
			play string
			guess int
		)
		fmt.Println(num)	
		for i:=1;i<=5;i++{		
			fmt.Println(i,".输入数字")
			fmt.Scanln(&guess)
			if guess < num {
				fmt.Println("太小了")
			} else if guess>num {
				fmt.Println("太大了")
			} else {
				fmt.Println("猜对了")
				flag = false
				break
			}
		}
		if flag {
			fmt.Println("你是猪，没猜对")
		}
		
		fmt.Println("你还要玩吗y/n?")
		fmt.Scanln(&play)
		if play == "n"{
			break
		}
}
	
}



