package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sonui.cn/meows-list-server/controller"
	"sonui.cn/meows-list-server/pkg/utils"
)

func UserLogin(c *gin.Context) {
	LoginParams := controller.LoginParams{}
	if err := c.ShouldBind(&LoginParams); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}
	res := controller.UserLogin(c, &LoginParams)

	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Status(http.StatusOK)
	_, _ = c.Writer.WriteString(res)

}

func Register(c *gin.Context) {
	SignupParams := controller.RegisterParams{}
	if err := c.ShouldBind(&SignupParams); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}
	res := controller.UserRegister(c, &SignupParams)

	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Status(http.StatusOK)
	_, _ = c.Writer.WriteString(res)
}

func Info(c *gin.Context) {
	// 获取请求头
	token := c.GetHeader("token")
	c.Set("token", token)

	res := controller.UserInfo(c)
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Status(http.StatusOK)
	_, _ = c.Writer.WriteString(res)
}

func Update(c *gin.Context) {
	token := c.GetHeader("token")
	c.Set("token", token)

	RegisterParams := controller.UpdateParams{}
	if err := c.ShouldBind(&RegisterParams); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	res := controller.UserUpdate(c, &RegisterParams)
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Status(http.StatusOK)
	_, _ = c.Writer.WriteString(res)
}

func Logout(c *gin.Context) {
	token := c.GetHeader("token")
	c.Set("token", token)

	res := controller.UserLogout(c)
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Status(http.StatusOK)
	_, _ = c.Writer.WriteString(res)
}
