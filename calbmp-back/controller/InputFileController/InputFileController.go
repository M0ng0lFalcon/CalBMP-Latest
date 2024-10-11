package InputFileController

import (
	"calbmp-back/Params/InputParams"
	"calbmp-back/RedisUtil"
	"calbmp-back/Res"
	"calbmp-back/service/InputService"
	"calbmp-back/service/ResultService"
	"calbmp-back/util/FileUtil"
	"calbmp-back/util/GlobalValueUtil"
	"calbmp-back/util/InputFileGenerateUtil"
	"calbmp-back/util/WeatherDataUtil"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func UserInputStep1(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	// set current time
	currentTimeBase := time.Now()
	curTime := currentTimeBase.Format("2006-01-02-15-04-05")
	RedisUtil.SetStringKey(username.(string)+"_curTime", curTime)
	// path of template
	TemplateFilePath := "./public/przm5Template/A1_4.txt"
	// get params
	var rec InputParams.UserInputStepReceiver
	errBind := ctx.BindJSON(&rec)
	if errBind != nil {
		log.Fatalln(errBind)
	}

	// set some param
	rec.Username = username.(string)
	rec.CurTime = curTime

	// main fun
	var gv GlobalValueUtil.GlobalVar
	resMap, _ := InputService.RepPlaceHolder2Fun(TemplateFilePath, rec, &gv, username.(string))

	// save res to file
	FileUtil.Mkdir("./przm5place/" + curTime)
	FileUtil.Mkdir("./przm5place/" + curTime + "/baseline")
	FileUtil.Mkdir("./przm5place/" + curTime + "/bmp")
	inp1Path := "./przm5place/" + curTime + "/baseline/baseline1.inp"
	err := FileUtil.WriteMapToFile(inp1Path, TemplateFilePath, resMap)

	// save step1 struct to file
	jsonString, _ := json.Marshal(rec)
	step1path := "./przm5place/" + curTime + "/baseline/step1.json"
	FileUtil.WriteString2File(step1path, string(jsonString))

	// response message
	msg := "False"
	if err == nil {
		msg = "True"
	}
	Res.SuccessMsg(ctx, msg)
}

func UserInputChemical(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	TemplateFilePath := "./public/przm5Template/5_U2.txt"
	// get params
	var rec InputParams.UserInputStepReceiver
	errBind := ctx.BindJSON(&rec)
	if errBind != nil {
		log.Fatalln(errBind)
	}

	// save some var to redis
	RedisUtil.SetStringKey(username.(string)+"_crop", rec.Crop)
	RedisUtil.SetStringKey(username.(string)+"_cokey", rec.CoKey)
	// set some param
	curTime := RedisUtil.GetStringVal(username.(string) + "_curTime")
	rec.Username = username.(string)
	rec.CurTime = curTime

	// main fun
	gv := GlobalValueUtil.GetGlobalVariables(rec)
	resMap, _ := InputService.RepPlaceHolder2Fun(TemplateFilePath, rec, &gv, username.(string))

	// save res to file
	inp2Path := "./przm5place/" + curTime + "/baseline/baseline2.inp"
	err := FileUtil.WriteMapToFile(inp2Path, TemplateFilePath, resMap)

	// merge test1.inp and test2.inp
	inp1Path := "./przm5place/" + curTime + "/baseline/baseline1.inp"
	targetPath := "./przm5place/" + curTime + "/baseline/PRZM5.inp"
	FileUtil.MergeInpFile(inp1Path, inp2Path, targetPath)

	// generate wea file
	baselineWeaPath := "./przm5place/" + curTime + "/baseline/weatherData.wea"
	WeatherDataUtil.GenerateWeaData(rec, baselineWeaPath)

	// check irrigation type
	if rec.IrrigationType == 0 {
		InputFileGenerateUtil.ChangeWeatherData(rec.IrrigationDate, rec.IrrigationAmount, baselineWeaPath)
	}

	// save step2 struct to file
	jsonString, _ := json.Marshal(rec)
	step2path := "./przm5place/" + curTime + "/baseline/step2.json"
	FileUtil.WriteString2File(step2path, string(jsonString))

	// response message
	msg := "False"
	if err == nil {
		msg = "True"
	}

	// run przm5.exe
	curOs := runtime.GOOS
	var cmd *exec.Cmd
	switch curOs {
	case "linux":
		wineLauncher := viper.GetString("wine.launcher")
		cmd = exec.Command(wineLauncher, "./przm5place/PRZM5.exe", "./przm5place/"+curTime+"/baseline")
	case "windows":
		cmd = exec.Command("./przm5place/PRZM5.exe", "./przm5place/"+curTime+"/baseline")
	}
	err = cmd.Run()
	if err != nil {
		log.Println("[!] przm5.exe error: ", err)
		Res.SuccessMsg(ctx, "przm5 error")
	} else {
		ResultService.Zts2Json(
			"./przm5place/"+curTime+"/baseline/baseline.zts",
			"./przm5place/"+curTime+"/baseline/baselineZts.json",
		)
	}

	Res.Success(ctx, gin.H{"created_time": curTime}, msg)
}

func GetStepJson(ctx *gin.Context) {
	var params InputParams.GetJsonParams
	err := ctx.BindJSON(&params)
	if err != nil {
		log.Println("[GetStepJson] :", err)
		Res.FailMsg(ctx, "[GetStepJson] err")
		return
	}

	filename := fmt.Sprintf("./przm5place/%s/baseline/step%d.json", params.CreatedAt, params.Step)
	content, err := os.ReadFile(filename)
	Res.Success(ctx, gin.H{
		"json_data": content,
	}, "[get step json]")

}
