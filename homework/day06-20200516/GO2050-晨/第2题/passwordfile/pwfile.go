package passwordfile

import (
	"fmt"
	"io"
	"os"
)

func FileIsExists(path string) bool  {
	_,err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err){
		return false
	}else {
		panic(err)
	}
}
func ReadPwFile(path string)string{
	file, err := os.OpenFile(path, os.O_RDONLY,os.ModePerm)
	if err != nil {
		fmt.Println("读取文件错误")
	}
	defer file.Close()
	text := make([]byte,10)
	full_text := make([]byte,0)

	for {
		n, err := file.Read(text)
		if err == io.EOF {
			break
		}
		full_text = append(full_text,text[:n] ...)
	}

	//fmt.Println(full_text)
	//fmt.Println(string(full_text))
	return string(full_text)
}
func WritePwFile(path ,input_pw string)  {
	file,err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()

	file.Write([]byte(input_pw))
}
//func Check_pwfile(path ,input_pw string)  bool{
//	if FileIsExists(path) {
//		real_pw := ReadPwFile(path)
//		if real_pw == input_pw {
//			return true
//		}else {
//			return false
//		}
//	} else {
//		WritePwFile(path,input_pw)
//	}
//}
