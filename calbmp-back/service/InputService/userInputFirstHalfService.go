package InputService

import (
	"calbmp-back/Params/InputParams"
	"calbmp-back/RedisUtil"
	"calbmp-back/Repository/ChorizonRepository"
	"calbmp-back/Repository/CommodityDatabaseCropRepository"
	"calbmp-back/Repository/RainDistributionIregRepository"
	"calbmp-back/Repository/SoilEvaAnetdRepository"
	"calbmp-back/util/GlobalValueUtil"
	"calbmp-back/util/InputFileGenerateUtil"
	"calbmp-back/util/StringUtil"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// -------------------------- user input step 1 --------------------------

func GetCurrentTime() string {
	currentTime := time.Now().Format("01/01/2006 15:04:05 PM")
	return "***  " + currentTime
}

func GetRecordA1(rec InputParams.UserInputStepReceiver, username string) string {
	res := ""
	fileRoot := "./przm5place/"

	// get params
	ifBmp := rec.IfBMP
	ifVfsm := rec.IfVfsm
	weaFilePath := rec.WeaFilePath

	curTime := RedisUtil.GetStringVal(username + "_curTime")
	if ifBmp == false {
		weaFilePath = curTime + "/baseline/weatherData.wea"
	} else {
		weaFilePath = curTime + "/bmp/" + weaFilePath
	}
	if ifVfsm {
		weaFilePath = curTime + "/vfsm_baseline/weatherData.wea"
	}

	// create wea file
	_, _ = os.Create(fileRoot + weaFilePath)

	// generate res
	res = fileRoot + weaFilePath

	return res
}

/*
1. get zts file path
2. create zts file
*/
func GetRecordA2(rec InputParams.UserInputStepReceiver, username string) string {
	res := ""
	fileRoot := "./przm5place/"

	// get params
	ifBmp := rec.IfBMP
	ztsFilePath := rec.ZtsFilePath

	curTime := RedisUtil.GetStringVal(username + "_curTime")
	if ifBmp == false {
		ztsFilePath = curTime + "/baseline/baseline.zts"
	} else {
		ztsFilePath = curTime + "/bmp/" + ztsFilePath
	}
	ifVfsm := rec.IfVfsm
	if ifVfsm {
		ztsFilePath = curTime + "/vfsm_baseline/baseline.zts"
	}

	// create zts file
	_, _ = os.Create(fileRoot + ztsFilePath)

	// generate res
	res = fileRoot + ztsFilePath

	return res
}

/*
! record A3 have default value
*/
func GetRecordA3() string {
	res := ""

	// default value
	recordA3 := "False,False,False,False,False,1,False"

	// generate res
	res += recordA3

	return res
}

/*
record 1: pfac, sfac, anetd
default value:

	pfac = 1
	sfac = 0.45
*/
func GetRecord1(rec InputParams.UserInputStepReceiver) string {
	res := ""

	// get params
	ZipCode := rec.ZipCode

	// find soil by zipcode
	SoilEvaAnetd := SoilEvaAnetdRepository.FindSoilEvaAnetdByZipCode(ZipCode)

	// get anetd value
	anetd := SoilEvaAnetd.ANETD_cm

	// generate res
	temp := "1, 0.45, %.1f"
	res = fmt.Sprintf(temp, anetd)

	return res
}

/*
record 2: MUSS
default value:

	MUSS = 4
*/
func GetRecord2() string {
	res := ""

	// generate res
	res = "4"

	return res
}

func GetRecord3(rec InputParams.UserInputStepReceiver, username string) string {
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
	if Chorizon.Mukey == "" {
		Chorizon = ChorizonRepository.FindChorizonListByMukeyAndCokey(Mukey, Cokey)[0]
	}
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

	RedisUtil.SetFloatKey(username+"_slope", slope)
	RedisUtil.SetFloatKey(username+"_slopeLength", slopeLength)
	RedisUtil.SetFloatKey(username+"_uslep", uslep)
	RedisUtil.SetFloatKey(username+"_fieldSize", fieldSize)

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

	RedisUtil.SetFloatKey(username+"_slope", slope)

	return res
}

/*
record 4 have default value
*/
func GetRecord4() string {
	res := ""

	CalBmpValue := "24, 0"
	res += CalBmpValue

	return res
}

// -------------------------- user input step 2 --------------------------

func GetRecord5(rec InputParams.UserInputStepReceiver, gv *GlobalValueUtil.GlobalVar) string {
	res := ""

	// get params
	USLECList := gv.USLEC
	NList := gv.N
	CNValue := gv.CNValue
	BareValue := gv.BareSoil

	// day, month, year, c = uslec, n, cn
	temp := "%2d, %2d, %4d,   %.5f, %.5f, %d"

	// day, month, year : int
	DateLiWithCN := InputFileGenerateUtil.GenerateDateListWithCN(rec, CNValue, BareValue)
	for i := 0; i < 24; i++ {
		if i != 0 {
			res += "\n"
		}
		res += fmt.Sprintf(temp,
			DateLiWithCN[i][0],
			DateLiWithCN[i][1],
			DateLiWithCN[i][2],
			USLECList[i],
			NList[i],
			DateLiWithCN[i][3],
		)
	}

	DateLiWithCNJson, _ := json.Marshal(DateLiWithCN)
	DateLiWithCNJsonStr := string(DateLiWithCNJson)
	RedisUtil.SetStringKey(rec.Username+"_DateLiWithCN", DateLiWithCNJsonStr)
	RedisUtil.SetStringKey(rec.Username+"_record5", res)

	return res
}

func GetRecord6() string {
	return "1"
}

func GetRecord7(rec InputParams.UserInputStepReceiver) string {
	// post params
	Emergence := StringUtil.FormatDate(rec.Emergence)
	Maturity := StringUtil.FormatDate(rec.Maturity)
	Harvest := StringUtil.FormatDate(rec.Harvest)
	SiteName := rec.Crop

	// db model
	CommodityDatabaseCrop := CommodityDatabaseCropRepository.FindCommodityDatabaseCropBySiteName(SiteName)

	// get depth
	depth := CommodityDatabaseCrop.Rootdepth_cm
	// get cover
	cover := CommodityDatabaseCrop.Canopycover
	// get height
	height := CommodityDatabaseCrop.Height_cm
	// get holdup
	holdup := CommodityDatabaseCrop.Holdup_cm

	// set emergence, maturity, harvest date
	// format: 16,2,1961
	EmergenceLi := strings.Split(Emergence, ",")
	Emergence = fmt.Sprintf("%s,%s, %s",
		StringUtil.DeleteFrontZero(EmergenceLi[2]),
		StringUtil.DeleteFrontZero(EmergenceLi[1]),
		EmergenceLi[0])
	MaturityLi := strings.Split(Maturity, ",")
	Maturity = fmt.Sprintf("%s,%s, %s",
		StringUtil.DeleteFrontZero(MaturityLi[2]),
		StringUtil.DeleteFrontZero(MaturityLi[1]),
		MaturityLi[0])
	HarvestLi := strings.Split(Harvest, ",")
	Harvest = fmt.Sprintf("%s,%s, %s",
		StringUtil.DeleteFrontZero(HarvestLi[2]),
		StringUtil.DeleteFrontZero(HarvestLi[1]),
		HarvestLi[0])

	// get post-harvest disposition
	postHarvestDisposition := CommodityDatabaseCrop.Scenarios_residue

	temp := "%s,   %s,  %s,      %s,      %s,       %s,     %s,        %s"
	res := fmt.Sprintf(temp,
		Emergence, Maturity, Harvest, depth, cover, height, holdup, postHarvestDisposition)
	return res
}

func GetRecord8() string {
	return " 2, False"
}

func GetRecord9(rec InputParams.UserInputStepReceiver, gv *GlobalValueUtil.GlobalVar) string {
	// post params
	IrrigationType := rec.IrrigationType
	CN := gv.CNValue

	template := "%d, 0.1, %f, %f, %s, 0"

	cn, _ := strconv.ParseFloat(CN, 64)

	S := (2540 / cn) - 25.4
	// if (user def) then S = (user def)

	var res string
	if IrrigationType == 0 { // irrigation type == flood irrigation
		res = fmt.Sprintf(template, IrrigationType, 0.0, S, "TRUE")
	} else if IrrigationType == 6 || IrrigationType == 7 {
		res = fmt.Sprintf("%d, 0.1, %f, %f, %s, 0",
			IrrigationType,
			0.5,
			rec.UserDefinedIrrgRate,
			"TRUE",
		)
	} else {
		res = fmt.Sprintf(template, IrrigationType, 0.5, S, "FALSE")
	}
	return res
}

func GetRecord14(rec InputParams.UserInputStepReceiver) string {
	res := ""
	// get params
	Mukey := rec.MuKey
	Cokey := rec.CoKey

	ChorizonList := ChorizonRepository.FindChorizonListByMukeyAndCokey(Mukey, Cokey)

	cnt := len(ChorizonList)

	res += strconv.Itoa(cnt)

	return res
}

/*
record 15: index,thk, Del, Dsp,   bd,  W0,    FC,    WP,    oc, snd, cly,  tmp
*/
func GetRecord15(rec InputParams.UserInputStepReceiver, gv *GlobalValueUtil.GlobalVar, username string) string {
	res := ""
	U2Num := 0

	// get params
	Mukey := rec.MuKey
	Cokey := rec.CoKey

	// database interaction
	ChorizonList := ChorizonRepository.FindChorizonListByMukeyAndCokey(Mukey, Cokey)

	temp := "%d, %d, %d, 0.0, %s, %f, %s, %s, %s, %s, %s, ,"
	// get values
	for i, v := range ChorizonList {
		if i != 0 {
			res += "\n"
		}
		// index
		thickess := StringUtil.ConvertToInt(v.Hzthk)
		// start: number of compartments
		NumOfCompartments := StringUtil.ConvertToInt(v.Hzthk)
		if thickess == 0 {
			NumOfCompartments = 1
		} else {
			for j := 10; j >= 1; j-- {
				if thickess%j == 0 {
					NumOfCompartments = thickess / j
					break
				}
			}
		}
		// stop: number of compartments
		bulkDensity := v.Dbthirdb_g_cm3
		W0 := StringUtil.Convert2Float(v.Wthirdbar_cm3_cm3)
		tjcsl := v.Wthirdbar_cm3_cm3
		dwhsl := v.Wfifteenbar_cm3_cm3
		organicCarbon := v.Orgc_per
		hsl := v.Sandtotal_per
		hntl := v.Claytotal_per

		RedisUtil.SetFloatKey(username+"_W0", W0)

		//U2Num += NumOfCompartments
		U2Num += NumOfCompartments
		res += fmt.Sprintf(temp,
			i+1,
			thickess,
			NumOfCompartments,
			bulkDensity,
			W0,
			tjcsl,
			dwhsl,
			organicCarbon,
			hsl,
			hntl,
		)
	}
	gv.U2Num = U2Num
	return res
}

func GetRecord16(rec InputParams.UserInputStepReceiver) string {
	//resLi := make([]string, 0)
	cam := rec.ApplicationMethod
	depthLi := rec.Depth
	rDepth := 2.0
	rDecline := 1.55
	Bypass := 0.266

	for i, v := range cam {
		if v == 8 && depthLi[i] > 2 {
			rDepth = depthLi[i] + 0.1
		}
	}

	res := fmt.Sprintf("%f, %f, %f", rDepth, rDecline, Bypass)
	return res
}

func GetRecord17() string {
	return "0.1,0,1.0"
}
