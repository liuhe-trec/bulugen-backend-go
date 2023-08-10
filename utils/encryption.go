package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// 密码比较 hash是密文和password明文比较
func CompareHashAndPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
