package services

import (
	"errors"
	"sonui.cn/meows-list-server/models"
)

func UserLogin(email, password string) error {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return errors.New("邮箱或密码错误")
	}
	if !models.CheckUserPassword(user, password) {
		return errors.New("邮箱或密码错误")
	}
	return nil
}

func UserSignup(name, password, email, phone string) error {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return err
	}
	if user.ID > 0 {
		return errors.New("邮箱已被注册")
	}
	if err := models.SignUp(name, password, email, phone); err != nil {
		return err
	}
	return nil
}
