package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func randString(n int) (string, string) {
	rt1 := make([]byte, 0, n)
	rt2 := make([]byte, n, n)
	// 定义取值范围切片
	// chars := []byte{
	// 	'a', 'b', 'c', 'd', 'e', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	// 	'!', '@', '#',
	// }

	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// 循环n次 每次生成随机数(切片范围内)，获取对应的字符
	for i := 0; i < n; i++ {
		rt1 = append(rt1, chars[rand.Intn(len(chars))])
		rt2[i] = chars[rand.Intn(len(chars))]
	}

	return string(rt1), string(rt2)

}

func md5String(text string, salt string) string {
	//salt + ":" + text
	bytes := []byte(salt)
	bytes = append(bytes, ':')
	bytes = append(bytes, []byte(text)...)

	return fmt.Sprintf("%x", md5.Sum(bytes))
}

func main() {
	// hash算法 => 签名 (不可逆)
	// MD5, sha1, sha256, sha512

	fmt.Printf("%x\n", md5.Sum([]byte("我是kk")))

	hasher := md5.New()
	hasher.Write([]byte("我是"))
	hasher.Write([]byte("kk"))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))

	// 加盐 + md5
	salt1, salt2 := randString(6)
	fmt.Println(salt1, salt2)
	fmt.Println(md5String("我是kk", salt1))

	fmt.Printf("%x\n", sha1.Sum([]byte("我是KK")))
	fmt.Printf("%x\n", sha256.Sum256([]byte("我是KK")))
	fmt.Printf("%x\n", sha512.Sum512([]byte("我是KK")))

	sha256Hasher := sha256.New()
	sha256Hasher.Write([]byte("我是"))
	sha256Hasher.Write([]byte("KK"))
	fmt.Println(hex.EncodeToString(sha256Hasher.Sum(nil)))

	// bcrypt
}

//b4cdaf444f7c87b746f11db2a71d28f2
