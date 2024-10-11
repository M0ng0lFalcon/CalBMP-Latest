package Router

import (
	"calbmp-back/controller/InputFileController"
	"calbmp-back/security"
	"github.com/gin-gonic/gin"
)

func GetInputRouter(v1 *gin.RouterGroup) *gin.RouterGroup {
	InputRouter := v1.Group("/input")
	{
		InputRouter.POST("/step1", security.AuthMiddleware(), InputFileController.UserInputStep1)
		InputRouter.POST("/step2", security.AuthMiddleware(), InputFileController.UserInputChemical)
		InputRouter.GET("/get_step_json", security.AuthMiddleware(), InputFileController.GetStepJson)
	}

	return InputRouter
}
