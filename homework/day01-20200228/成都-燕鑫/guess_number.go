package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	guess_number()
}

func isContinue() (contmp string) {
	for {
		contmp := ""
		fmt.Print("Do you wanna continue?(Y/N): ")
		fmt.Scan(&contmp)
		if contmp != "Y" && contmp != "N" {
			print("Please enter Y or N !")
			continue
		}
		return contmp
	}
}

func compareNum(a int, b int) (result bool) {
	switch {
	case a > b:
		fmt.Println("CLUE:Your guess is bigger!")
		return false
	case a < b:
		fmt.Println("CLUE:Your guess is smaller!")
		return false
	case a == b:
		fmt.Println("Bingo!How clever you are!")
		return true
	}
	return false
}

func guess_number() {
	rand.Seed(time.Now().Unix())
	fmt.Println("Welcome to GuessNumber Game!")
	con := "Y"
	for con == "Y" {
		num := rand.Int() % 101
		answer := ""
		fmt.Print("Please enter your guess: ")
		for i := 0; i < 5; {
			fmt.Scan(&answer)
			anstmp, err := strconv.ParseInt(answer, 10, 0)
			if err != nil {
				fmt.Print("Please enter a number!Enter again: ")
				continue
			}
			if compareNum(int(anstmp), num) == true {
				con = isContinue()
				break
			} else {
				i++
				if i < 4 {
					fmt.Printf("Sorry,please guess it again(%d CHANCES LEFT): ", 5-i)
				} else if i == 4 {
					fmt.Printf("Come on!Last one chance: ")
				} else {
					fmt.Printf("So sorry!The correct answer is %d!\n", num)
					con = isContinue()
				}
			}

		}
	}
}
