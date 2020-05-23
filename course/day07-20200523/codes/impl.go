package main

import "fmt"

// 定义只有Send方法的接口
type SingleSender interface {
	Send(string, string) error
}

// 通过匿名嵌入对SingleSender进行扩展(复用)
type Sender interface {
	SingleSender
	SendAll([]string, string) error
}

type EmailSender struct{}

func (s EmailSender) Send(to, msg string) error {
	return nil
}

func (s EmailSender) SendAll(to []string, msg string) error {
	return nil
}

func main() {
	var sender Sender = EmailSender{}
	fmt.Println(sender)
}
