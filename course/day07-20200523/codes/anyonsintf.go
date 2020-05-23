package main

import "fmt"

type EmailSender struct {
}

func (s EmailSender) Send(to, msg string) error {
	return nil
}

func main() {
	// 定义匿名接口
	var sender interface {
		Send(to, msg string) error
	}

	fmt.Println(sender)
	sender = EmailSender{}
	fmt.Println(sender)
}
