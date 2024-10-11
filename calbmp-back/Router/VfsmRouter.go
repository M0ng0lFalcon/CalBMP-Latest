package Router

import (
	"calbmp-back/controller/VFSMController"
	"calbmp-back/security"
	"github.com/gin-gonic/gin"
)

func GetVfsmRouter(v1 *gin.RouterGroup) *gin.RouterGroup {
	VfsmRouter := v1.Group("/vfsm")
	// post
	{
		VfsmRouter.POST("/run", security.AuthMiddleware(), VFSMController.GenerateInputFile)
		VfsmRouter.POST("/storm", security.AuthMiddleware(), VFSMController.Storm)
		VfsmRouter.POST("/specify_length", security.AuthMiddleware(), VFSMController.SpecifyLengthController)
	}
	// get
	{
		VfsmRouter.GET("/get_vegetation", security.AuthMiddleware(), VFSMController.GetVegetation)
		VfsmRouter.GET("/get_vegetation_model", security.AuthMiddleware(), VFSMController.GetVegetationByName)
		VfsmRouter.GET("/get_basic_result", security.AuthMiddleware(), VFSMController.GetBasicResult)
		VfsmRouter.GET("/get_overall", security.AuthMiddleware(), VFSMController.GetPesticideReductionEff)
		VfsmRouter.GET("/get_res_visualization", security.AuthMiddleware(), VFSMController.Result2Visualization)
		VfsmRouter.GET("/get_progress", security.AuthMiddleware(), VFSMController.GetProgress)
		VfsmRouter.GET("/get_comparison", security.AuthMiddleware(), VFSMController.ComparisonData)
		VfsmRouter.GET("/get_inputs", security.AuthMiddleware(), VFSMController.GetInputs)
		VfsmRouter.GET("/get_outputs", security.AuthMiddleware(), VFSMController.GetOutputs)
	}
	return VfsmRouter
}
