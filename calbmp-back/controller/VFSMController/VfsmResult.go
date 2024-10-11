package VFSMController

import (
	"calbmp-back/Params/InputParams"
	"calbmp-back/Params/ResultParams"
	"calbmp-back/Repository/CropPesticideFinalRepository"
	"calbmp-back/Res"
	"calbmp-back/service/ResultService"
	"calbmp-back/service/VFSMService"
	"calbmp-back/util/JsonUtil"
	"calbmp-back/util/StringUtil"
	"github.com/gin-gonic/gin"
)

// Result2Visualization
//
//	@Description: 获取可进行可视化的VFS结果数据
//	@param ctx
func Result2Visualization(ctx *gin.Context) {
	// 获取基本参数
	CreatedTime := ctx.Query("created_time")
	MinId := ctx.Query("min_id")
	MaxId := ctx.Query("max_id")

	// 解析baseline的数据，获得日期序列
	inp2Path := "./przm5place/" + CreatedTime + "/baseline/step2.json"
	var inp2 InputParams.UserInputStepReceiver
	JsonUtil.ReadJson(inp2Path, &inp2)
	var TextRes ResultParams.TextResultParams
	TextRes.CreatedTime = CreatedTime
	TextRes.PesticideList = inp2.Pesticide // type: []string
	TextRes.ScenarioType = "baseline"
	ztsResult := ResultService.GetTextResultFun(TextRes, "org")
	textZtsJson := "./przm5place/" + CreatedTime + "/vfsm_baseline/textZts.json"
	JsonUtil.WriteJson(textZtsJson, ztsResult)

	// 解析owq文件
	owq := VFSMService.ParseOwq(
		CreatedTime,
		StringUtil.ConvertToInt(MinId),
		StringUtil.ConvertToInt(MaxId),
	)

	// 解析osm文件
	osm := VFSMService.ParseOsm(
		CreatedTime,
		StringUtil.ConvertToInt(MinId),
		StringUtil.ConvertToInt(MaxId),
	)

	// 初始化返回数据
	runoff := make([]float64, 0)
	sediment := make([]float64, 0)
	concentration := make([]float64, 0)
	pesticide_liquid := make([]float64, 0)
	pesticide_solid := make([]float64, 0)

	// runoff单位换算, m3 -> acre-feet
	for _, v := range osm["runoff"] {
		runoff = append(runoff, v/1233.48186)
	}

	// sediment 单位换算, g -> lbs
	for _, v := range osm["sediment"] {
		sediment = append(sediment, v/453.59702440351987)
	}

	// concentration 单位换算
	for i, v := range owq["PesticideOutflowInLiquidPhase"] {
		concentrationVal := 0.0
		if osm["runoff"][i] != 0 {
			concentrationVal = v / osm["runoff"][i]
		} else {
			concentrationVal = 0
		}
		concentration = append(concentration, concentrationVal)
	}

	// pesticide_liquid
	for _, v := range owq["PesticideOutflowInLiquidPhase"] {
		temp := v
		pesticide_liquid = append(pesticide_liquid, temp)
	}

	// pesticide_solid
	for _, v := range owq["PesticideOutflowInSolidPhase"] {
		temp := v
		pesticide_solid = append(pesticide_solid, temp)
	}

	Res.Success(ctx, gin.H{
		"result": map[string]interface{}{
			"date": ztsResult.Date,
			"water": map[string]interface{}{
				"RUNF": runoff,
			},
			"sediment": map[string]interface{}{
				"ESLS": sediment,
			},
			"pesticide": map[string]interface{}{
				"RFLX_" + inp2.Pesticide[0] + "_1_TSER": pesticide_liquid,
				"EFLX_" + inp2.Pesticide[0] + "_1_TSER": pesticide_solid,
			},
			"concentration": map[string]interface{}{
				inp2.Pesticide[0] + "_1": concentration,
			},
			"benchmark": map[string]interface{}{
				inp2.Pesticide[0] + "_1": CropPesticideFinalRepository.FindBenchMarkValueByChemicalName(inp2.Pesticide[0]) * 1000,
			},
		},
	}, "test")

}

//func SummaryData(ctx *gin.Context) {
//	// 获取基本参数
//	CreatedTime := ctx.Query("created_time")
//	MinId := ctx.Query("min_id")
//	MaxId := ctx.Query("max_id")
//
//	// 解析baseline的数据，获得日期序列
//	inp2Path := "./przm5place/" + CreatedTime + "/baseline/step2.json"
//	var inp2 InputParams.UserInputStepReceiver
//	JsonUtil.ReadJson(inp2Path, &inp2)
//	var TextRes ResultParams.TextResultParams
//	TextRes.CreatedTime = CreatedTime
//	TextRes.PesticideList = inp2.Pesticide // type: []string
//	TextRes.ScenarioType = "baseline"
//	ztsResult := ResultService.GetTextResultFun(TextRes, "org")
//	textZtsJson := "./przm5place/" + CreatedTime + "/vfsm_baseline/textZts.json"
//	JsonUtil.WriteJson(textZtsJson, ztsResult)
//
//	// 解析owq文件
//	owq := VFSMService.ParseOwq(
//		CreatedTime,
//		StringUtil.ConvertToInt(MinId),
//		StringUtil.ConvertToInt(MaxId),
//	)
//
//	// 解析osm文件
//	osm := VFSMService.ParseOsm(
//		CreatedTime,
//		StringUtil.ConvertToInt(MinId),
//		StringUtil.ConvertToInt(MaxId),
//	)
//
//}

func ComparisonData(ctx *gin.Context) {
	res := map[string]float64{
		"comparison_erosion":        10.12,
		"comparison_runoff":         12.12,
		"comparison_volatilization": 0,
	}
	Res.Success(ctx, gin.H{"res": res}, "test")
}
