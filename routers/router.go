package routers

import (
	"github.com/gin-gonic/gin"
	"sonui.cn/meows-list-server/routers/handler"
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
		user.POST("login", handler.UserLogin)
		// 注册
		user.POST("signup", handler.Signup)
		// 获取账号信息
		user.GET("/:id/info", handler.Info)
		// 设置账号信息
		user.PUT("set", handler.Update)
		// 退出登录
		user.POST("quit", handler.Quit)
	}
	return r
}
