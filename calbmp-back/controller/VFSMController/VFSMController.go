package VFSMController

import (
	"calbmp-back/Params/InputParams"
	"calbmp-back/Params/VFSMParams"
	"calbmp-back/RedisUtil"
	"calbmp-back/Repository/ChorizonRepository"
	"calbmp-back/Repository/CropPesticideFinalRepository"
	"calbmp-back/Repository/PfstationRepository"
	"calbmp-back/Repository/SoilTextureFinalRepository"
	"calbmp-back/Repository/VegetationRepository"
	"calbmp-back/Res"
	"calbmp-back/service/InputService"
	"calbmp-back/service/ResultService"
	"calbmp-back/service/VFSMService"
	"calbmp-back/util/FileUtil"
	"calbmp-back/util/GlobalValueUtil"
	"calbmp-back/util/InputFileGenerateUtil"
	"calbmp-back/util/JsonUtil"
	"calbmp-back/util/StringUtil"
	"calbmp-back/util/TimeUtil"
	"calbmp-back/util/WeatherDataUtil"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

func GenerateInputFile(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	var params VFSMParams.VfsmMainParams
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		Res.FailMsg(ctx, "[GenerateInputFile] bind json error")
		return
	}

	rootDir := "./inputs" // dir of vfsm input files
	VfsmMin := 0
	VfsmMax := -1

	// parse baseline input file
	inp1Path := "./przm5place/" + params.CreatedAt + "/baseline/step1.json"
	inp2Path := "./przm5place/" + params.CreatedAt + "/baseline/step2.json"
	var inp1 InputParams.UserInputStepReceiver
	var inp2 InputParams.UserInputStepReceiver
	JsonUtil.ReadJson(inp1Path, &inp1)
	JsonUtil.ReadJson(inp2Path, &inp2)

	// parse baseline wea file
	weaPath := "./przm5place/" + params.CreatedAt + "/baseline/weatherData.wea"
	weaJsonPath := "./przm5place/" + params.CreatedAt + "/baseline/weatherData.json"
	WeatherDataUtil.Wea2json(weaPath, weaJsonPath)
	weaJson := make(map[string]map[string]string)
	JsonUtil.ReadJson(weaJsonPath, &weaJson)

	// init parameter -------------- start --------------
	vegModel := VegetationRepository.FindByVegetation(params.IkwParam.Vegetation)
	// igr
	params.IgrParam.Vegetation = params.IkwParam.Vegetation
	params.IgrParam.VegModel = vegModel
	// ikw
	params.IkwParam.CropName = inp2.Crop
	params.IkwParam.VegModel = vegModel
	params.IkwParam.Username = username.(string)
	// irn
	params.IrnParam.ZipCode = inp1.ZipCode
	params.IrnParam.Pfstation = PfstationRepository.FindByZipCode(params.IrnParam.ZipCode)
	// isd
	params.IsdParam.Mukey = inp1.MuKey
	params.IsdParam.Cokey = inp1.CoKey
	params.IsdParam.Chorizon = ChorizonRepository.FindChorizonListByMukeyAndCokey(
		params.IsdParam.Mukey,
		params.IsdParam.Cokey,
	)[0]
	// iso
	params.IsoParam.Username = username.(string)
	params.IsoParam.Mukey = inp1.MuKey
	params.IsoParam.Cokey = inp1.CoKey
	params.IsoParam.SoilTexture = SoilTextureFinalRepository.FindByMukeyAndCokey(
		params.IsoParam.Mukey,
		params.IsoParam.Cokey,
	)
	// iwq
	params.IwqParam.CreatedAt = params.CreatedAt
	params.IwqParam.Chorizon = ChorizonRepository.FindChorizonListByMukeyAndCokey(
		params.IsdParam.Mukey,
		params.IsdParam.Cokey,
	)[0]

	// get text result of zts file
	params.TextRes.CreatedTime = params.CreatedAt
	params.TextRes.PesticideList = inp2.Pesticide
	params.TextRes.ScenarioType = "baseline"
	params.TextRes.FieldSize = RedisUtil.GetFloatVal(username.(string) + "_fieldSize")
	// init parameter -------------- start --------------

	// make dir of input files and output dir
	targetDir := filepath.Join(rootDir)
	outputDir := filepath.Join("./output")
	jsonDir := filepath.Join(rootDir, params.CreatedAt, "json")
	if _, errExist := os.Stat(targetDir); os.IsNotExist(errExist) {
		errMkDir := os.MkdirAll(targetDir, os.ModePerm)
		if errMkDir != nil {
			log.Println(errMkDir)
		}
	}
	if _, errExist := os.Stat(outputDir); os.IsNotExist(errExist) {
		errMkDir := os.MkdirAll(outputDir, os.ModePerm)
		if errMkDir != nil {
			log.Println(errMkDir)
		}
	}
	if _, errExist := os.Stat(jsonDir); os.IsNotExist(errExist) {
		errMkDir := os.MkdirAll(jsonDir, os.ModePerm)
		if errMkDir != nil {
			log.Println(errMkDir)
		}
	}

	// rerun baseline for iwq
	runBaseline(inp1, inp2, username.(string), params.CreatedAt)

	ztsResult := ResultService.GetTextResultFun(params.TextRes, "org")
	textZtsJson := "./przm5place/" + params.CreatedAt + "/vfsm_baseline/textZts.json"
	JsonUtil.WriteJson(textZtsJson, ztsResult)

	reRunZtsJson := make(map[string]map[string]string)
	baselineZtsJsonPath := "./przm5place/" + params.CreatedAt + "/vfsm_baseline/baselineZts.json"
	JsonUtil.ReadJson(baselineZtsJsonPath, &reRunZtsJson)
	pesticideDate, _, _ := getPesticideDate(textZtsJson)

	// init VFSM progress
	progressKeyName := username.(string) + "_vfsm_progress"
	RedisUtil.SetFloatKey(progressKeyName, 0.0)

	// Run VFSM model
	for i, v := range ztsResult.Date {
		dd, _ := time.Parse("2006/1/2", v)
		if dd.Before(pesticideDate) {
			continue
		}
		// set parameter -------------- start --------------
		// iro
		if ztsResult.Water["IRRG"][i] != 0.0 {
			params.IroParam.Prec = ztsResult.Water["IRRG"][i]
			params.IroParam.IsRain = true
		} else {
			params.IroParam.Prec = ztsResult.Water["PRCP"][i]
			params.IroParam.IsRain = false
		}
		// isd
		ESLS := ztsResult.Sediment["ESLS"][i]
		RUNF := ztsResult.Water["RUNF"][i]
		AREA := StringUtil.Convert2Float(inp1.FieldSize)
		params.IsdParam.CI = ESLS / (RUNF * AREA * 40.4685642)
		// iwq
		DgT := make([]string, 0)
		DgTheta := make([]string, 0)
		if i != 0 {
			days, _ := TimeUtil.GetDaysBetween2Date("2006/1/2", ztsResult.Date[i-1], v)
			params.IwqParam.NDGDAY = days

			d1, _ := time.Parse("2006/1/2", ztsResult.Date[i-1])
			d2, _ := time.Parse("2006/1/2", ztsResult.Date[i])
			//d2 = d2.AddDate(0, 0, 1)
			for d1 != d2 {
				tTemp := weaJson[d1.Format("2006/01/02")]["Temperature"]
				DgT = append(DgT, tTemp)

				thetaTemp := reRunZtsJson[d1.Format("2006/1/2")]["SWTR1"]
				DgTheta = append(DgTheta, thetaTemp)
				d1 = d1.AddDate(0, 0, 1)
			}
			params.IwqParam.DgT = DgT
			params.IwqParam.DgTheta = DgTheta
		} else {
			params.IwqParam.NDGDAY = int64(TimeUtil.GetDOY("2006/1/2", ztsResult.Date[i]))

			d1, _ := time.Parse("2006/1/2", ztsResult.Date[i])
			d1 = d1.AddDate(0, 0, -d1.YearDay()+1)

			d2, _ := time.Parse("2006/1/2", ztsResult.Date[i])
			d2 = d2.AddDate(0, 0, 1)
			for d1 != d2 {
				tTemp := weaJson[d1.Format("2006/01/02")]["Temperature"]
				DgT = append(DgT, tTemp)

				thetaTemp := reRunZtsJson[d1.Format("2006/1/2")]["SWTR1"]
				DgTheta = append(DgTheta, thetaTemp)
				d1 = d1.AddDate(0, 0, 1)
			}
			params.IwqParam.DgT = DgT
			params.IwqParam.DgTheta = DgTheta
		}

		// irn
		params.IrnParam.Prec = ztsResult.Water["IRRG"][i]
		// set parameter -------------- end --------------
		pesticideCnt := make(map[string]int)
		for j, pesticide := range inp2.Pesticide {
			// set parameter -------------- start --------------
			// set json path
			pesticideId := j + 1
			vfsmID := int64(pesticideId + i*len(inp2.Pesticide))
			if VfsmMin == 0 {
				VfsmMin = int(vfsmID)
			}
			params.IgrParam.JsonPath = filepath.Join(jsonDir, fmt.Sprintf("Igr_%d.json", vfsmID))
			params.IkwParam.JsonPath = filepath.Join(jsonDir, fmt.Sprintf("Ikw_%d.json", vfsmID))
			params.IrnParam.JsonPath = filepath.Join(jsonDir, fmt.Sprintf("Irn_%d.json", vfsmID))
			params.IroParam.JsonPath = filepath.Join(jsonDir, fmt.Sprintf("Iro_%d.json", vfsmID))
			params.IsdParam.JsonPath = filepath.Join(jsonDir, fmt.Sprintf("Isd_%d.json", vfsmID))
			params.IsoParam.JsonPath = filepath.Join(jsonDir, fmt.Sprintf("Iso_%d.json", vfsmID))
			params.IwqParam.JsonPath = filepath.Join(jsonDir, fmt.Sprintf("Iwq_%d.json", vfsmID))
			// ikw
			params.IkwParam.VfsmID = vfsmID
			// iwq
			if val, ok := pesticideCnt[pesticide]; ok {
				pesticideCnt[pesticide] = val + 1
			} else {
				pesticideCnt[pesticide] = 1
			}
			rflxKey := fmt.Sprintf("RFLX_%s_%d_TSER", pesticide, pesticideCnt[pesticide])
			eflxKey := fmt.Sprintf("EFLX_%s_%d_TSER", pesticide, pesticideCnt[pesticide])
			RFLX := ztsResult.Pesticide[rflxKey][i]
			EFLX := ztsResult.Pesticide[eflxKey][i]
			params.IwqParam.DGPIN = (RFLX + EFLX) * 1e7
			params.IwqParam.CropPesticide = CropPesticideFinalRepository.FindCropPesticideByChemicalName(pesticide)
			// iro
			params.IroParam.IrnJson = params.IrnParam.JsonPath

			// set parameter -------------- end --------------

			// generate input file string
			vfsmStringMap := map[string]string{
				"igr": VFSMService.GenerateIgrFun(params.IgrParam), // +
				"ikw": VFSMService.GenerateIkwFun(params.IkwParam), // slope
				"irn": VFSMService.GenerateIrnFun(params.IrnParam), //
				"iro": VFSMService.GenerateIroFun(params.IroParam), //
				"isd": VFSMService.GenerateIsdFun(params.IsdParam), //
				"iso": VFSMService.GenerateIsoFun(params.IsoParam), //
				"iwq": VFSMService.GenerateIwqFun(params.IwqParam), //
			}

			// write to file
			for suffix, content := range vfsmStringMap {
				// write input files
				outFilename := fmt.Sprintf("%s/%s_%d.%s", targetDir, params.CreatedAt, params.IkwParam.VfsmID, suffix)
				errWrite := FileUtil.WriteToFile(outFilename, []string{content})
				if errWrite != nil {
					log.Println("[GenerateInputFile] -> write inp file, err:", errWrite)
				}

				// rewrite vfsm max
				VfsmMax = int(params.IkwParam.VfsmID)
				// run przm5.exe
				exeFi := "./vfsm.exe"
				arg := fmt.Sprintf("%s_%d", params.CreatedAt, params.IkwParam.VfsmID)
				FileUtil.DoExecute(exeFi, arg)
			}

			progress := float64(i+1) / float64(len(ztsResult.Date))
			RedisUtil.SetFloatKey(progressKeyName, progress)
		}
	}
	// set min and max val to redis
	vfsmMinKeyName := username.(string) + "_vfsm_min"
	vfsmMaxKeyName := username.(string) + "_vfsm_max"
	RedisUtil.SetIntVal(vfsmMinKeyName, VfsmMin)
	RedisUtil.SetIntVal(vfsmMaxKeyName, VfsmMax)
	Res.Success(ctx, gin.H{"min": VfsmMin, "max": VfsmMax}, "[*] VFSM model run successfully")
}

