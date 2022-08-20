package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"sonui.cn/meows-list-server/pkg/utils"
)

const (
	NotActive = iota
	Active
	Baned
)
const (
	UserLevelNormal = iota
	UserLevelVIP
	UserLevelAdmin
)

type User struct {
	gorm.Model
	Name     string
	Level    int
	Email    string
	Password string
	Phone    string
	Status   int
}

func GetUserByEmail(email string) (*User, error) {
	user := User{}
	if err := DB.Where("email = ?", email).Limit(1).Find(&user).Error; err != nil {
		return nil, err
	}
	utils.Logger.Debug("user: %v", user)
	return &user, nil
}

func CheckUserPassword(user *User, password string) bool {
	hashedPassword := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	return err == nil
}

func SignUp(name, password, email, phone string) error {
	user := User{
		Name:     name,
		Password: password,
		Email:    email,
		Phone:    phone,
		Status:   NotActive,
		Level:    UserLevelNormal,
	}
	pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(pw)
	return DB.Create(&user).Error
}
