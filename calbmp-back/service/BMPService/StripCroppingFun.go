package BMPService

import (
	"calbmp-back/Params/BmpParams"
	"calbmp-back/Params/InputParams"
	"calbmp-back/Repository/ChorizonRepository"
	"calbmp-back/Repository/CommodityDatabaseCropRepository"
	"calbmp-back/Repository/CommodityDatabaseUSLECRepository"
	"calbmp-back/Repository/RainDistributionIregRepository"
	"calbmp-back/util/GlobalValueUtil"
	"calbmp-back/util/InputFileGenerateUtil"
	"fmt"
	"strconv"
)

func StripCropping(
	rec1 InputParams.UserInputStepReceiver,
	rec2 InputParams.UserInputStepReceiver,
	gv *GlobalValueUtil.GlobalVar,
	SC BmpParams.StripCroppingParams,
) (string, string, string) {
	Record3 := GetStripCroppingRecord3(rec1)
	Record5 := GetStripCroppingRecord5(rec2, gv, SC)
	Record9 := GetStripCroppingRecord9(rec2, gv)

	return Record3, Record5, Record9
}

func GetStripCroppingRecord3(rec InputParams.UserInputStepReceiver) string {
	res := ""

	// get params
	ZipCode := rec.ZipCode
	Mukey := rec.MuKey
	Cokey := rec.CoKey
	FieldSize := rec.FieldSize

	fieldSize, _ := strconv.ParseFloat(FieldSize, 64)
	fieldSize /= 2.471054

	// get orm obj
	Chorizon := ChorizonRepository.FindChorizonByMukeyAndCokeyAndHzdept(Mukey, Cokey, "0")
	RainDistributionIreg := RainDistributionIregRepository.FindRainDistributionIregByZipCode(ZipCode)

	// get basic values
	uslek, usles, uslep, IREG, slope, slopeLength := InputFileGenerateUtil.GetRecord3BasicValues(
		Chorizon,
		RainDistributionIreg,
		true,
		rec.KnowSlope,
		rec.Slope,
	)

	if rec.KnowSlope {
		slope = rec.Slope
	}

	// -------------- BMP part --------------
	if slope > 24 {
		uslep = 1.0
	} else if slope <= 24 && slope > 18 {
		uslep = 0.4
	} else if slope <= 18 && slope > 12 {
		uslep = 0.35
	} else if slope <= 12 && slope > 7 {
		uslep = 0.3
	} else if slope <= 7 && slope > 2 {
		uslep = 0.25
	} else {
		uslep = 0.3
	}
	// -------------- BMP part --------------

	// generate res
	temp := "%s, %f, %.2f, %f, %s, %f, %f"
	res = fmt.Sprintf(temp,
		uslek,
		usles,
		uslep,
		fieldSize,
		IREG,
		slope,
		slopeLength,
	)

	return res
}

func GetStripCroppingRecord5(
	rec InputParams.UserInputStepReceiver,
	gv *GlobalValueUtil.GlobalVar,
	SC BmpParams.StripCroppingParams,
) string {
	res := ""

	// get params
	USLECList := gv.USLEC
	NList := gv.N
	CNValueStr := gv.CNValue
	BareValueStr := gv.BareSoil

	// day, month, year, c = uslec, n, cn
	temp := "%2d, %2d, %4d,   %.5f, %.5f, %d"

	// CN value for another crop
	Crop := SC.Crop

	CDCrop := CommodityDatabaseCropRepository.FindCommodityDatabaseCropBySiteName(Crop)
	CropClass := CDCrop.Classes

	// these lists are from commodity database USLEC
	USLECListOther := CommodityDatabaseUSLECRepository.FindUSLECValuesByCropClass(CropClass)
	NListOther := CommodityDatabaseUSLECRepository.FindNValueByClass(CropClass)

	// day, month, year : int
	DateLiWithCN := InputFileGenerateUtil.GenerateDateListWithCN(rec, CNValueStr, BareValueStr)
	for i := 0; i < 24; i++ {
		if i != 0 {
			res += "\n"
		}
		// bmp cn
		CNValue := DateLiWithCN[i][3]
		CNFlag := DateLiWithCN[i][4]
		if CNFlag == 1 {
			CNValue -= 3
		}

		// bmp c
		C := USLECList[i]
		C = C*(SC.Rate/100) + USLECListOther[i]*(1-SC.Rate/100)

		// bmp n
		N := NList[i]
		N = N*(SC.Rate/100) + NListOther[i]*(1-SC.Rate/100)

		res += fmt.Sprintf(temp,
			DateLiWithCN[i][0],
			DateLiWithCN[i][1],
			DateLiWithCN[i][2],
			C,
			N,
			CNValue,
		)
	}

	return res
}

func GetStripCroppingRecord9(rec InputParams.UserInputStepReceiver, gv *GlobalValueUtil.GlobalVar) string {
	// post params
	IrrigationType := rec.IrrigationType
	CN := gv.CNValue

	cn, _ := strconv.ParseFloat(CN, 64)
	cn -= 3

	S := (2540 / cn) - 25.4

	res := fmt.Sprintf("%d, 0.1, 0.5, %f, False, 0", IrrigationType, S)
	return res
}
