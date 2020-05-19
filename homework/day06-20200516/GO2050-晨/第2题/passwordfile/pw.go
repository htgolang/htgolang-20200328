package passwordfile

import (
	"crypto/md5"
	"fmt"
	"math/rand"
)
func gen_RandString()  string{
	strs := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTXYZ1234567890"
	bytes := []byte{}
	for i:=0;i<5;i++{
		bytes = append(bytes,strs[rand.Intn(len(strs))])
	}
	return string(bytes)  //salt
}
func md5Str(pass,salt string) string {
	pass_salt := pass + salt
	return fmt.Sprintf("%x",md5.Sum([]byte(pass_salt)))
}
func Check_password(pw,real_pw string)  bool{
	salt := gen_RandString()
	md5pass := md5Str(real_pw,salt)
	input_md5 := md5Str(pw,salt)
	if md5pass == input_md5 {
		return true
	}
	return false
}
func Check_password_nosalt(pw,real_pw string) bool {
	if pw == real_pw{
		return true
	}
	return false
}


