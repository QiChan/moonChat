package route

import (
	"moonChat/gateway/handler"

	"github.com/gin-gonic/gin"
)

func RouteUserCenter(router *gin.Engine) {
	router.POST("/userCenter/userInfo/GetUser", handler.GetUserHandler)
}
