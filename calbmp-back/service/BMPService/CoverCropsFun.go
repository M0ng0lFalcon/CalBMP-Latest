package BMPService

import (
	"calbmp-back/Params/BmpParams"
	"calbmp-back/Params/InputParams"
	"calbmp-back/Repository/CommodityDatabaseCropRepository"
	"calbmp-back/Repository/CommodityDatabaseUSLECRepository"
	"calbmp-back/dto/Record5DTO"
	"calbmp-back/util/GlobalValueUtil"
	"calbmp-back/util/InputFileGenerateUtil"
	"calbmp-back/util/StringUtil"
	"calbmp-back/util/TimeUtil"
	"sort"
	"time"
)

func CoverCrops(
	rec InputParams.UserInputStepReceiver,
	gv *GlobalValueUtil.GlobalVar,
	CP BmpParams.CoverCropsParams,
) string {
	res := ""

	// make crop unit list
	CropUnitList := make([]Record5DTO.Record5DTO, 2)
	// get org crop unit
	CropUnitList[0] = Record5DTO.Record5DTO{
		EmergenceDate: TimeUtil.ParseTimeString(rec.Emergence),
		HarvestDate:   TimeUtil.ParseTimeString(rec.Harvest),
		USLECList:     gv.USLEC,
		NList:         gv.N,
		CNValue:       StringUtil.ConvertToInt(gv.CNValue),
		BareValue:     StringUtil.ConvertToInt(gv.BareSoil),
	}
	// get cover crop unit
	// -------------- bmp part --------------
	// get cover crop cn value
	Mukey := rec.MuKey
	Cokey := rec.CoKey

	CoverCropCNValue, CoverCropBareValue := CommodityDatabaseCropRepository.GetCNValueAndBareValue(CP.CoverCrop, Cokey, Mukey)

	// get cover crop usles value and N value
	CoverCropCDCrop := CommodityDatabaseCropRepository.FindCommodityDatabaseCropBySiteName(CP.CoverCrop)
	CoverCropClass := CoverCropCDCrop.Classes
	CoverCropUSLECList := CommodityDatabaseUSLECRepository.FindUSLECValuesByCropClass(CoverCropClass)
	CoverCropNList := CommodityDatabaseUSLECRepository.FindNValueByClass(CoverCropClass)

	// -------------- bmp part --------------
	CropUnitList[1] = Record5DTO.Record5DTO{
		EmergenceDate: TimeUtil.ParseTimeString(CP.CoverCropEmergence),
		HarvestDate:   TimeUtil.ParseTimeString(CP.CoverCropHarvest),
		USLECList:     CoverCropUSLECList,
		NList:         CoverCropNList,
		CNValue:       StringUtil.ConvertToInt(CoverCropCNValue),
		BareValue:     StringUtil.ConvertToInt(CoverCropBareValue),
	}

	sort.SliceStable(CropUnitList, func(i, j int) bool {
		return CropUnitList[i].EmergenceDate.Before(CropUnitList[j].EmergenceDate)
	})

	curIdx := 0
	curUnit := CropUnitList[curIdx]

	// generate date list
	/*
		flag = true  -> day = 16
		flag = false -> day = 1
	*/
	flag := true   // control day
	CNFlag := true // cn value or bare soil
	CoverFlag := false
	// cnt control day of a year,
	// a year have 2 days
	cnt := 0

	curDay := curUnit.EmergenceDate.Day()
	curMonth := int(curUnit.EmergenceDate.Month())
	curYear := curUnit.EmergenceDate.Year()
	for i := 0; i < 24; i++ {
		// Check the CNFlag and decide which CN to use
		// CN Value:
		// CNFlag = true  -> CNValue
		// CNFlag = false -> BareValue
		var line string
		if CNFlag == true {
			if CoverFlag == true {
				line = InputFileGenerateUtil.MakeRecord5line(
					curDay,
					curMonth,
					curYear,
					curUnit.CNValue-2,
					curUnit.USLECList[i],
					0.25,
				)
			} else {
				line = InputFileGenerateUtil.MakeRecord5line(
					curDay,
					curMonth,
					curYear,
					curUnit.CNValue,
					curUnit.USLECList[i],
					curUnit.NList[i],
				)
			}
			res += line
		} else {
			N := curUnit.NList[i]
			CN := curUnit.BareValue
			if CoverFlag == true {
				N = 0.25
				CN -= 2
			}
			line = InputFileGenerateUtil.MakeRecord5line(
				curDay,
				curMonth,
				curYear,
				CN,
				curUnit.USLECList[i],
				N,
			)
			res += line
		}

		if i != 23 {
			res += "\n"
		}

		// check day
		if flag == true {
			curDay = 16
		} else {
			curDay = 1
		}

		// Change day from 1 to 16
		// and then increase month by one
		if cnt >= 1 {
			curMonth = curMonth%12 + 1

			// If curMonth equals 1,
			// represent for next year
			if curMonth == 1 {
				curYear++
			}
			// reset cnt to 0
			cnt = 0
		} else {
			cnt++
		}

		curDate := time.Date(
			curYear,
			time.Month(curMonth),
			curDay,
			0,
			0,
			0,
			0,
			time.UTC,
		)
		if curDate.After(curUnit.HarvestDate) {
			CNFlag = false
			CoverFlag = true
			if curDate.After(CropUnitList[1].EmergenceDate) {
				curUnit = CropUnitList[1]
				CNFlag = true

				curDay = curUnit.EmergenceDate.Day()
				curMonth = int(curUnit.EmergenceDate.Month())
				curYear = curUnit.EmergenceDate.Year()
			}
		}

		flag = !flag
	}

	return res
}
