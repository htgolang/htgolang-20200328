package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type FileFilter func(string) bool

type FileCallback func(string)

func Dir(path string, filter FileFilter, callback FileCallback) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		return
	}
	for _, fileInfo := range fileInfos {
		fpath := path + "/" + fileInfo.Name()
		if fileInfo.IsDir() {
			Dir(fpath, filter, callback)
		}
		if filter == nil || filter(fpath) {
			if callback != nil {
				callback(fpath)
			}
		}
	}
}

func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()
	txt := make([]byte, 0, 1024*1024)
	ctx := make([]byte, 1024)
	for {
		n, err := file.Read(ctx)
		if err == io.EOF {
			break
		}
		txt = append(txt, ctx[:n]...)
	}
	return string(txt)
}

func main() {
	Dir("password", func(path string) bool {
		return strings.HasSuffix(path, ".go")
	}, func(path string) {
		fmt.Println("file:", path)
		fmt.Println("content:", ReadFile(path))
	})
}
