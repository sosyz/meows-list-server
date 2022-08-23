package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	return &user, nil
}

func GetUserByID(id uint) (*User, error) {
	user := User{}
	if err := DB.Where("id = ?", id).Limit(1).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(name, password, email, phone string) error {
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

func UpdateUser(id uint, name, email, phone, password string) error {
	user := User{}
	if err := DB.Where("id = ?", id).Find(&user).Error; err != nil {
		return err
	}
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	if phone != "" {
		user.Phone = phone
	}
	if password != "" {
		pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(pw)
	}
	return DB.Save(&user).Error
}
