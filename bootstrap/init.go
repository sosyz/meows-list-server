package bootstrap

import (
	"sonui.cn/meows-list-server/pkg/cache"
	"sonui.cn/meows-list-server/pkg/conf"
	"sonui.cn/meows-list-server/pkg/logger"
)

func Init(path string) {
	conf.Init(path)
	cache.Init(conf.Redis.Host, conf.Redis.Port, conf.Redis.Password, conf.Redis.Database)
	logger.Init(conf.Run.LogLevel)
}
