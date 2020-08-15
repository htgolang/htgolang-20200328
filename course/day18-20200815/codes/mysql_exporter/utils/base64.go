package utils

import (
	"encoding/base64"
)

func Base64Encode(text string) string {
	if plaintext, err := base64.StdEncoding.DecodeString(text); err == nil {
		return string(plaintext)
	}
	return ""
}
