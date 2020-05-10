package handlefile

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileCp struct {
	src  string
	dest string
}

func NewFileCp(src string, dest string) *FileCp {
	return &FileCp{src: src, dest: dest}
}

func (f *FileCp) walkFunc(path string, info os.FileInfo, err error) error {
	destpath:=""
	//To handle src like this for handling the situation like " fcp yx/ kk " and " fcp yx/TED.txt aa/bb/1.txt"
	if strings.HasSuffix(f.src,"/") || strings.HasSuffix(f.src,`\`){
		//if src has a suffix "/",then I have to keep the "/" in the str after substr.
		destpath = f.dest + path[len(f.src)-1:]
	}else {
		//I don't simply add a "/" between dest and substr of src,
		//because if I'd used tool like this: fcp yx/TED.txt aa/bb/1.txt,
		//there would be a problem.
		destpath = f.dest + path[len(f.src):]
	}

	if info.IsDir() {
		err := os.MkdirAll(destpath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		f1, _ := os.Open(path)
		f2, _ := os.OpenFile(destpath, os.O_CREATE|os.O_TRUNC, os.ModePerm)
		for {
			ctx := make([]byte, 8192)
			n, err := f1.Read(ctx)
			if err == io.EOF {
				break
			}
			_, _ = f2.Write(ctx[:n])
		}
		_ = f2.Close()
		_ = f1.Close()
	}
	return nil
}

func (f *FileCp) CP() error {
	finfo, err := os.Stat(f.src)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("The source file/directory doesn't exist!!")
		} else {
			return errors.New("Please check the format of source file/directory!")
		}
	} else {
		if !finfo.IsDir() && (strings.HasSuffix(f.dest, "/") || strings.HasSuffix(f.dest, "\\")) {
			_ = os.MkdirAll(f.dest, os.ModePerm)
			destfile := f.dest + finfo.Name()
			f1, _ := os.Open(f.src)
			f2, _ := os.OpenFile(destfile, os.O_CREATE|os.O_TRUNC, os.ModePerm)
			for {
				ctx := make([]byte, 8192)
				n, err := f1.Read(ctx)
				if err == io.EOF {
					break
				}
				_, _ = f2.Write(ctx[:n])

			}
			_ = f2.Close()
			_ = f1.Close()
			return nil
		} else {
			return filepath.Walk(f.src, f.walkFunc)
		}
	}
}
