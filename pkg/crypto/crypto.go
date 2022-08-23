package crypto

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func RandString(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	token := ""
	for _, v := range b {
		// 转为可视ascii字符
		token += fmt.Sprintf("%x", v)
	}
	return token, nil
}

func EncPassword(password string) (string, error) {
	ret, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func CheckPassword(password, hashPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)) == nil
}
