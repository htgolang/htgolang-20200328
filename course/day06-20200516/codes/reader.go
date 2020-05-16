package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 通过New函数创建Reader结构体指针，初始化字符串
	reader := strings.NewReader("abc123456789xyz")

	// 定义切片，用于存放读取的内容
	ctx := make([]byte, 10)
	for {

		// 读取内容到切片，n 读取字节数量, err 用于判断是否发生错误或者是否读取到文件末尾
		n, err := reader.Read(ctx)
		if err != nil { // 由错误结束循环
			break
		}
		fmt.Println(n, err, string(ctx[:n])) // 打印读取内容
	}

	// 重新设置流指针位置
	reader.Seek(0, os.SEEK_SET)

	n, err := reader.Read(ctx)
	fmt.Println(n, err, string(ctx[:n]))
	fmt.Println(reader.Size())

	// 重置reader对象中内容
	reader.Reset("123123123")

	// n, err = reader.Read(ctx)
	// fmt.Println(n, err, string(ctx[:n]))

	// 将reader对象中内容输出到输出流
	reader.WriteTo(os.Stdout)

}
