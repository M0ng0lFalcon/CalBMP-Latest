package BasicDataController

import (
	"calbmp-back/Repository/ChorizonRepository"
	"calbmp-back/Res"
	"calbmp-back/service/BasicDataService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAllSoil(ctx *gin.Context) {
	SoilList := BasicDataService.GetAllSoilFun()
	Res.Success(ctx, gin.H{"soil_list": SoilList}, "[*] Get all soil")
}

func GetZipCodeBySoil(ctx *gin.Context) {
	mukey := ctx.Query("mukey")
	cokey := ctx.Query("cokey")
	zipCode := BasicDataService.GetZipCodeBySoil(mukey, cokey)
	Res.Success(ctx, gin.H{"zip_code": zipCode}, "[*] Get zipcode by soil")
}

func GetCountyNames(ctx *gin.Context) {
	CountyNames := BasicDataService.GetCountyNames()

	Res.Success(
		ctx,
		gin.H{"counties": CountyNames},
		"[*] Get County Names, Count:"+fmt.Sprintf("%d", len(CountyNames)),
	)
}

func GetZipcode(ctx *gin.Context) {
	CountyName := ctx.Query("CountyName")

	ZipCodes := BasicDataService.GetZipCodes(CountyName)

	Res.Success(
		ctx,
		gin.H{"zipCode": ZipCodes},
		"[*] Get Zip Codes, Count:"+fmt.Sprintf("%d", len(ZipCodes)),
	)
}

func GetCompnamesMukeyCokey(ctx *gin.Context) {
	ZipCode := ctx.Query("zip_code")
	CompnameMukeyCokeyList := BasicDataService.GetCompnamesMukeyCokey(ZipCode)

	Res.Success(ctx, gin.H{"compname_mukey_cokey": CompnameMukeyCokeyList}, "")
}

func GetStationData(ctx *gin.Context) {
	ZipCode := ctx.Query("ZipCode")
	StationData := BasicDataService.GetStation(ZipCode)
	Res.Success(
		ctx,
		gin.H{"Station": StationData},
		"[*] Get Station Data",
	)
}

func GetMuname(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	MunameLi := ChorizonRepository.FindMuname(keyword)
	Res.Success(
		ctx,
		gin.H{"muname": MunameLi},
		"[*] Get muname Data",
	)
}

func GetCompname(ctx *gin.Context) {
	muname := ctx.Query("muname")
	CompnameLi := ChorizonRepository.FindCompname(muname)
	Res.Success(
		ctx,
		gin.H{"compname": CompnameLi},
		"[*] Get compname Data",
	)
}

func GetMukeyCokeyByName(ctx *gin.Context) {
	muname := ctx.Query("muname")
	compname := ctx.Query("compname")
	mukey, cokey := ChorizonRepository.FindByMunameCompname(muname, compname)
	Res.Success(
		ctx,
		gin.H{
			"mukey": mukey,
			"cokey": cokey,
		},
		"[*] Get key Data",
	)
}

// ----------------------------------------

func GetCropNames(ctx *gin.Context) {
	CropNames := BasicDataService.GetCropNames()

	Res.Success(
		ctx,
		gin.H{"CropNames": CropNames},
		"[*] Get CropNames, count:"+fmt.Sprintf("%d", len(CropNames)),
	)
}

func GetPesticide(ctx *gin.Context) {
	Pesticides := BasicDataService.GetPesticides()

	Res.Success(
		ctx,
		gin.H{"Pesticides": Pesticides},
		"[*] Get Pesticides, count:"+fmt.Sprintf("%d", len(Pesticides)),
	)
}
