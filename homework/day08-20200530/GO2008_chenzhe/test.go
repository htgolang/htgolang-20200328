package main

import (
	pool2 "GO2008_chenzhe/pool"
	"fmt"
)

func main() {
	p:=pool2.NewPool(100,4)

	for i:=0;i<10 ;i++  {
		p.AddTask(func() interface{} {

			return "任务结束"
		})
	}
	resluts:=p.Start()
	for i:=range resluts{
		fmt.Println(i)
	}


}
