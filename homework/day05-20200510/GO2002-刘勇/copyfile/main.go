package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	src  *string
	dest string
	help bool
)

func menuFlag() {
	src = flag.String("s", "", "src file")
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

func copyfile(src, dest string) {
	//判断源文件是否存在
	if !Exist(src) {
		fmt.Println("源文件", src, "不存在，请检查")
		os.Exit(-1)
	}
	//打开源文件
	srcfile, err := os.Open(src)
	//延迟关闭文件
	defer srcfile.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		//判断目标文件是否存在
		if Exist(dest) {
			fmt.Println(dest, "已经存在，文件名冲突")
			os.Exit(-1)
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
}

func main() {
	menuFlag()

	if help || *src == "" || dest == "" {
		flag.Usage()

	} else {
		copyfile(*src, dest)

	}
}
