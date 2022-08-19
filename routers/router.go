package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sonui.cn/meows-list-server/routers/handler"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s [Info] Gin %s %s \"%s\" %s \n",
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Header.Get("User-Agent"),
		)
	}))
	r.Use(gin.Recovery())
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
