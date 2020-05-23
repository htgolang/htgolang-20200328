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

// 定义SmsSender结构体并定义Send方法和SendAll方法
type SMSSender struct {
	api string
	id  string
	key string
}

func (s SMSSender) Send(to, msg string) error {
	fmt.Println("发送短信给: %s, 内容: %s\n", to, msg)
	return nil
}

func (s SMSSender) SendAll(to []string, msg string) error {
	fmt.Println("发送短信给: %#v, 内容: %s\n", to, msg)
	return nil
}

// 定义函数参数时接口类型变量
func PrintConfig(sender Sender) {
	// 需要将sender转换为结构体对象
	// 断言(转换后的一个对象, 转换失败)
	if obj, ok := sender.(EmailSender); ok {
		fmt.Printf("smtp服务器地址: %s\n", obj.addr)
	} else if obj, ok := sender.(SMSSender); ok {
		fmt.Printf("(值)api服务器地址: %s\n", obj.api)
	} else if obj, ok := sender.(*SMSSender); ok {
		fmt.Printf("(指针)api服务器地址: %s\n", obj.api)
	}
}

func PrintConfigV2(sender Sender) {
	// 需要将sender转换为结构体对象
	// 类型查询
	// .(type)只能用在switch
	switch v := sender.(type) {
	case EmailSender:
		fmt.Printf("v2: smtp服务器地址: %s\n", v.addr)
	case SMSSender:
		fmt.Printf("v2: (值)api服务器地址: %s\n", v.api)
	case *SMSSender:
		fmt.Printf("v2: (指针)api服务器地址: %s\n", v.api)
	default:
		fmt.Println("类型错误")
	}
}

func main() {
	emailSender := &EmailSender{"smtp.qq.com", 456, "kk", "kk123"}
	// PrintConfig(emailSender)
	PrintConfigV2(emailSender)

	smsSender := &SMSSender{"sms.tenant.com", "123123123", "fagfdsa"}
	// PrintConfig(smsSender)
	PrintConfigV2(smsSender)

}
