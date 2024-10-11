package VFSMController

import (
	"calbmp-back/Repository/VegetationRepository"
	"calbmp-back/Res"
	"calbmp-back/service/VFSMService"
	"calbmp-back/util/FileUtil"
	"calbmp-back/util/StringUtil"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func GetVegetation(ctx *gin.Context) {
	vegetationList := VegetationRepository.FindAllList()
	data := make([]map[string]string, 0)
	for _, v := range vegetationList {
		temp := map[string]string{
			"value": v,
			"text":  v,
		}
		data = append(data, temp)
	}
	Res.Success(ctx, gin.H{"vegetation_list": data}, "get vegetation")
}

func GetVegetationByName(ctx *gin.Context) {
	vegetationName := ctx.Query("vegetation")
	vegetation := VegetationRepository.FindByVegetation(vegetationName)
	Res.Success(ctx, gin.H{"vegetation": vegetation}, "get vegetation")
}

func GetBasicResult(ctx *gin.Context) {
	//username, _ := ctx.Get("username")
	CreatedTime := ctx.Query("created_time")
	vfsmId := ctx.Query("vfsm_id")

	targetDir := "/output/"
	readLi := []string{"osm", "owq"} // file type waiting for get
	res := make(map[string]string)   // result
	for _, v := range readLi {
		filename := targetDir + CreatedTime + "_" + vfsmId + "." + v
		var lines []string
		err := FileUtil.ReadFile(filename, &lines)
		if err != nil {
			log.Println(err)
		}

		for _, line := range lines {
			if strings.HasPrefix(line, "Trapping efficiency") {
				li := strings.Fields(line)
				num := li[len(li)-2]
				res["Trapping efficiency"] = num
			} else if strings.HasSuffix(line, "Sediment reduction (dE)") {
				li := strings.Fields(line)
				num := li[0]
				res["Sediment reduction (dE)"] = num
			} else if strings.HasSuffix(line, "Runoff inflow reduction") {
				li := strings.Fields(line)
				num := li[0]
				res["Runoff inflow reduction"] = num
			} else if strings.HasSuffix(line, "Pesticide reduction (dP)") {
				li := strings.Fields(line)
				num := li[0]
				res["Pesticide reduction (dP)"] = num
			}
		}
	}
	Res.Success(ctx, gin.H{"res": res}, "vfsm res")
}

// GetPesticideReductionEff
//
//	@Description: 获取减少率
//	@param ctx
func GetPesticideReductionEff(ctx *gin.Context) {
	CreatedTime := ctx.Query("created_time")
	MinId := ctx.Query("min_id")
	MaxId := ctx.Query("max_id")

	owq := VFSMService.ParseOwq(
		CreatedTime,
		StringUtil.ConvertToInt(MinId),
		StringUtil.ConvertToInt(MaxId),
	)

	osm := VFSMService.ParseOsm(
		CreatedTime,
		StringUtil.ConvertToInt(MinId),
		StringUtil.ConvertToInt(MaxId),
	)

	Res.Success(ctx, gin.H{
		"owq": owq,
		"osm": osm,
	}, "test")
}
