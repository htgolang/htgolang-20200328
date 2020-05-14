package passwd

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateSalt(saltLen int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	salt := make([]byte, saltLen)
	for i := 0; i < saltLen; i++ {
		salt[i] = chars[rand.Intn(len(chars))]
	}
	return string(salt)
}

// ComparePasswd 检查密码
func ComparePasswd(text, hash string) bool {
	pos := strings.LastIndex(hash, "$$")
	if HashPasswd(text, hash[:pos]) == hash {
		return true
	}
	return false
}

func sha256text(text string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
}

// HashPasswd 加密密码
func HashPasswd(text, salt string) string {
	if salt == "" {
		salt = generateSalt(6)
	}
	return fmt.Sprintf("%s$$%s", salt, sha256text(fmt.Sprintf("%s:%s", text, salt)))
}
