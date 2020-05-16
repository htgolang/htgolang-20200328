package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件1
	file1, _ := os.Open("test/1.log")

	// 打开文件2
	file2, _ := os.Open("test/2.log")

	// 创建MultiReader 依次读取所有文件中的内容
	reader := io.MultiReader(file1, file2)

	// 读取文件中的内容
	buffer := make([]byte, 5)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(n, err, string(buffer[:n]))

	}

	file1.Close()
	file2.Close()
}
