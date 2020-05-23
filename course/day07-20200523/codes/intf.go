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

// 值接收者
func (s EmailSender) Send(to string, msg string) error {
	fmt.Printf("发送邮件给: %s, 内容: %s\n", to, msg)
	return nil
}

func (s EmailSender) SendAll(tos []string, msg string) error {
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

	// 给接口变量sender进行赋值
	// 创建EmailSender对象
	emailSender := EmailSender{"stmp.qq.com", 456, "kk", "123"}
	sender = emailSender

	fmt.Printf("%T %#v\n", sender, sender)

	// 接口对象不能访问属性
	// fmt.Println(sender.addr)

	// 接口只能调用声明的方法
	sender.Send("浩泽宇", "作业结果")
	sender.SendAll([]string{"浩泽宇", "大圈", "烽火"}, "作业结果")
	// sender.SendCc("浩泽宇", "大A", "作业结果")

	pEmailSender := &EmailSender{"stmp.qq.com", 456, "kk", "123"}
	sender = pEmailSender

	fmt.Printf("%T %#v\n", sender, sender)
}
