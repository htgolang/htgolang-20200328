package main

import (
	"io"
	"os"
)

func main() {

	// 创建文件1
	logFile1, _ := os.Create("test/1.log")

	// 创建文件文件2
	logFile2, _ := os.Create("test/2.log")

	// 创建MultiWriter对象，给所有的输出流中都写入内容
	writer := io.MultiWriter(logFile1, logFile2, os.Stdout)

	writer.Write([]byte("Hello world"))

	logFile1.Close()
	logFile2.Close()
}
