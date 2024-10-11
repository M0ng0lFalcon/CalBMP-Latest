package Router

import (
	"calbmp-back/controller/UserController"
	"calbmp-back/security"
	"github.com/gin-gonic/gin"
)

func GetUserRouter(v1 *gin.RouterGroup) *gin.RouterGroup {
	UserRouter := v1.Group("/user")
	{
		UserRouter.POST("/register", UserController.Register)
		UserRouter.POST("/login", UserController.Login)
		UserRouter.GET("/info", security.AuthMiddleware(), UserController.Info)
		UserRouter.POST("/change_password", UserController.ChangePassword)
	}

	return UserRouter
}
