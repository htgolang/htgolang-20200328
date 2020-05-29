package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

//复制
func copyfile(srcfile, dstfile string) {
	//第一步打开文件
	src, err := os.Open(srcfile)
	if err != nil {
		fmt.Println(err)
	} else {
		//最后文件关闭
		defer src.Close()
		//第二步创建目标文件
		dst, err := os.Create(dstfile)
		if err != nil {
			fmt.Println(err)
		} else {
			//最后关闭文件
			defer dst.Close()
			//第三步读取源文件
			bytes, err := ioutil.ReadAll(src)
			if err != nil {
				fmt.Println(err)
			} else {
				//第四步写入目标文件
				dst.Write(bytes)
			}
		}
	}
}
func main() {
	//定义变量  存放参数
	var (
		srcfile string
		dstfile string
		help    bool
		h       bool
	)
	//参数定义以寄默认值
	flag.StringVar(&srcfile, "s", "", "src file")
	flag.StringVar(&dstfile, "d", "", "dst file")
	flag.BoolVar(&help, "help", false, "help")
	flag.BoolVar(&h, "h", false, "help")

	//帮助说明提示
	flag.Usage = func() {
		fmt.Println("usage:  cp -s [srcfile] -d [dstfile]")
		flag.PrintDefaults()
	}
	//解析输入格式
	flag.Parse()
	if help || h || srcfile == "" || dstfile == "" {
		flag.Usage()
	} else {
		copyfile(srcfile, dstfile)
		fmt.Println("cp end")
	}
}
