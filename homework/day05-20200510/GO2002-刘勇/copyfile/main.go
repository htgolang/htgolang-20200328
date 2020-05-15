package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	src  string
	dest string
	help bool
)

func menuFlag() {
	flag.StringVar(&src, "s", "", "src file")
	flag.StringVar(&dest, "d", "", "dest file")
	flag.BoolVar(&help, "h", false, "help")
	flag.Usage = func() {
		fmt.Println(`
Usage:copyfile -s src srcfile  -d destfile
Options:`)
		flag.PrintDefaults()
	}

	flag.Parse()

}

func Exist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func copyFile(src, dest string) {

	//打开源文件
	srcfile, err := os.Open(src)
	//延迟关闭文件
	defer srcfile.Close()

	if err != nil {
		fmt.Println(err)
	}
	//创建目标文件
	destfile, err := os.Create(dest)
	//延迟关闭
	defer destfile.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		//没有报错就执行下面代码
		//创建10MB大小的bytes
		bytes := make([]byte, 10*1024*1024)

		for {
			n, err := srcfile.Read(bytes)
			if err != nil {
				if err != io.EOF {
				}
				break
			}
			destfile.Write(bytes[:n])
		}
	}

}
func copyDir(src, dest string) {
	//创建最上层文件夹
	os.MkdirAll(dest, 0644)
	files, _ := ioutil.ReadDir(src)
	for _, file := range files {
		fmt.Println("copyDir:", file.Name, src, dest)
		if file.IsDir() {
			copyDir(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))
		} else {
			copyFile(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))

		}
	}
}

func main() {
	menuFlag()

	if help || src == "" || dest == "" {
		flag.Usage()
	}

	//判断文件/文件夹是否存在
	if !Exist(src) {
		fmt.Println(src, "文件/文件夹不存在")
		os.Exit(-1)
	}
	if Exist(dest) {
		fmt.Println(dest, "文件/文件夹已存在")
		os.Exit(-1)
	}

	//判断是文件还是文件夹，使用不同的函数
	v, err := os.Stat(src)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if v.IsDir() {
		copyDir(src, dest)
		fmt.Println(src, dest)
	} else {
		copyFile(src, dest)
	}

}
