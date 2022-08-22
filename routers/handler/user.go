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

func Signup(c *gin.Context) {
	SignupParams := controller.SignupParams{}
	if err := c.ShouldBind(&SignupParams); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}
	res := controller.UserSignup(c, &SignupParams)

	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Status(http.StatusOK)
	_, _ = c.Writer.WriteString(res)
}

func Info(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Quit(c *gin.Context) {

}
