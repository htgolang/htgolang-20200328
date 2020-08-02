package utils

import "golang.org/x/crypto/bcrypt"

// GeneratePassword 生成bcrypt hash
func GeneratePassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		panic(err.Error())
	}
	return string(hash)
}

// CheckPassword 检查密码正确性
func CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
