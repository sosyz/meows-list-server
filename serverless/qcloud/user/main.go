package main

import (
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"sonui.cn/meows-list-server/controller"
)

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(controller.UserLogin)
}
