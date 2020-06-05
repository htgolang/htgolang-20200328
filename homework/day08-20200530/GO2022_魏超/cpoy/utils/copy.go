package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func copyFile(src, dst string, done chan<- struct{}) {
	defer func() {
		done <- struct{}{}
	}()
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dstFile.Close()

	srcbuf := bufio.NewScanner(srcFile)
	dstbuf := bufio.NewWriter(dstFile)
	defer dstbuf.Flush()
	for srcbuf.Scan() {
		context := srcbuf.Bytes()
		context = append(context, '\n')
		_, err := dstbuf.Write(context)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}

func copyDir(src, dst string, done chan<- struct{}) {
	defer func() {
		done <- struct{}{}
	}()
	doneLocal := make(chan struct{})
	err := os.Mkdir(dst, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 检查源
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcFile.Close()

	files, err := srcFile.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			go copyDir(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()), doneLocal)
		} else {
			go copyFile(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()), doneLocal)
		}
	}
	for range files {
		<-doneLocal
	}
	close(doneLocal)
}

// Copy 拷贝文件或目录
func Copy(src, dst string) {
	fmt.Println(time.Now())
	done := make(chan struct{})
	srcFileInfo, err := os.Stat(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	if srcFileInfo.IsDir() {
		go copyDir(src, dst, done)
	} else {
		go copyFile(src, dst, done)
	}
	<-done
	fmt.Println(time.Now())
	close(done)
}
