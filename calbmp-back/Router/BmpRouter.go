package Router

import (
	"calbmp-back/controller/BMPController"
	"calbmp-back/security"
	"github.com/gin-gonic/gin"
)

func GetBmpRouter(v1 *gin.RouterGroup) *gin.RouterGroup {
	BmpRouter := v1.Group("/bmp")
	{
		BmpRouter.POST("/bmpScenario", security.AuthMiddleware(), BMPController.BmpScenario)
	}
	return BmpRouter
}
