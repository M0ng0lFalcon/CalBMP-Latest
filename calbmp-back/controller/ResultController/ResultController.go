package ResultController

import (
	"calbmp-back/Params/ResultParams"
	"calbmp-back/Res"
	"calbmp-back/service/ResultService"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetDataByName : get data by names
// Route : /result/getBasicByName
func GetDataByName(ctx *gin.Context) {
	// get request body
	var rec ResultParams.ResultReceiver
	errBind := ctx.Bind(&rec)
	if errBind != nil {
		panic(errBind)
	}

	// get needed data
	resData := ResultService.GetDataByName(rec, "change")

	// return result
	Res.Success(ctx, gin.H{
		"result": resData,
	}, "[*] Get result data")
}

func GetComparisonData(ctx *gin.Context) {
	// get cnt of bmp
	BmpCnt, _ := strconv.Atoi(ctx.Query("bmp_cnt"))
	HarvestData := ctx.Query("harvest")
	CreatedTime := ctx.Query("created_time")

	ComparisonRunoff,
		ComparisonErosion,
		ComparisonVolatilization := ResultService.GetComparisonData(BmpCnt, HarvestData, CreatedTime)

	Res.Success(ctx, gin.H{
		"comparison_runoff":         ComparisonRunoff,
		"comparison_erosion":        ComparisonErosion,
		"comparison_volatilization": ComparisonVolatilization,
	}, "[*] Get comparison data")
}

func GetTextResult(ctx *gin.Context) {
	var params ResultParams.TextResultParams
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		Res.FailMsg(ctx, "[!] params error")
		return
	}
	params.ScenarioType = "baseline"

	textRes := ResultService.GetTextResultFun(params, "change")
	Res.Success(ctx, gin.H{"text_res": textRes}, "[*] get text outputs")
}

func GetBestBMP(ctx *gin.Context) {

}

func GetInputFiles(ctx *gin.Context) {
	CreatedTime := ctx.Query("created_time")

	files := ResultService.GetInputFilesFun(CreatedTime)

	Res.Success(ctx, gin.H{"input_files": files}, "[*] get input file list")
}

func ZipInputFile(ctx *gin.Context) {
	var param ResultParams.ZipInputFileParam
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		Res.FailMsg(ctx, "[*] params error")
		return
	}
	ResultService.ZipInputFileFun(param)
	Res.SuccessMsg(ctx, "[*] zip successfully")
}

func DownloadInputFile(ctx *gin.Context) {
	CreatedTime := ctx.Query("created_time")
	zipFilePath := "./przm5place/" + CreatedTime + "/down.zip"
	ctx.File(zipFilePath)
}
