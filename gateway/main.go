package main

import (
	"moonChat/gateway/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	handler.UserCenterConnPool = handler.InitUserCenterGrpcConnPool()
	router := gin.Default()
	router.POST("/userCenter/userInfo/GetUser", handler.GetUserHandler)

	if err := router.Run(":8091"); err != nil {
		panic(err)
	}
}
