package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPasswd 生成加密密码
func HashPasswd(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		panic(err.Error())
	}
	return string(hash)
}

// ValidPassword 校验密码
func ValidPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
