package main

import "fmt"

func main() {
	var msg = "我的名\\n字是KK"

	var msgRaw = `我的名\n字是kk`

	fmt.Printf("%T %s\n", msg, msg)
	fmt.Printf("%T %s\n", msgRaw, msgRaw)

	// 操作
	// 字符串连接 +
	fmt.Println(msg + msgRaw)
	// 关系运算 > >= <= < != ==

	fmt.Println("abc" > "acd")  //false
	fmt.Println("abc" >= "acd") //false

	// 赋值 = +=
	msg += "---kk"
	fmt.Println(msg)

	// 索引 切片 ascii
	msg = "abcdef"
	fmt.Printf("%T %#v %c\n", msg[0], msg[0], msg[0])
	fmt.Println(msg[1:3])

	// len 字节大小
	fmt.Println(len(msg))
	fmt.Println(len(msgRaw))

}
