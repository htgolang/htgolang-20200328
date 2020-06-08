package codeCount

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// 计算单个文件的代码数量
func fileCodeCount(file string) (fileCount int) {
	var multilineNotes bool
	f, err := os.Open(file)
	if err != nil {
		return
	}
	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return
		}
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		if strings.HasPrefix(line, "/*") || strings.HasSuffix(line, "*/") {
			multilineNotes = !multilineNotes
			continue
		}
		if multilineNotes {
			continue
		}
		fileCount++
	}
	return
}

// CodeCount 统计root目录下或root文件
func CodeCount(root string) int {
	var wg sync.WaitGroup
	countChan := make(chan int)
	total := make(chan int)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			wg.Add(1)
			go func(f string, countChan chan int) {
				defer wg.Done()
				countChan <- fileCodeCount(f)
			}(path, countChan)
		}
		return nil
	})
	go func(total chan int) {
		var count int
		for c := range countChan {
			count += c
		}
		total <- count
	}(total)
	wg.Wait()
	close(countChan)
	t := <-total
	return t
}
