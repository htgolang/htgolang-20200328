package main

import (
	"fmt"
	"io"
	"os"
)

// 复制文件(src -> dst)
func CopyFile(src, dst string) error {

	// 以读方式打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	// 延迟关闭
	defer srcFile.Close()

	// 以写方式打开目的文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	// 延迟关闭
	defer dstFile.Close()

	// 定义buffer用于存储读取的文件内容
	// 1024 KB, 1024 * 1024 MB, 10MB
	buffer := make([]byte, 1024)
	for {
		// 从源文件读取内容到buffer中
		n, err := srcFile.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// 将buffer内容写入到目的文件
		n, err = dstFile.Write(buffer[:n])
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	fmt.Println(CopyFile("os.go", "./test2/os.go.2"))
}
