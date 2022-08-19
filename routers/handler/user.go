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
	if res, err := controller.UserLogin(c, &LoginParams); err != nil {
		c.JSON(http.StatusOK, utils.ErrorResponse(err.Error()))
	} else {
		c.JSON(http.StatusOK, utils.SuccessResponse("", res))
	}
}

func Signup(c *gin.Context) {

}

func Info(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Quit(c *gin.Context) {

}
