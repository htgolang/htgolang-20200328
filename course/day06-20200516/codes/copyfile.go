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

	// 从srcFile复制内容到dstFile
	_, err = io.Copy(dstFile, srcFile)

	// 从srcFile复制内容到dstFile
	// buffer := make([]byte, 1024*1024)
	// _, err = io.CopyBuffer(dstFile, srcFile, buffer)

	// 只拷贝前N前字节
	// _, err = io.CopyN(dstFile, srcFile, 10)

	return err
}

func main() {
	fmt.Println(CopyFile("os.go", "./os.go.5"))
}
