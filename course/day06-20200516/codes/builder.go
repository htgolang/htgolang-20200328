package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix())
}

// 生成随机数的(n位)
func RandString(n int) string {
	chars := "qwertyuiopasdfghjklzxcvbnm1234567890QWERTYUIOPASDFGHJKLZXCVBNM"
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteByte(chars[rand.Intn(len(chars))])
	}
	return builder.String()
}

func main() {
	// 定义strings.Builder结构体对象 类似在内存中的流对象(写文件对象)
	var builder strings.Builder

	//builder中写入内容: 字节切片
	builder.Write([]byte("我是KK\n"))

	//builder中写入内容: 字符串
	builder.WriteString("我是阿宁")

	//builder中写入内容: rune，码点
	builder.WriteRune('b')

	//builder中写入内容: byte, 字节
	builder.WriteByte('b')

	//获取写入的内容 字符串
	fmt.Println(builder.String())

	// 获取写入内容的字节数量
	fmt.Println(builder.Len())

	// 清空builder对象
	builder.Reset()
	fmt.Println(builder.String())
	fmt.Println(builder.Len())

	fmt.Println(RandString(5))
	fmt.Println(RandString(6))
}
