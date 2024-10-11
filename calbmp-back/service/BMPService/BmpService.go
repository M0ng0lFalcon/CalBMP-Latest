package BMPService

import (
	"calbmp-back/Params/BmpParams"
	"calbmp-back/RedisUtil"
	"calbmp-back/service/InputService"
	"calbmp-back/util/FileUtil"
	"calbmp-back/util/GlobalValueUtil"
	"calbmp-back/util/InputFileGenerateUtil"
	"calbmp-back/util/WeatherDataUtil"
	"github.com/spf13/viper"
	"log"
	"os/exec"
	"runtime"
)

func BmpScenario(rec BmpParams.BmpBasicParams, username string) {
	// 1. store params to var
	Step1Params := rec.Step1Params
	Step2Params := rec.Step2Params
	BmpId := rec.BmpId

	// 2. Variables required to initialize inP files
	TemplateFilePath1 := "./public/przm5Template/A1_4.txt"
	TemplateFilePath2 := "./public/przm5Template/5_U2.txt"

	// 3. generate inp value
	var gv GlobalValueUtil.GlobalVar
	inp1, _ := InputService.RepPlaceHolder2Fun(TemplateFilePath1, Step1Params, &gv, username)
	gv = GlobalValueUtil.GetGlobalVariables(Step2Params)
	inp2, _ := InputService.RepPlaceHolder2Fun(TemplateFilePath2, Step2Params, &gv, username)

	// 4. change inp value
	inp1, inp2 = ParseBmpFun(rec, inp1, inp2, &gv)

	// 5. save inp value
	curTime := RedisUtil.GetStringVal(username + "_curTime")
	basePath := "./przm5place/" + curTime + "/bmp/" + BmpId
	FileUtil.Mkdir(basePath)
	BmpInp1Path := basePath + "/bmp1.inp"
	BmpInp2Path := basePath + "/bmp2.inp"
	errStep1 := FileUtil.WriteMapToFile(BmpInp1Path, TemplateFilePath1, inp1)
	errStep2 := FileUtil.WriteMapToFile(BmpInp2Path, TemplateFilePath2, inp2)
	if errStep1 != nil {
		log.Println(errStep1)
	}
	if errStep2 != nil {
		log.Println(errStep2)
	}

	// 6. merge inp file
	BmpResultPath := basePath + "/PRZM5.inp"
	FileUtil.MergeInpFile(BmpInp1Path, BmpInp2Path, BmpResultPath)

	// 7. generate wea file
	baselineWeaPath := basePath + "/weatherData.wea"
	WeatherDataUtil.GenerateWeaData(Step2Params, baselineWeaPath)

	// check irrigation type
	if Step2Params.IrrigationType == 0 {
		InputFileGenerateUtil.ChangeWeatherData(Step2Params.IrrigationDate, Step2Params.IrrigationAmount, baselineWeaPath)
	}

	// 8. generate zts file
	curOs := runtime.GOOS
	var cmd *exec.Cmd
	switch curOs {
	case "linux":
		wineLauncher := viper.GetString("wine.launcher")
		cmd = exec.Command(wineLauncher, "./przm5place/PRZM5.exe", basePath)
	case "windows":
		cmd = exec.Command("./przm5place/PRZM5.exe", basePath)
	}
	err := cmd.Run()
	if err != nil {
		log.Println("[!] przm5.exe error", err)
	}
}

// ParseBmpFun : Returns the result using the specified function according to BmpOpts
func ParseBmpFun(
	rec BmpParams.BmpBasicParams,
	inp1 map[string]string,
	inp2 map[string]string,
	gv *GlobalValueUtil.GlobalVar,
) (map[string]string, map[string]string) {
	// get params
	BmpOpts := rec.BmpOpts
	rec1 := rec.Step1Params
	rec2 := rec.Step2Params
	// bmp options
	PAR := rec.PesticideAppReduction
	SC := rec.StripCropping
	PT := rec.ParallelTerracing
	CP := rec.CoverCrops
	CRRM := rec.CropRotationAndResidueManagement

	for _, opt := range BmpOpts {
		switch opt {
		case "contourFarming":
			Record3, Record5, Record9 := ContourFarming(rec1, rec2, gv)
			inp1["3"] = Record3 + "\n"
			inp2["5"] = Record5 + "\n"
			inp2["9"] = Record9 + "\n"
			//log.Println("[BMP] contour farming")
		case "stripCropping":
			Record3, Record5, Record9 := StripCropping(rec1, rec2, gv, SC)
			inp1["3"] = Record3 + "\n"
			inp2["5"] = Record5 + "\n"
			inp2["9"] = Record9 + "\n"
			//log.Println("[BMP] strip cropping")
		case "parallelTerracing":
			Record3, Record5, Record9 := ParallelTerracing(rec1, rec2, gv, PT)
			inp1["3"] = Record3 + "\n"
			inp2["5"] = Record5 + "\n"
			inp2["9"] = Record9 + "\n"
			//log.Println("[BMP] parallel terracing")
		case "cropRotation":
			Record5 := CropRotation(rec2, gv, CRRM)
			inp2["5"] = Record5 + "\n"
		case "coverCrops":
			Record5 := CoverCrops(rec2, gv, CP)
			inp2["5"] = Record5 + "\n"
			//log.Println("[BMP] cover crops")
		case "pesticideAppReduction":
			RecordC2 := PesticideAppReduction(rec2, PAR)
			inp2["C2"] = RecordC2 + "\n"
			//log.Println("[BMP] pesticide application reduction")
		case "pesticideAppTiming":
			RecordC1 := PesticideAppTiming(rec2)
			inp2["C1"] = RecordC1 + "\n"
			//log.Println("[BMP] pesticide application timing")
		}
	}
	return inp1, inp2
}
