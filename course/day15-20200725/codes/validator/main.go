package main

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

func main() {
	valid := &validation.Validation{}
	text := ""
	// 验证
	valid.Required(text, "required.required.required").Key("required").Message("输入内容不能为空")

	valid.Alpha("123abc.", "alpha.alpha.alpha").Message("输入内容只能是大小写英文字母")
	valid.Tel("152xxxxxx.", "tel.tel.tel").Message("输入联系方式不正确")
	valid.MaxSize("152xxxxxx.", 3, "tel.tel.tel").Message("字符串长度超过长度3")

	// 获取验证结果
	fmt.Println(valid.HasErrors())
	if valid.HasErrors() {
		fmt.Println(valid.Errors)
		fmt.Println(valid.ErrorsMap)
	}

}
