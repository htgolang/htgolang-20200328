package utils

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// check phone
func IsPhone(phone string) bool {
	reg, err := regexp.Compile("^((1[3,5,8][0-9])|(14[5,7])|(17[0,6,7,8]))\\d{8}$")
	if err != nil {
		log.Println(err)
		return false
	}
	return reg.MatchString(phone)
}

// generate salt
func generateSalt(saltlen int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	salt := make([]byte, saltlen)
	for i := 0; i < saltlen; i++ {
		salt[i] = chars[rand.Intn(len(chars))]
	}
	return string(salt)
}

// sha256 data
func sha256Text(text string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
}

// hash password
func HashPasswd(text, salt string) string {
	if salt == "" {
		salt = generateSalt(rand.Int()%5 + 5) // salt lenght random scope 5 - 10
	}
	return fmt.Sprintf("%s$$%s", salt, sha256Text(fmt.Sprintf("%s:%s", text, salt)))
}

// check valid password
func ValidPassword(text, hash string) bool {
	pos := strings.LastIndex(hash, "$$")
	if HashPasswd(text, hash[:pos]) == hash {
		return true
	}
	return false
}
