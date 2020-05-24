package main

import (
	"fmt"
	"homework0523_persistence/inf"
)

const(
	GOBFILEPATH = `F:\go\practice\homework0523_persistence\1.txt`
	JOSNFILEPATH = `F:\go\practice\homework0523_persistence\2.txt`
)

func Buf(pinf inf.PersisInf,filepath string)  {
	err:=pinf.Decode(filepath)
	if err!=nil{
		fmt.Println(err)
	}
	//fmt.Println(pinf.GetObj())
	fmt.Println(pinf)
}

func Persis(pinf inf.PersisInf,filepath string)  {
	err:=pinf.Encode(filepath)
	if err!=nil{
		fmt.Println(err)
	}
}
