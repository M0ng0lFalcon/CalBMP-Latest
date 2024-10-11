package Router

import (
	"calbmp-back/controller/HistoryController"
	"github.com/gin-gonic/gin"
)

func GetHistoryRouter(v1 *gin.RouterGroup) *gin.RouterGroup {
	HistoryRouter := v1.Group("/history")
	{
		HistoryRouter.POST("/addHistory", HistoryController.AddHistory)
		HistoryRouter.GET("/checkHistory", HistoryController.CheckHistory)
	}

	return HistoryRouter
}
