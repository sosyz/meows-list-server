package routers

import (
	"github.com/gin-gonic/gin"
	"sonui.cn/meows-list-server/services"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(Cors())
	//服务器状态检测
	r.GET("/ping", func(context *gin.Context) {
		context.String(200, "!pong")
	})

	// v1接口
	v1 := r.Group("/v1/api")

	user := v1.Group("/user")
	{
		// 登录
		user.POST("login", services.Login)
		// 注册
		user.POST("signup", services.Signup)
		// 获取账号信息
		user.GET("/:id/info", services.Info)
		// 设置账号信息
		user.PUT("set", services.Update)
		// 退出登录
		user.POST("quit", services.Quit)
	}
	return r
}
