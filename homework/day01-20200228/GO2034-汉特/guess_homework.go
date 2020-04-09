package main
import "fmt"
import "time"
import "math/rand"


func main() {
	var guessVal int
	var answerVal int
	var againVal string
	var i bool = true
	for i==true {
		rand.Seed(time.Now().Unix())
		answerVal = rand.Intn(100)
		fmt.Print("答案已就位，猜测开始")
		fmt.Println(answerVal)					//测试过程，事后注释
		for numVal:=1;numVal>=1;numVal++ {      //死循环
			if  numVal==6{							 //判断是不是猜次数超过5次
				fmt.Println("您已经猜测次数超过5次，真菜")
				break
			}
			fmt.Scan(&guessVal)
			if guessVal!=answerVal{            //判断是否猜对
				switch {
				case guessVal>answerVal:		//判断是大是小
				fmt.Print("大了，请再次猜测")
				case guessVal<answerVal:		//判断是大是小
				fmt.Print("小了，请再次猜测")
				}
			}else {
				fmt.Printf("恭喜你猜对了，你一共猜了%d次\n",numVal)
				break
			}
		}
		fmt.Println("请问还要再来一次吗？按“y”再来一次，按“n”退出程序")
		fmt.Scan(&againVal)
		//开始判断是否再来  if
		// if againVal==1{
		// 	i=5
		// }else if againVal=2{
		// 	i=11
		// }else {
		// 	fmt.Println("输入错误，退出程序")
		// 	i=11
		// }
		//开始判断是否再来  switch
		switch againVal{
		case "y","yes","Yes","YES":
			i=true
		case "no","N","NO":
			i=false
		default:
			fmt.Println("输入错误,退出程序")
			i=false
		}
	}
}