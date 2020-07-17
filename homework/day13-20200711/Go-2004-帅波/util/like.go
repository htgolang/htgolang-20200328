package util

import (
	"strings"
)

func Like(method string) string {
	//去除空格
	method = strings.TrimSpace(method)
	//  /必须放在最前面因为如果替换了%之后会有/  那么就会出现问题
	method = strings.Replace(method,"/","//",-1)
	method = strings.Replace(method,"%","/%",-1)
	method = strings.Replace(method,"_","/_",-1)
	//sprintf用%对%解除特殊字符
	//return fmt.Sprintf("%%%s%%",method)
	//返回"%123%"  用于sql模糊匹配
	return "%" + method + "%"
}
