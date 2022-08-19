package main

import (
	"github.com/gin-gonic/gin"
	"sonui.cn/meows-list-server/pkg/utils"
	"sonui.cn/meows-list-server/routers"
)

func main() {
	// 定义日志输出

	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = utils.Logger.Write()
	api := routers.InitRouter()

	utils.Logger.Info("start meows-list-server listen on %s", utils.Config.RunConfig.Host+":"+utils.Config.RunConfig.Port)
	if err := api.Run(utils.Config.RunConfig.Host + ":" + utils.Config.RunConfig.Port); err != nil {
		utils.Logger.Panic("start meows-list-server error: %v", err)
	}
}
