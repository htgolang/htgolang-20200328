package main

import (
	"bufio"
	"os"
)

func main() {

	// 打开文件
	file, _ := os.Open("user.txt")
	defer file.Close()

	// 创建带缓冲IO 读对象
	reader := bufio.NewReader(file)
	// reader.Read()
	// reader.ReadByte()

	// 当读取到字节,返回
	// ctx, _ := reader.ReadBytes(',')
	// fmt.Println(string(ctx))

	// ctx, _ = reader.ReadBytes('\n')
	// fmt.Println(string(ctx))

	// 读取一行
	// ctx, prefix, _ := reader.ReadLine()
	// fmt.Println(string(ctx), prefix)

	// ctx, prefix, _ = reader.ReadLine()
	// // fmt.Println(string(ctx), prefix)

	// 读取一行字符串
	// ctx, _ := reader.ReadString('\n')
	// fmt.Println(ctx)
	// ctx, _ = reader.ReadString('\n')
	// fmt.Println(ctx)

	// ctx, _ := reader.ReadSlice('\n')
	// fmt.Println(string(ctx))

	// 将流输出到标准输出
	reader.WriteTo(os.Stdout)
}
