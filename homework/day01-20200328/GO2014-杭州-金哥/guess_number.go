package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().Unix())
	var randomNumber int =  rand.Intn(100)
	var userNumber int
	i := 1
	fmt.Println(randomNumber)

	for i<=5 {
		fmt.Print("请输入数字:")
		fmt.Scan(&userNumber)
		if userNumber > randomNumber {
			fmt.Println("你输入的太大了")
			i++
			continue
		}else if userNumber < randomNumber {
			fmt.Println("你输入的太小了")
			i++
			continue
		}else {
			fmt.Printf("经过%d次输入,终于对了，你太聪明了",i)
			i++
			break
		}
	}
	if i>5 {
		fmt.Println("你已经猜了5次了，还没猜正确，你太笨了")
	}
}