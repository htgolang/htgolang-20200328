package main

import (
	"flag"
	"fmt"
	"homework0529/copyfiles"
	"homework0529/countlines"
	"homework0529/taskpool"
	"homework0529/testticker"
	"os"
)

var (
	h    bool
	t    bool
	p    bool
	cp   bool
	src  string
	dest string
	dir    string
)

func init() {
	flag.BoolVar(&h, "h", false, "Help information")
	flag.BoolVar(&t, "t", false, "Test Tricker")
	flag.BoolVar(&p, "p", false, "Test TaskPool")
	flag.BoolVar(&cp, "cp", false, "CP src to dest, must specify -src and -dest\n -cp src xxx -dest xxx")
	flag.StringVar(&src, "src", "", "source of dir or file")
	flag.StringVar(&dest, "dest", "", "dest of dir or file")
	flag.StringVar(&dir, "dir", "", "Counting the total lines of files of specific dir")

}

func main() {
	flagParse()
}

func flagParse() {
	flag.Parse()
	if h || flag.NFlag() == 0 {
		flag.Usage()
		return
	}
	if t {
		testticker.HWTrigger()
		return
	}
	if p {
		taskpool.TestTaskPool()
		return
	}
	if cp {
		if src == "" || dest == "" {
			flag.Usage()
			fmt.Println("Missing src or dest!")
			return
		}
		fcp := copyfiles.NewFileCp(src, dest)
		err := fcp.CP()
		if err != nil {
			fmt.Println(err)
			os.Exit(12)
		}
		return
	}
	if dir!=""{
		countlines.ReadDir(dir)
	}
}