func GetProgress(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	progressKeyName := username.(string) + "_vfsm_progress"
	progress := RedisUtil.GetFloatVal(progressKeyName)
	if progress == 1 {
		vfsmMinKeyName := username.(string) + "_vfsm_min"
		vfsmMaxKeyName := username.(string) + "_vfsm_max"
		VfsmMin := RedisUtil.GetIntVal(vfsmMinKeyName)
		VfsmMax := RedisUtil.GetIntVal(vfsmMaxKeyName)
		Res.Success(ctx, gin.H{"min": VfsmMin, "max": VfsmMax}, "[*] VFSM model run successfully")
	} else {
		Res.Success(ctx, gin.H{"progress": progress}, "[*] get vfsm progress")
	}
}

// ----------------------------------------------- Util -----------------------------------------------

func runBaseline(inp1, inp2 InputParams.UserInputStepReceiver, username, curTime string) {
	// run step 1 ------------------------------ start  ------------------------------
	// path of template
	inp1.IfVfsm = true
	TemplateFilePath := "./public/przm5Template/A1_4.txt"
	// main fun
	var gv GlobalValueUtil.GlobalVar
	resMap, _ := InputService.RepPlaceHolder2Fun(TemplateFilePath, inp1, &gv, username)
	// save res to file
	FileUtil.Mkdir("./przm5place/" + curTime)
	FileUtil.Mkdir("./przm5place/" + curTime + "/vfsm_baseline")
	FileUtil.Mkdir("./przm5place/" + curTime + "/vfsm_bmp")
	inp1Path := "./przm5place/" + curTime + "/vfsm_baseline/baseline1.inp"
	err := FileUtil.WriteMapToFile(inp1Path, TemplateFilePath, resMap)
	if err != nil {
		log.Println("[runBaseline] WriteMapToFile,err:", err)
		return
	}
	// save step1 struct to file
	jsonString, _ := json.Marshal(inp1)
	step1path := "./przm5place/" + curTime + "/vfsm_baseline/step1.json"
	FileUtil.WriteString2File(step1path, string(jsonString))
	// run step 1 ------------------------------ end  ------------------------------

	// run step 2 ------------------------------ start  ------------------------------
	inp2.Crop = "ALFALFA-GRASS MIXTURE"
	for i, _ := range inp2.Amount {
		inp2.Amount[i] = "0.0"
	}
	// save some var to redis
	RedisUtil.SetStringKey(username+"_crop", inp2.Crop)
	RedisUtil.SetStringKey(username+"_cokey", inp2.CoKey)

	// main fun
	TemplateFilePath = "./public/przm5Template/5_U2.txt"
	gv = GlobalValueUtil.GetGlobalVariables(inp2)
	resMap, _ = InputService.RepPlaceHolder2Fun(TemplateFilePath, inp2, &gv, username)
	// save res to file
	inp2Path := "./przm5place/" + curTime + "/vfsm_baseline/baseline2.inp"
	err = FileUtil.WriteMapToFile(inp2Path, TemplateFilePath, resMap)
	if err != nil {
		log.Println("[runBaseline] WriteMapToFile,err:", err)
		return
	}
	// merge test1.inp and test2.inp
	targetPath := "./przm5place/" + curTime + "/vfsm_baseline/PRZM5.inp"
	FileUtil.MergeInpFile(inp1Path, inp2Path, targetPath)
	// generate wea file
	baselineWeaPath := "./przm5place/" + curTime + "/vfsm_baseline/weatherData.wea"
	WeatherDataUtil.GenerateWeaData(inp2, baselineWeaPath)
	// check irrigation type
	if inp2.IrrigationType == 0 {
		InputFileGenerateUtil.ChangeWeatherData(inp2.IrrigationDate, inp2.IrrigationAmount, baselineWeaPath)
	}

	// save step2 struct to file
	jsonString, _ = json.Marshal(inp2)
	step2path := "./przm5place/" + curTime + "/vfsm_baseline/step2.json"
	FileUtil.WriteString2File(step2path, string(jsonString))

	// run przm5.exe
	curOs := runtime.GOOS
	var cmd *exec.Cmd
	switch curOs {
	case "linux":
		wineLauncher := viper.GetString("wine.launcher")
		cmd = exec.Command(wineLauncher, "./przm5place/PRZM5.exe", "./przm5place/"+curTime+"/vfsm_baseline")
	case "windows":
		cmd = exec.Command("./przm5place/PRZM5.exe", "./przm5place/"+curTime+"/vfsm_baseline")
	}
	err = cmd.Run()
	ResultService.Zts2Json(
		"./przm5place/"+curTime+"/vfsm_baseline/baseline.zts",
		"./przm5place/"+curTime+"/vfsm_baseline/baselineZts.json",
	)
}
