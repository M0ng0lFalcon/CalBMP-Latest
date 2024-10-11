package GlobalValueUtil

import (
	"calbmp-back/Params/InputParams"
	"calbmp-back/Repository/CommodityDatabaseCropRepository"
	"calbmp-back/Repository/CommodityDatabaseUSLECRepository"
)

func GetGlobalVariables(rec InputParams.UserInputStepReceiver) GlobalVar {
	var gv GlobalVar
	// --- get uslec and crop data ---
	// get params
	crop := rec.Crop // crop = site name
	cokey := rec.CoKey
	mukey := rec.MuKey
	// get value
	CDCrop := CommodityDatabaseCropRepository.FindCommodityDatabaseCropBySiteName(crop)
	CropClass := CDCrop.Classes

	// these lists are from commodity database USLEC
	USLEC := CommodityDatabaseUSLECRepository.FindUSLECValuesByCropClass(CropClass)
	N := CommodityDatabaseUSLECRepository.FindNValueByClass(CropClass)

	// generate gv
	gv.USLEC = USLEC
	gv.N = N
	gv.CNValue, gv.BareSoil = CommodityDatabaseCropRepository.GetCNValueAndBareValue(crop, cokey, mukey)
	return gv
}
