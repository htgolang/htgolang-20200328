package main

import (
	"io/ioutil"
	"time"
)

func main() {
	// 创建临时文件，所在目录./test, 文件命名以log为前缀+随机数字
	// 返回文件结构体指针对象
	file, _ := ioutil.TempFile("./test", "log")
	file.WriteString(time.Now().Format("15:04:05"))
}
