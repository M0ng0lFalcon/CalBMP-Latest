package CommodityDatabaseCropRepository

import (
	"calbmp-back/Database"
	"calbmp-back/Repository/ChorizonRepository"
	"calbmp-back/model"
	"fmt"
	"strconv"
	"strings"
)

func GetCNValueByHSGAndSiteName(HSG string, siteName string) string {
	db := Database.GetDB()
	var CNValue string
	var CommodityDatabaseCrop model.CommodityDatabaseCrop

	HSGString := "hsg_" + HSG
	db.Select(HSGString).Where("site_name=?", siteName).First(&CommodityDatabaseCrop).Scan(&CNValue)

	return CNValue
}

func GetBareSoil(HSG string) string {
	temp := []string{"77", "86", "91", "94"}
	switch HSG {
	case "A":
		return temp[0]
	case "B":
		return temp[1]
	case "C":
		return temp[2]
	case "D":
		return temp[3]
	case "a":
		return temp[0]
	case "b":
		return temp[1]
	case "c":
		return temp[2]
	case "d":
		return temp[3]
	}
	return ""
}

func GetAllCrops() []string {
	db := Database.GetDB()
	CropNames := make([]string, 0)
	var Crops model.CommodityDatabaseCrop

	// site name equal crop name
	param := "site_name"

	db.Distinct(param).Where(param + " IS NOT NULL").Order(param).Find(&Crops).Scan(&CropNames)

	return CropNames
}

func FindCommodityDatabaseCropBySiteName(SiteName string) model.CommodityDatabaseCrop {
	var CommodityDatabaseCrop model.CommodityDatabaseCrop
	db := Database.GetDB()

	db.Where("site_name = ?", SiteName).First(&CommodityDatabaseCrop)

	return CommodityDatabaseCrop
}

func GetCNValueAndBareValue(
	SiteName string,
	cokey string,
	mukey string,
) (string, string) {
	//HSG := ComponentRepository.GetHSGType(compName, cokey, mukey)
	HSG := ChorizonRepository.FindHSGTypeByCompNameAndMukeyAndCokey(mukey, cokey)
	HSG = strings.ToLower(HSG)

	OrgCNValueStr := GetCNValueByHSGAndSiteName(HSG, SiteName)
	OrgCnValue, _ := strconv.Atoi(OrgCNValueStr)
	CNValue := fmt.Sprintf("%d", OrgCnValue+3)

	BareSoilValue := GetBareSoil(HSG)

	return CNValue, BareSoilValue
}

func GetUSLECmnBySiteName(SiteName string) float64 {
	db := Database.GetDB()
	var CommodityDatabaseCrop model.CommodityDatabaseCrop

	db.Where("site_name=?", SiteName).First(&CommodityDatabaseCrop)

	return CommodityDatabaseCrop.USLECmn
}
