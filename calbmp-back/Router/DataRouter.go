package Router

import (
	"calbmp-back/controller/BasicDataController"
	"github.com/gin-gonic/gin"
)

func GetDataRouter(v1 *gin.RouterGroup) *gin.RouterGroup {
	DataRouter := v1.Group("/data")
	{
		// step 1
		DataRouter.GET("/get_all_soil", BasicDataController.GetAllSoil)
		DataRouter.GET("/get_zipcode_by_soil", BasicDataController.GetZipCodeBySoil)
		DataRouter.GET("/getCounties", BasicDataController.GetCountyNames)
		DataRouter.GET("/getZipcode", BasicDataController.GetZipcode)
		DataRouter.GET("/getCompnameMukeyCokey", BasicDataController.GetCompnamesMukeyCokey)
		DataRouter.GET("/getStation", BasicDataController.GetStationData)

		// step 2
		DataRouter.GET("/getCropName", BasicDataController.GetCropNames)
		DataRouter.GET("/getPesticide", BasicDataController.GetPesticide)

		// new step
		DataRouter.GET("/get_muname", BasicDataController.GetMuname)
		DataRouter.GET("/get_compname", BasicDataController.GetCompname)
		DataRouter.GET("/get_mukey_cokey", BasicDataController.GetMukeyCokeyByName)

	}

	return DataRouter
}
