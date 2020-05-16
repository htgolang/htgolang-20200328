package main

import (
	"bufio"
	"os"
)

func main() {
	// 打开文件
	file, _ := os.Create("test/test.log")
	defer file.Close()

	// 创建带缓冲IO的写对象
	writer := bufio.NewWriter(file)

	// 延迟刷新缓冲内容
	defer writer.Flush()

	// 像缓冲中写入数据
	writer.Write([]byte("abcdef\n"))
	writer.WriteString("123456789")
}
