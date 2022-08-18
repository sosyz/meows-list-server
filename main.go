package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sonui.cn/meows-list-server/pkg/utils"
	"sonui.cn/meows-list-server/routers"
)

func main() {
	// 定义日志输出
	var logLevel logrus.Level

	logrus.SetLevel(logLevel)

	gin.SetMode(gin.ReleaseMode)
	api := routers.InitRouter()

	logrus.Info("start meows-list-server listen on %s...", utils.Config.RunConfig.Host+":"+utils.Config.RunConfig.Port)
	if err := api.Run(utils.Config.RunConfig.Host + ":" + utils.Config.RunConfig.Port); err != nil {
		logrus.Error("start meows-list-server error: %v", err)
	}
}
