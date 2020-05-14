package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)
var (
	cwdPath string
	allFilename []string
	srcBasePath string = ``
	destBasePath string = ``
	//srcBasePath string = `G:\rds\test\aaa.txt`
	//destBasePath string = `G:\rds\test\bbb.txt`
)

func cpDir(fullPathName string)  {
	newFileName := strings.Replace(fullPathName,srcBasePath,destBasePath,-1)
	if isExist(newFileName){
		fmt.Println(newFileName," -- 该文件已存在")
		return
	}
 	os.Mkdir(newFileName,os.ModePerm)
}

func cpFile(fullPathName string)  {
	newFileName := strings.Replace(fullPathName,srcBasePath,destBasePath,-1)
	if isExist(newFileName){
		fmt.Println(newFileName," -- 该文件已存在")
		return
	}
	newfile,_ :=os.Create(newFileName)
	oldfile,_ :=os.OpenFile(fullPathName,os.O_RDONLY,os.ModePerm)
	defer oldfile.Close()
	defer newfile.Close()
	txt := make([]byte, 0, 1024*1024)
	ctx := make([]byte, 1024)
	for {
		n, err := oldfile.Read(ctx)
		if err == io.EOF {
			break
		}
		txt = append(txt, ctx[:n]...)
	}
	newfile.Write(txt)


}

func isExist(path string)(bool){
	_, err := os.Stat(path)
	if err != nil{
		if os.IsExist(err){
			return true
		}
		if os.IsNotExist(err){
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func isDir(pathName string)bool  {
	fileInfo,_ := os.Stat(pathName)
	if fileInfo == nil{
		return false
	}
	if fileInfo.IsDir(){
		return true
	}
	return false
}

func listAllFile(pathName string)  {
	if pathName==srcBasePath{
		allFilename = append(allFilename,pathName)
	}
	if isExist(pathName){
		if isDir(pathName){
			file,err := os.Open(pathName)
			defer file.Close()
			if err != nil{
				return
			}
			fileinfo,err := file.Readdir(-1)
			if err != nil{
				return
			}
			for _,filename := range fileinfo{
				fullPathNow :=pathName+string(os.PathSeparator)+filename.Name()
				allFilename = append(allFilename,fullPathNow)
				if isDir(fullPathNow){
					listAllFile(fullPathNow)
				}
			}


		}else {
			allFilename = append(allFilename,pathName)
		}
	}else {
		panic("源文件不存在")
	}
}

func isAbsPath(name string) bool {
	match1,_ := regexp.Match(`^[A-Z]:\\`,[]byte(name))
	match2,_ := regexp.Match(`^/`,[]byte(name))
	if match1 || match2{
		return true
	}else {
		return false
	}
}
func main() {

	//获取当前路径
	cwdPath,_ = os.Getwd()
	//解析命令行参数
	flag.StringVar(&srcBasePath,"s","","源地址")
	flag.StringVar(&destBasePath,"d","","目标地址")


	flag.Usage = func() {
		fmt.Println("usage: cp [-s srcpath] [-d destpaht]")
		flag.PrintDefaults()
	}
	flag.Parse()

	//判断路径是否为绝对路径，并统一组合为绝对路径
	if !isAbsPath(srcBasePath){
		srcBasePath = cwdPath +string(os.PathSeparator)+srcBasePath
	}
	if !isAbsPath(destBasePath){
		destBasePath = cwdPath +string(os.PathSeparator)+destBasePath
	}
	if isDir(destBasePath) {
		s := strings.Split(srcBasePath, string(os.PathSeparator))
		destBasePath = destBasePath + string(os.PathSeparator) + s[len(s)-1]
	}
	//遍历源文件的所有文件
	listAllFile(srcBasePath)
	//复制所有文件
	for _,file:= range allFilename{
		if isDir(file){
			cpDir(file)
		}else {
			cpFile(file)
		}
	}
}

