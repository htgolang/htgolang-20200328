package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// 通过字节切片创建Buffer结构体指针对象
	buffer1 := bytes.NewBuffer([]byte("abc,"))

	// 通过字符串创建Buffer结构体指针对象
	buffer2 := bytes.NewBufferString("123,")

	// 写操作
	buffer1.Write([]byte("xyz,"))
	buffer2.Write([]byte("789,"))

	buffer1.WriteString("mn")
	buffer2.WriteString("55")

	// 读内容
	ctx := make([]byte, 3)
	n, _ := buffer1.Read(ctx)
	fmt.Println(string(ctx[:n]))

	n, _ = buffer2.Read(ctx)
	fmt.Println(string(ctx[:n]))

	// 当碰到指定的字节元素后停止读取
	ctx, _ = buffer2.ReadBytes(',')
	fmt.Println(string(ctx))

	ctx, _ = buffer2.ReadBytes(',')
	fmt.Println(string(ctx))

	txt, _ := buffer1.ReadString(',')
	fmt.Println(txt)

	txt, _ = buffer1.ReadString(',')
	fmt.Println(txt)

	// 将buffer流中剩余内容转换为字节切片
	fmt.Println(string(buffer1.Bytes()))

	// 将buffer流中剩余内容转换为字节切片
	fmt.Println(buffer1.String())

	buffer1.WriteTo(os.Stdout)

	buffer2.Reset() // 清空内容
	buffer2.WriteTo(os.Stdout)

}
