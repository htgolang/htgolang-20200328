package ioutils

import (
	"bufio"
	"fmt"
	"github.com/howeyc/gopass"
	"os"
	"strings"
)

//定义一个input函数
func Input(prompt string) string {
	/*
	使用带缓冲的io,可以使用下面的任意一种方式
		1.	bufio.Scanner{}
		2.	bufio.NewReader()
	 */


	//打印调用函数时传入的参数（打印传入的内容）
	fmt.Println(prompt)
	//创建一个读io缓冲
	scanner := bufio.NewScanner(os.Stdin)
	//扫描输入的内容
	scanner.Scan()
	//返回用户输入的内容
	return scanner.Text()

}

func Password(prompt string) string {
	fmt.Print(prompt)
	if ctx, err := gopass.GetPasswd();err != nil {
		return ""
	}else {
		return strings.TrimSpace(string(ctx))
	}
}