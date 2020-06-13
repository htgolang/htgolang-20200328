package main

import (
	"fmt"
	"regexp"
)

func main() {

	// . 任意
	// \d 数字 \D非数字
	// \w 数字，大小写英文字母 _
	// \S 非空白字符 \s非空白字符
	// \d 数字 0,1,2,3,4,5,6,7,8,9 0|1|2|3|4|5|6|7|8|9 [0123456789] [0-9]
	// [a-z] /a, -, z/ [a\-z]
	// [^a-z] 取反
	// ? 0个或1个
	// + 至少1个
	// * 任意多个
	// {n,m} 字符串数量>=n, <=m
	var pattern string = "^132" // 已132开头
	fmt.Println(regexp.MatchString(pattern, "132xxxx"))
	fmt.Println(regexp.MatchString(pattern, "122xxxx"))

	// 以132开头,都是数字,长度为11位 [0-9]
	pattern = "^132\\d{8}$"
	fmt.Println(regexp.MatchString(pattern, "132xxxx"))
	fmt.Println(regexp.MatchString(pattern, "132123123"))
	fmt.Println(regexp.MatchString(pattern, "13212312323"))
	fmt.Println(regexp.MatchString(pattern, "13212312323x"))

	// 132 158
	// 1[35][28] 132 138 152 158
	// 分组 ()
	// ^(132)|(158)[0-9]{8}$
	// 邮箱格式
	// xxxx@xxx.xx
	// xxxx(@之前) 数字,大小写英文字母长度1-32个字符
	// xxx(.之前) 小写英文字母组成 长度1-12字符
	// xx(.之后) edu
	// [a-zA-Z0-9]{1,32}@[a-z]{1,12}.edu

	fmt.Println("----email----")
	//[.]
	pattern = "[a-zA-Z0-9]{1,32}@[a-z]{1,12}\\.edu"

	pattern = "^[a-zA-Z0-9]{1,32}@[a-z]{1,12}[.]edu$"

	fmt.Println(regexp.MatchString(pattern, "a@b.edu"))    // 可以匹配
	fmt.Println(regexp.MatchString(pattern, "a@1.edu"))    // 不匹配
	fmt.Println(regexp.MatchString(pattern, "?@b.edu"))    // 不匹配
	fmt.Println(regexp.MatchString(pattern, "a@bxedu"))    // 不匹配
	fmt.Println(regexp.MatchString(pattern, "我是a@b.edux")) // 不匹配

	pattern = regexp.QuoteMeta("^ab")

	fmt.Println(regexp.MatchString(pattern, "ab"))
	fmt.Println(pattern)
	fmt.Println(regexp.MatchString(pattern, "^ab"))
}
