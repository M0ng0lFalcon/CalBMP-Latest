package BMPService

import (
	"calbmp-back/Params/BmpParams"
	"calbmp-back/Params/InputParams"
	"calbmp-back/Repository/CommodityDatabaseCropRepository"
	"calbmp-back/Repository/CommodityDatabaseUSLECRepository"
	"calbmp-back/dto/Record5DTO"
	"calbmp-back/util/CRRMUtil"
	"calbmp-back/util/GlobalValueUtil"
	"calbmp-back/util/InputFileGenerateUtil"
	"calbmp-back/util/StringUtil"
	"calbmp-back/util/TimeUtil"
	"math"
	"time"
)

// CropRotation : crop rotation with residue management
// c, n, cn 前一个收获之后到原作物收获为止
// cn : 前一个有residue的话 fallow的时候 -2
// cn : 前一个有residue的话 cn -2
func CropRotation(
	rec InputParams.UserInputStepReceiver,
	gv *GlobalValueUtil.GlobalVar,
	CRRM BmpParams.CropRotationAndResidueManagementParams,
) string {
	res := ""

	// get params
	Mukey := rec.MuKey
	Cokey := rec.CoKey
	CropInfos := CRRM.CropInfos

	TotalCnt := len(CropInfos)

	// 1. sort crop infos
	CRRMUtil.SortDateByEmergence(CropInfos)

	// 2. generate crop rotation and residue management dto
	CRRMList := make([]Record5DTO.Record5DTO, 0)
	for i := 0; i < TotalCnt; i++ {
		// 1. get crop class
		CropName := CropInfos[i].CropName
		CDCrop := CommodityDatabaseCropRepository.FindCommodityDatabaseCropBySiteName(CropName)
		CropClass := CDCrop.Classes

		// 2. get list of uslec and n
		USLECList := CommodityDatabaseUSLECRepository.FindUSLECValuesByCropClass(CropClass)
		NList := CommodityDatabaseUSLECRepository.FindNValueByClass(CropClass)

		// 3. get cn value and bare value
		CNValue, BareValue := CommodityDatabaseCropRepository.GetCNValueAndBareValue(
			CropName,
			Cokey,
			Mukey,
		)

		// make crop rotation and residue management unit
		tmpCN := StringUtil.ConvertToInt(CNValue)
		if CropInfos[i].ResidueValue != 0 {
			tmpCN -= 3
		}
		CRRMUnit := Record5DTO.Record5DTO{
			UnitId:        i,
			CropName:      CropInfos[i].CropName,
			EmergenceDate: TimeUtil.ParseTimeString(CropInfos[i].EmergenceDate),
			HarvestDate:   TimeUtil.ParseTimeString(CropInfos[i].HarvestDate),
			USLECList:     USLECList,
			NList:         NList,
			CNValue:       tmpCN,
			BareValue:     StringUtil.ConvertToInt(BareValue),
			ResidueValue:  CropInfos[i].ResidueValue,
		}

		CRRMList = append(CRRMList, CRRMUnit)
	}

	// generate record 5

	/*
		flag = true  -> day = 16
		flag = false -> day = 1
	*/
	flag := true // control day
	/*
		ResidueAndCNFlag Control values in the case of no planting
		== 1 -> change cn, n, uslec by residue formula
		== 2 -> use org cn, n, uslec and check residue state, if pre crop have residue -> change c,cn,n
		== 3 -> use org bare, n, uslec
	*/
	ResidueAndCNFlag := 2
	// cnt control day of a year,
	// a year have 2 days
	cnt := 0

	curIdx := 0
	curUnit := CRRMList[curIdx]

	curDay := curUnit.EmergenceDate.Day()
	curMonth := int(curUnit.EmergenceDate.Month())
	curYear := curUnit.EmergenceDate.Year()

	var line string
	for i := 0; i < 24; i++ {
		// value of c,n when residue
		// value of C

		// change cn,n,c by state
		if ResidueAndCNFlag == 2 {
			// residue and cn flag == 2
			// use org cn, n, uslec and check residue state, if pre crop have residue -> change c,cn,n
			var residueValue float64
			//var cropName string
			if curUnit.UnitId != 0 {
				residueValue = CRRMList[curUnit.UnitId-1].ResidueValue
				//cropName = CRRMList[curUnit.UnitId-1].CropName
			}
			//k := math.Exp(-0.00115 * residueValue * 1000)
			//uslecMn := CommodityDatabaseCropRepository.GetUSLECmnBySiteName(cropName)
			//C := math.Pow(0.8, k) * math.Pow(uslecMn, 1-k)
			// value of N
			var N float64
			if residueValue >= 0 && residueValue < 2 {
				N = 0.2
			} else if residueValue >= 2 && residueValue <= 10 {
				N = 0.3
			}
			// check residue state of pre crop
			var c, n float64
			var cn int
			if curUnit.UnitId != 0 {
				if CRRMList[curUnit.UnitId-1].ResidueValue != 0 {
					//c = C
					c = curUnit.USLECList[i]
					n = N
					cn = curUnit.CNValue + 1
				}
			} else {
				c = curUnit.USLECList[i]
				n = curUnit.NList[i]
				cn = curUnit.CNValue + 3
			}
			line = InputFileGenerateUtil.MakeRecord5line(
				curDay,
				curMonth,
				curYear,
				cn,
				c,
				n,
			)
			res += line
		} else if ResidueAndCNFlag == 3 {
			// residue and cn flag == 3
			// use bare soil
			c := curUnit.USLECList[i]
			n := curUnit.NList[i]
			cn := curUnit.BareValue
			line = InputFileGenerateUtil.MakeRecord5line(
				curDay,
				curMonth,
				curYear,
				cn,
				c,
				n,
			)
			res += line
		} else if ResidueAndCNFlag == 1 {
			// residue and cn flag == 1
			// use formula of residue
			k := math.Exp(-0.00115 * curUnit.ResidueValue * 1000)
			uslecMn := CommodityDatabaseCropRepository.GetUSLECmnBySiteName(curUnit.CropName)
			C := math.Pow(0.8, k) * math.Pow(uslecMn, 1-k)

			var N float64
			if curUnit.ResidueValue >= 0.5 && curUnit.ResidueValue < 2 {
				N = 0.2
			} else if curUnit.ResidueValue >= 2 && curUnit.ResidueValue <= 10 {
				N = 0.3
			}
			// change uslec -> c
			c := C
			// change n
			n := N
			// change cn
			cn := curUnit.BareValue - 2
			line = InputFileGenerateUtil.MakeRecord5line(
				curDay,
				curMonth,
				curYear,
				cn,
				c,
				n,
			)
			res += line
		}

		// -------------------------------------------
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

		// check date

		// 1. make a time obj
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
		// change to next CRRM item
		if curIdx != len(CRRMList)-1 {
			if curDate.After(CRRMList[curIdx+1].EmergenceDate) {
				curIdx++
				curUnit = CRRMList[curIdx]
				// refresh date
				curDay = curUnit.EmergenceDate.Day()
				curMonth = int(curUnit.EmergenceDate.Month())
				curYear = curUnit.EmergenceDate.Year()
				ResidueAndCNFlag = 2
			}
		}
		// check bare or residue
		if curDate.After(curUnit.HarvestDate) {
			if curUnit.ResidueValue != 0 {
				ResidueAndCNFlag = 1
			} else {
				ResidueAndCNFlag = 3
			}
		}

		flag = !flag
	}

	return res
}
