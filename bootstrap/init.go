package bootstrap

import (
	"sonui.cn/meows-list-server/models"
	"sonui.cn/meows-list-server/pkg/cache"
	"sonui.cn/meows-list-server/pkg/conf"
	"sonui.cn/meows-list-server/pkg/logger"
)

func Init(path string) {
	logger.Init("debug")
	conf.Init(path)
	logger.SetLevel(conf.Run.LogLevel)
	models.Init()
	cache.Init(conf.Redis.Host, conf.Redis.Port, conf.Redis.Password, conf.Redis.Database)
}
