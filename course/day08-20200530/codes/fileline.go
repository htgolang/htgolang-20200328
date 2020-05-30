package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// 计算文件内容行数
func fileLine(path string) int {
	cnt := 0

	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return cnt
	}
	// 延迟关闭
	defer file.Close()

	//使用带缓冲IO读取文件，按行读取，计算行数
	reader := bufio.NewReader(file)
	for {
		// 读取每行数据
		ctx, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		// 过滤空行数据及注释数据
		txt := strings.TrimSpace(string(ctx))
		if txt == "" || strings.HasPrefix(txt, "//") {
			continue
		}

		cnt++
	}
	return cnt
}

func main() {

	dir := "./../.."
	total := 0

	var wg sync.WaitGroup

	channel := make(chan int, 10)
	// 遍历文件夹, 计算每隔go文件的行数, 并计算总数

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && ".go" == filepath.Ext(path) {
			wg.Add(1)
			go func() {
				cnt := fileLine(path)
				channel <- cnt
				wg.Done()
			}()
		}
		return nil
	})

	exit := make(chan struct{})

	go func() {
		for cnt := range channel {
			total += cnt
		}
		exit <- struct{}{}
	}()

	wg.Wait()
	close(channel)

	<-exit
	fmt.Println(total)

	// a. wark之前 x
	// b. wait之前 x
	// c. wait之后 x
	// d. wark之前, goroutine v
	// e. wait之前, goroutine v
	// f. wait之后, goroutine
}
