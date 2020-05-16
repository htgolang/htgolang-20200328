package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("multireader.go")
	defer file.Close()

	// 利用io.Copy+内存中流对象

	// bytes.buffer
	// buffer := bytes.NewBuffer([]byte(""))

	// 复制
	// io.Copy(buffer, file)
	// fmt.Println(buffer.String())

	// strings.Builder
	builder := new(strings.Builder)

	// 复制
	io.Copy(builder, file)

	fmt.Println(builder.String())
}
