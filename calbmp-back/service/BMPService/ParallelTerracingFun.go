package BMPService

import (
	"calbmp-back/Params/BmpParams"
	"calbmp-back/Params/InputParams"
	"calbmp-back/Repository/ChorizonRepository"
	"calbmp-back/Repository/RainDistributionIregRepository"
	"calbmp-back/util/GlobalValueUtil"
	"calbmp-back/util/InputFileGenerateUtil"
	"fmt"
	"strconv"
)

func ParallelTerracing(
	rec1 InputParams.UserInputStepReceiver,
	rec2 InputParams.UserInputStepReceiver,
	gv *GlobalValueUtil.GlobalVar,
	PT BmpParams.ParallelTerracingParams,
) (string, string, string) {
	Record3 := GetParallelTerracingRecord3(rec1, PT)
	Record5 := GetParallelTerracingRecord5(rec2, gv)
	Record9 := GetParallelTerracingRecord9(rec2, gv)

	return Record3, Record5, Record9
}

func GetParallelTerracingRecord3(
	rec InputParams.UserInputStepReceiver,
	PT BmpParams.ParallelTerracingParams,
) string {
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
		false,
		rec.KnowSlope,
		rec.Slope,
	)

	if rec.KnowSlope {
		slope = rec.Slope
	}

	// -------------- BMP part --------------
	if PT.Type == "type1" {
		if slope > 24 {
			uslep = 1.0
		} else if slope <= 24 && slope > 18 {
			uslep = 0.12
		} else if slope <= 18 && slope > 12 {
			uslep = 0.1
		} else if slope <= 12 && slope > 7 {
			uslep = 0.12
		} else if slope <= 7 && slope > 2 {
			uslep = 0.14
		} else {
			uslep = 0.17
		}
	} else if PT.Type == "type2" {
		if slope > 24 {
			uslep = 1.0
		} else if slope <= 24 && slope > 18 {
			uslep = 0.05
		} else if slope <= 18 && slope > 12 {
			uslep = 0.05
		} else if slope <= 12 && slope > 7 {
			uslep = 0.05
		} else if slope <= 7 && slope > 2 {
			uslep = 0.05
		} else {
			uslep = 0.06
		}
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

func GetParallelTerracingRecord5(
	rec InputParams.UserInputStepReceiver,
	gv *GlobalValueUtil.GlobalVar,
) string {
	res := ""

	// get params
	USLECList := gv.USLEC
	NList := gv.N
	CNValueStr := gv.CNValue
	BareValueStr := gv.BareSoil

	// day, month, year, c = uslec, n, cn
	temp := "%2d, %2d, %4d,   %.5f, %.5f, %d"

	// day, month, year : int
	DateLiWithCN := InputFileGenerateUtil.GenerateDateListWithCN(rec, CNValueStr, BareValueStr)
	for i := 0; i < 24; i++ {
		if i != 0 {
			res += "\n"
		}

		var CNValue int
		if DateLiWithCN[i][4] == 1 {
			CNValue = DateLiWithCN[i][3] - 6
		} else {
			CNValue = DateLiWithCN[i][3]
		}

		res += fmt.Sprintf(temp,
			DateLiWithCN[i][0],
			DateLiWithCN[i][1],
			DateLiWithCN[i][2],
			USLECList[i],
			NList[i],
			CNValue,
		)
	}

	return res
}

func GetParallelTerracingRecord9(rec InputParams.UserInputStepReceiver, gv *GlobalValueUtil.GlobalVar) string {
	// post params
	IrrigationType := rec.IrrigationType
	CN := gv.CNValue

	cn, _ := strconv.ParseFloat(CN, 64)
	cn -= 6

	S := (2540 / cn) - 25.4

	res := fmt.Sprintf("%d, 0.1, 0.5, %f, False, 0", IrrigationType, S)
	return res
}
