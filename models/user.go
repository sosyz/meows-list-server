package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);"`
	Salt string `gorm:"type:char(64);"`
}
