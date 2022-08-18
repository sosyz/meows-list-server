package services

import (
	"github.com/gin-gonic/gin"
)

type LoginParams struct {
	Name   string `json:"name"`
	Pass   string `json:"pass"`
	Verify string `json:"verify"`
}

func Login(c *gin.Context) {

}

func Signup(c *gin.Context) {

}

func Info(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Quit(c *gin.Context) {

}
