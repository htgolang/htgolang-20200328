package main

import (
	"bufio"
	"fmt"
	. "homework0418/question2"
	. "homework0418/question3"
	. "homework0418/question4"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Println("\n\n2:找TOP  3:分析IP统计信息  4:任务管理")
		fmt.Print("\nPlease enter your choice: ")
		var inputReader *bufio.Reader
		inputReader = bufio.NewReader(os.Stdin)
		qno,_ := inputReader.ReadString('\n')
		switch strings.TrimSpace(qno) {
		case "2":
			fmt.Println("作业第二题：找TOP10")
			FindTOPmain()
		case "3":
			fmt.Println("作业第三题：分析IP统计信息")
			IpAnalysisMain()
		case "4":
			fmt.Println("作业第四题：任务管理")
			TaskMain()
		case "q":
			goto END
		default:
			fmt.Println("Only can enter 2,3,4,q")
		}
	}
END:
	fmt.Println("BYE!")
}
