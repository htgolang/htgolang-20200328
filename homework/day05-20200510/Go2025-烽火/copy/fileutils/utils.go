package fileutils

import (
	"fmt"
	"io"
	"os"
)

var DirList = make([]string, 0)
var FileList = make([]string, 0)

// read file content
func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ctx := make([]byte, 1024)
	txt := make([]byte, 1024*1024)

	for {
		n, err := file.Read(ctx)
		if err == io.EOF {
			break
		}
		txt = append(txt, ctx[:n]...)
	}
	return string(txt)
}

// write content into file
func WriteFile(path, txt string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	file.WriteString(txt)
}

// judge file exists and return filename
func FileIsExists(path string) (string, bool) {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return fileInfo.Name(), true
	} else if os.IsNotExist(err) {
		return "", false
	} else {
		panic(err)
	}
}

// judge dir
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return fileInfo.IsDir()
}

// traverse path (subfiles & subdir)
func TraverseCopy(src, dest string) {
	file, err := os.Open(src)
	if err != nil {
		return
	}
	defer file.Close()

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		return
	}

	for _, fileInfo := range fileInfos {
		spath := src + "/" + fileInfo.Name()
		dpath := dest + "/" + fileInfo.Name()

		if fileInfo.IsDir() {
			fmt.Println("dir: ", spath)
			os.Mkdir(dpath, fileInfo.Mode())
			TraverseCopy(spath, dpath)
		} else {
			fmt.Println("file: ", spath)
			WriteFile(dpath, ReadFile(spath))
		}
	}
}
