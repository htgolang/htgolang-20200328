package hashValue


import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

//生成随机字符串的函数，该函数生成的随机字符串就作为salt
func RandString(n int) string {
	//声明并初始化一个字节类型空切片
	s1 := make([]byte,n)

	chars := "$%*abcdefghijklmnopqrstuvw%*xyzABCDEFGHIJKLMN%*OPQRSTUVWXYZ"
	for i:=0;i<n;i++ {
		s1 = append(s1,chars[rand.Intn(len(chars))])
	}
	return string(s1)
}

func Md5String(pass, salt string) string {
	bytes := []byte(pass)
	bytes = append(bytes,[]byte(salt)...)
	md5Value := md5.Sum(bytes)
	return fmt.Sprintf("%x",md5Value)
}

