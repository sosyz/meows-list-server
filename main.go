package main

import (
	"github.com/gin-gonic/gin"
	"sonui.cn/meows-list-server/bootstrap"
	"sonui.cn/meows-list-server/pkg/conf"
	"sonui.cn/meows-list-server/pkg/logger"
	"sonui.cn/meows-list-server/routers"
)

func main() {
	// 定义日志输出
	bootstrap.Init("./")

	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = logger.Write()
	api := routers.InitRouter()

	logger.Info("start meows-list-server listen on %s", conf.Run.Host+":"+conf.Run.Port)
	if err := api.Run(conf.Run.Host + ":" + conf.Run.Port); err != nil {
		logger.Panic("start meows-list-server error: %v", err)
	}
}
