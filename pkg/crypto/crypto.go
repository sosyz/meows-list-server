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

func EncPassword(data []byte) ([]byte, error) {
	ret, err := bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func CheckPassword(data, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, data) == nil
}
