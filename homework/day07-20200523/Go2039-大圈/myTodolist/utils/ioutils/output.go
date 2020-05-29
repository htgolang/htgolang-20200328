package ioutils

import "fmt"

//定义一个Error类型的输出
func Error(txt string) {
	fmt.Printf("[---] %s\n",txt)
}

////定义一个Successful类型的输出
//func Success(txt string) {
//	fmt.Printf("[+++] %s\n", txt)
//}
//
////定义一个Output类型的输出
//func Output(txt string) {
//	fmt.Printf("[===] %s\n", txt)
//}

