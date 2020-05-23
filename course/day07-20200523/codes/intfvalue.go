package main

import "fmt"

// 定义有Send和SendAll方法的接口
type Sender interface {
	Send(string, string) error
	SendAll([]string, string) error
}

// 定义只有Send方法的接口
type SingleSender interface {
	Send(string, string) error
}

// 定义结构体,并定义结构体Send和SendAll方法
type EmailSender struct {
}

func (s EmailSender) Send(to string, msg string) error {
	fmt.Printf("发送邮件给: %s, 内容: %s\n", to, msg)
	return nil
}

func (s EmailSender) SendAll(to []string, msg string) error {
	fmt.Printf("发送邮件给: %#v, 内容: %s\n", to, msg)
	return nil
}

func main() {
	// 定义sender接口和singleSender接口对象
	var sender Sender
	var singleSender SingleSender

	// 定义结构体EmailSender对象
	emailSender := EmailSender{}

	/*
		// 将emailSender赋值给sender接口对象
		sender = emailSender

		// 将Sender接口的对象赋值给SingleSender接口对象
		singleSender = sender
		fmt.Printf("%T %#v\n", sender, sender)
		fmt.Printf("%T %#v\n", singleSender, singleSender)
	*/

	/*
		// 将emailSender赋值给singlesender接口对象
		singleSender = emailSender

		// 不能赋值(没有sendAll方法)
		sender = singleSender

		fmt.Printf("%T %#v\n", sender, sender)
		fmt.Printf("%T %#v\n", singleSender, singleSender)
	*/
}
