package models

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sonui.cn/meows-list-server/pkg/utils"
)

var DB *gorm.DB

// Init 初始化连接
func init() {
	var (
		db  *gorm.DB
		err error
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		utils.Config.DataBaseConfig.Host,
		utils.Config.DataBaseConfig.User,
		utils.Config.DataBaseConfig.Password,
		utils.Config.DataBaseConfig.Database,
		utils.Config.DataBaseConfig.Port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "meows_",
		},
	})

	if err != nil {
		utils.Logger.Panic("DateBase.init err:", err)
	}

	//设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		utils.Logger.Panic("DateBase.init err:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	// 更新结构
	err = db.AutoMigrate(&User{})
	if err != nil {
		utils.Logger.Panic("DateBase.init err:", err)
	}
	DB = db
}
