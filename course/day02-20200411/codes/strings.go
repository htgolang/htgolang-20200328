package main

import (
	"fmt"
	"strings"
)

func main() {
	// 比较
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))

	// 包含
	fmt.Println(strings.Contains("我是kk", "kk"))
	fmt.Println(strings.Contains("我是kk", "kk1"))

	fmt.Println(strings.ContainsAny("我是kk", "kk1"))
	fmt.Println(strings.ContainsAny("我是kk", "123"))

	fmt.Println(strings.ContainsRune("我是kk", '我'))
	fmt.Println(strings.ContainsRune("我是kk", 'a'))

	// 计算次数
	fmt.Println(strings.Count("我是kk", "kk"))
	fmt.Println(strings.Count("我是kk", "k"))
	fmt.Println(strings.Count("我是kk", "a"))

	// 比较
	fmt.Println(strings.EqualFold("abc", "ABC"))
	fmt.Println(strings.EqualFold("abc", "abc"))
	fmt.Println(strings.EqualFold("abc", "xyz"))

	// 空白符
	// 空格 tab 回车 换行 换页 ...
	fmt.Printf("%#v\n", strings.Fields("aafds b\tc\nd\re\ff"))

	// 开头*结尾
	fmt.Println(strings.HasPrefix("abc", "ab"))
	fmt.Println(strings.HasPrefix("abc", "bc"))

	fmt.Println(strings.HasSuffix("abc", "bc"))
	fmt.Println(strings.HasSuffix("abc", "ab"))

	// 字符串出现位置
	fmt.Println(strings.Index("abcdefdef", "def"))
	fmt.Println(strings.Index("abcdefdef", "xxx"))
	fmt.Println(strings.LastIndex("abcdefdef", "def"))
	fmt.Println(strings.LastIndex("abcdefdef", "xxx"))

	// 连接 分割
	fmt.Println(strings.Join([]string{"ab", "cd", "ef"}, "-"))

	fmt.Printf("%#v\n", strings.Split("ab-ab-ab", "-"))
	fmt.Printf("%#v\n", strings.SplitN("ab-ab-ab", "-", 2))

	// 重复

	fmt.Println(strings.Repeat("*", 5))

	// 替换
	fmt.Println(strings.Replace("xyzxyzxxxxyz", "yz", "mn", -1))
	fmt.Println(strings.Replace("xyzxyzxxxxyz", "yz", "mn", 1))
	fmt.Println(strings.ReplaceAll("xyzxyzxxxxyz", "yz", "mn"))

	// 首字母大写
	fmt.Println(strings.Title("my name is kk"))

	fmt.Println(strings.ToLower("abcABC"))
	fmt.Println(strings.ToUpper("abcABC"))

	// trim
	fmt.Println(strings.Trim("abcdefabc", "bc"))
	fmt.Println(strings.Trim("abcdefabc", "abc"))
	fmt.Println(strings.TrimSpace(" \n\f\tabcdefabc\t"))
	fmt.Println(strings.TrimLeft("cabcdefabca", "abc")) // 左边字符出现在子字符串中则替换
	fmt.Println(strings.TrimRight("cabcdefabca", "abc"))

	fmt.Println(strings.TrimPrefix("abccabcdefabca", "abc")) // 字符串当成一个整体替换
	fmt.Println(strings.TrimSuffix("cabcdefabca", "abc"))
}
