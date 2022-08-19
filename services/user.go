package services

import (
	"errors"
	"sonui.cn/meows-list-server/models"
)

func UserLogin(email, password string) (string, error) {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("邮箱或密码错误")
	}
	if !models.CheckUserPassword(user, password) {
		return "", errors.New("邮箱或密码错误")
	}

	return "", nil
}
