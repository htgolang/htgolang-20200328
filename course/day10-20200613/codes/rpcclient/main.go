package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"

	"rpcclient/data"
)

func main() {
	addr := "127.0.0.1:9999"
	conn, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 定义请求对象
	request := &data.CalculatorRequest{2, 5}

	//定义响应对象
	response := &data.CalculatorResponse{}

	// 调用远程方法
	err = conn.Call("calc.Add", request, response)

	// 获取结果
	fmt.Println(err, response.Result)
}
