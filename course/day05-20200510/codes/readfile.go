package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 文件路径
	// 相对路径 程序执行执行路径
	// password
	// ./password
	// ../passwrod
	// 绝对路径 根目录。磁盘目录写起的路径
	// /opt/todolist/etc/password
	// e:/todolist/etc/password

	// 读文件
	// 程序 不会退出，重复读取文件
	file, err := os.Open("password.txt")
	if err != nil {
		return
	}
	defer file.Close()
	ctx := make([]byte, 10)

	for {
		n, err := file.Read(ctx)

		if err == io.EOF {
			break
		}
		fmt.Println(n, err, string(ctx[:n]), len(string(ctx[:n])))
	}

}
