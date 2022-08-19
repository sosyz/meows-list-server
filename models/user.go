package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	Active = iota
	NotActive
	Baned
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);"`
	Level    int    `gorm:"type:int(11);"`
	Email    string `gorm:"type:varchar(100);"`
	Password string `gorm:"type:varchar(100);"`
	Phone    string `gorm:"type:varchar(100);"`
	Status   int    `gorm:"type:int(11);"`
}

func GetUserByEmail(email string) (User, error) {
	var user User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CheckUserPassword(user User, password string) bool {
	hashedPassword := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	return err == nil
}

func SignUp(user User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(password)
	return DB.Create(&user).Error
}
