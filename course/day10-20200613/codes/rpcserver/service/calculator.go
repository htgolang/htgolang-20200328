package service

import (
	"log"
	"rpcserver/data"
)

// 定义算法服务
type Calculator struct {
}

// 定义+方法
func (c *Calculator) Add(request *data.CalculatorRequest, response *data.CalculatorResponse) error {
	log.Printf("[+] call add method\n")
	response.Result = request.Left + request.Right
	return nil
}
