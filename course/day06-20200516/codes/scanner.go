package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("请输入内容:")
		scanner.Scan()
		if "q" == scanner.Text() {
			break
		}
		fmt.Println("你输入的内容是:", scanner.Text())
	}

}
