package countlines

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func fileline(fpath string) int {
	f, _ := os.Open(fpath)
	defer f.Close()
	linecount := 0
	br := bufio.NewReader(f)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		//Don't take count the blank line and annotation
		if strings.TrimSpace(string(line)) != "" && !strings.HasPrefix(strings.TrimSpace(string(line)), "//") {
			linecount++
		}
	}
	return linecount
}

func ReadDir(dpath string) {
	total := 0
	var wg sync.WaitGroup
	syncChan1 := make(chan struct{}, 1)
	intChan := make(chan int, 10)
	_ = filepath.Walk(dpath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			wg.Add(1)
			go func() {
				cnt := fileline(path)
				intChan <- cnt
				wg.Done()
			}()
		}
		return nil
	})
	go func() {
		for e := range intChan {
			total += e
		}
		syncChan1 <- struct{}{}
	}()
	wg.Wait()
	close(intChan)
	<-syncChan1
	fmt.Println(total)
}
