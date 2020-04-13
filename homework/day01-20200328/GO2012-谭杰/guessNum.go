package main


import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)


func main()  {


	// generate randnumer
	rand.Seed(time.Now().Unix())
	randInt := rand.Intn(10)

	var guess_fre int = 5


	// fmt.Println(guess_fre)
	// fmt.Println(randInt)

	plzNum := ""

	for {
		fmt.Print("plz input number:")

		for i:=1;i<=guess_fre;i++{
			fmt.Scan(&plzNum)
			plzNumNew,_ := strconv.Atoi(plzNum)

			switch  {
			case randInt == plzNumNew :
				fmt.Printf(">>guess success, total %#v", i)
				goto Loop1
			case randInt > plzNumNew :
				fmt.Println("guess number is little ")
				if i< guess_fre{
					fmt.Print("plz try again: ")
					continue
				}else if i== plzNumNew{
					fmt.Printf(">>guess success, total %#v", i)
					goto Loop1
				}
			case randInt < plzNumNew :
				fmt.Println("guess number is bigger ,plz try again: ")
				if i< guess_fre{
					fmt.Print("plz try again: ")
					continue
				}else if i== plzNumNew{
					fmt.Printf(">>guess success, total %#v", i)
					goto Loop1
				}
			default:
				fmt.Println("输入错误")
			}
		}
		Loop1:
			fmt.Println("Success Exiting....\n")

		fmt.Println("=================")
		isContinue := ""
		fmt.Print("is continue(Y|N)？")
		fmt.Scan(&isContinue)
		if isContinue == "Y" {
			continue
		}else{
			break
		}

	}
}