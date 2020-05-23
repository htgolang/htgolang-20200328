package main

import "fmt"

// 定义接口Sender
type Sender interface {
	// 定义接口的行为(方法名，参数列表，返回值列表)
	Send(string, string) error
	SendAll([]string, string) error
}

type EmailSender struct {
	addr     string
	port     int
	user     string
	password string
}

// 指针接收者
func (s *EmailSender) Send(to string, msg string) error {
	fmt.Printf("发送邮件给: %s, 内容: %s\n", to, msg)
	return nil
}

func (s *EmailSender) SendAll(tos []string, msg string) error {
	fmt.Printf("发送邮件给: %#v, 内容: %s\n", tos, msg)
	return nil
}

func (s EmailSender) SendCc(to string, cc string, msg string) error {
	fmt.Printf("发送邮件给: %s, 抄送: %s, 内容: %s\n", to, cc, msg)
	return nil
}

func main() {
	// 定义接口Sender的变量
	var sender Sender
	fmt.Printf("%T %#v\n", sender, sender)

	// 定义值对象 并进行赋值
	// var emailSender = EmailSender{}
	// sender = emailSender
	// fmt.Printf("%T %#v\n", sender, sender)

	var pemailSender = &EmailSender{}
	sender = emailSender
	fmt.Printf("%T %#v\n", sender, sender)
}
