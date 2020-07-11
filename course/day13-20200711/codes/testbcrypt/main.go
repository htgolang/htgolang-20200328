package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "123abc"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)

	fmt.Println(string(hash), err)
	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte("123abcd")))
	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte("123abc")))
}
