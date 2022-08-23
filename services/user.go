package services

import (
	"errors"
	"sonui.cn/meows-list-server/models"
	"sonui.cn/meows-list-server/pkg/cache"
	"sonui.cn/meows-list-server/pkg/crypto"
)

func UserLogin(email, password string) (string, error) {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("邮箱或密码错误")
	}
	if !crypto.CheckPassword(password, user.Password) {
		return "", errors.New("邮箱或密码错误")
	}

	token, err := crypto.RandString(64)
	if err != nil {
		return "", errors.New("登录失败")
	}
	if err := cache.Put(token, user, 0); err != nil {
		return "", errors.New("登录失败")
	}
	return token, nil
}

func GetUserByToken(token string) *models.User {
	user, err := cache.Get[models.User](token)
	if err != nil {
		return nil
	}
	return user
}

func UserRegister(name, password, email, phone string) error {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return errors.New("服务器错误")
	}
	if user.ID > 0 {
		return errors.New("邮箱已被注册")
	}
	if err := models.CreateUser(name, password, email, phone); err != nil {
		return err
	}
	return nil
}

// UserUpdate 更新用户信息
// 无需更新的信息传递空文本
func UserUpdate(token, name, email, phone, oldPassword, password string) error {
	if user, err := cache.Get[models.User](token); err != nil {
		return errors.New("未登录")
	} else if oldPassword != "" && !crypto.CheckPassword(oldPassword, user.Password) {
		return errors.New("旧密码错误")
	} else if err := models.UpdateUser(user.ID, name, email, phone, password); err != nil {
		return errors.New("更新失败")
	} else {
		user, _ = models.GetUserByID(user.ID)
		if err := cache.Del(token); err != nil {
			return errors.New("服务器错误")
		}
		if err := cache.Put(token, user, 0); err != nil {
			return errors.New("服务器错误")
		}
		return nil
	}
}

func RemoveToken(token string) error {
	return cache.Del(token)
}
