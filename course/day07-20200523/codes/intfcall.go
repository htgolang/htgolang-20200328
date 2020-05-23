package main

import "fmt"

type Sender interface {
	Send(string, string) error
}

// 定义EmailSender结构体 并实现Sender接口
type EmailSender struct {
}

func (s EmailSender) Send(to, msg string) error {
	fmt.Println("发送邮件")
	return nil
}

// 定义SmsSender结构体 并实现Sender接口
type SmsSender struct {
}

func (s SmsSender) Send(to, msg string) error {
	fmt.Println("发送短信")
	return nil
}

/*

JsonEncoder
CsvEncoder
GobEncoder
XXXXEncoder

typ => Encoder

func save(encoder, []*Task) {
	// 打开文件
	// 延迟关闭
	// encode => 传入的 json csv gob
	写入文件
}
func save(typ, []*Task) {
	if typ == "json" {
		// 打开文件
		// 延迟关闭
		// 创建encoder对象
		// encode

	} else if typ == "csv" {

		// ...
	} else {

		//...

	}
}
*/

func main() {
	var sender Sender

	sender = DongdongSender{}

	sender.Send("", "") // 调用的EmailSender

	//把不同类型的变量赋值给了同一个变量

}
