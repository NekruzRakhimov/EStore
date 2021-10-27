package utils

import (
	"crypto/sha1"
	"fmt"
)

const (
	salt = "hjqrhjqw124617ajfhajs"
)

// Создание хеша sha1
func GenerateSha1(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
