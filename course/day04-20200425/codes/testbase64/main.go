package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// base64
	// 通常说的base64 0-9a-zA-Z+/ 64
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("我发动机开始了拉拉啊解开了是kk")))
	txt, _ := base64.StdEncoding.DecodeString("5oiR5piva2s=")
	fmt.Println(string(txt))
	// 在URL中+_特殊字符, base64url (+(-)/(_)替换)
	fmt.Println(base64.URLEncoding.EncodeToString([]byte("我发动机开始了拉拉啊解开了是kk")))
	txt, _ = base64.URLEncoding.DecodeString("5oiR5Y-R5Yqo5py65byA5aeL5LqG6Kej5byA5LqG5piva2s=")
	fmt.Println(string(txt))
	// 非对齐 3的整数倍 =补齐
	// 标准
	fmt.Println(base64.RawStdEncoding.EncodeToString([]byte("我发动机开始了拉拉啊解开了是kk")))
	// url
	fmt.Println(base64.RawURLEncoding.EncodeToString([]byte("我发动机开始了拉拉啊解开了是kk")))

}
