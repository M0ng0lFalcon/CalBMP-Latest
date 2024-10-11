package Router

import (
	"calbmp-back/controller/ResultController"
	"github.com/gin-gonic/gin"
)

func GetResultRouter(v1 *gin.RouterGroup) *gin.RouterGroup {
	ResultRouter := v1.Group("/result")
	{
		ResultRouter.POST("/getBasicByName", ResultController.GetDataByName)       // visualization data
		ResultRouter.GET("/getComparisonData", ResultController.GetComparisonData) // comparison visualization data
		ResultRouter.POST("/get_text_result", ResultController.GetTextResult)
		ResultRouter.GET("/get_best_bmp", ResultController.GetBestBMP)
		ResultRouter.GET("/get_input_files", ResultController.GetInputFiles)
		ResultRouter.POST("/zip_input_file", ResultController.ZipInputFile)
		ResultRouter.GET("/download_input_file", ResultController.DownloadInputFile)
	}
	return ResultRouter
}
