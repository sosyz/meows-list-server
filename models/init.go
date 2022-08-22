package models

import (
	"fmt"
	"sonui.cn/meows-list-server/pkg/conf"
	"sonui.cn/meows-list-server/pkg/logger"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// Init 初始化连接
func Init() {
	var (
		db  *gorm.DB
		err error
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Database.Host,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Database,
		conf.Database.Port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "meows_",
		},
	})

	if err != nil {
		logger.Panic("DateBase.init err:", err)
	}

	//设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		logger.Panic("DateBase.init err:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	// 更新结构
	err = db.AutoMigrate(&User{})
	if err != nil {
		logger.Panic("DateBase.init err:", err)
	}
	DB = db
}
