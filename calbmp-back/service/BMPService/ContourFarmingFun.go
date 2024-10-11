package BMPService

import (
	"calbmp-back/Params/InputParams"
	"calbmp-back/Repository/ChorizonRepository"
	"calbmp-back/Repository/RainDistributionIregRepository"
	"calbmp-back/util/GlobalValueUtil"
	"calbmp-back/util/InputFileGenerateUtil"
	"fmt"
	"strconv"
)

func ContourFarming(
	rec1 InputParams.UserInputStepReceiver,
	rec2 InputParams.UserInputStepReceiver,
	gv *GlobalValueUtil.GlobalVar,
) (string, string, string) {
	record3 := ContourFarming3(rec1)
	record5 := ChangeRecordCNValue(3, gv, rec2)
	record9 := ContourFarming9(rec2, gv)

	return record3, record5, record9
}

func ContourFarming3(rec InputParams.UserInputStepReceiver) string {
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
		uslep = 0.8
	} else if slope <= 18 && slope > 12 {
		uslep = 0.7
	} else if slope <= 12 && slope > 7 {
		uslep = 0.6
	} else if slope <= 7 && slope > 2 {
		uslep = 0.5
	} else {
		uslep = 0.6
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

func ContourFarming9(rec InputParams.UserInputStepReceiver, gv *GlobalValueUtil.GlobalVar) string {
	// post params
	IrrigationType := rec.IrrigationType
	CN := gv.CNValue

	cn, _ := strconv.ParseFloat(CN, 64)
	cn -= 3

	S := (2540 / cn) - 25.4

	res := fmt.Sprintf("%d, 0.1, 0.5, %f, False, 0", IrrigationType, S)
	return res
}
