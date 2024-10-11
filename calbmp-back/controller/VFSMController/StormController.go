package VFSMController

import (
	"calbmp-back/Params/InputParams"
	"calbmp-back/Params/VFSMParams"
	"calbmp-back/RedisUtil"
	"calbmp-back/Repository/ChorizonRepository"
	"calbmp-back/Repository/SoilTextureFinalRepository"
	"calbmp-back/Res"
	"calbmp-back/dto/ResultDataDTO"
	"calbmp-back/service/VFSMService"
	"calbmp-back/util/FileUtil"
	"calbmp-back/util/JsonUtil"
	"calbmp-back/util/StringUtil"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"math"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func Storm(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	var params VFSMParams.DesignParam
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		Res.FailMsg(ctx, "[GenerateInputFile] bind json error")
		return
	}

	// get basic data
	slope := RedisUtil.GetFloatVal(username.(string) + "_slope") // !!!
	slope *= 0.01
	slopeLength := RedisUtil.GetFloatVal(username.(string) + "_slopeLength") // !!!
	uslep := RedisUtil.GetFloatVal(username.(string) + "_uslep")             // !!!
	fieldSize := RedisUtil.GetFloatVal(username.(string) + "_fieldSize")     // !!!
	//fmt.Println(slopeLength, slope, uslep)
	DateLiWithCNJson := RedisUtil.GetStringVal(username.(string) + "_DateLiWithCN")
	DataLiWithCN := make([][]int, 0) // !!!
	errUnmarshal := json.Unmarshal([]byte(DateLiWithCNJson), &DataLiWithCN)
	if errUnmarshal != nil {
		log.Println("json string to int array, errUnmarshal:", errUnmarshal)
	}
	curTime := RedisUtil.GetStringVal(username.(string) + "_curTime") // !!!
	record5 := RedisUtil.GetStringVal(username.(string) + "_record5")

	// parse baseline input file
	inp1Path := "./przm5place/" + curTime + "/baseline/step1.json"
	inp2Path := "./przm5place/" + curTime + "/baseline/step2.json"
	var inp1 InputParams.UserInputStepReceiver
	var inp2 InputParams.UserInputStepReceiver
	JsonUtil.ReadJson(inp1Path, &inp1)
	JsonUtil.ReadJson(inp2Path, &inp2)

	// get pesticide date
	vfsmPath := "./przm5place/" + curTime + "/vfsm_baseline/"
	pesticideDate, prec, vfsmId := getPesticideDate(vfsmPath + "textZts.json") // !!!

	// get cn & c
	cn := getCN(pesticideDate, DataLiWithCN) // !!!
	c := getCfactor(pesticideDate, record5)  // !!!

	// get soil
	SoilTexture := SoilTextureFinalRepository.FindByMukeyAndCokey(
		inp1.MuKey,
		inp2.CoKey,
	)
	Chorizon := ChorizonRepository.FindChorizonListByMukeyAndCokey(
		inp1.MuKey,
		inp2.CoKey,
	)[0]

	params.P = prec
	params.CN = cn
	params.A = fieldSize
	params.L = slopeLength
	params.Y = slope
	params.SoilType = SoilTexture.Texcl
	params.K = 0.1317 * (0.02606*(12-StringUtil.Convert2Float(Chorizon.Om)) + 0.0650 + 0.05)
	params.C = c
	params.P2 = uslep
	params.SoilOrgMatter = Chorizon.Om

	res := VFSMService.GenerateInp(params)

	rootDir := "./inputs"
	targetDir := filepath.Join(rootDir)
	outFilename := fmt.Sprintf("%s/%s_%d_storm.%s", targetDir, curTime, vfsmId, "inp")
	errWrite := FileUtil.WriteToFile(outFilename, []string{res})
	if errWrite != nil {
		log.Println("[GenerateInputFile] -> write inp file, err:", errWrite)
	}

	// copy basic input files
	baseFileSuffix := []string{"igr", "ikw", "irn", "iro", "isd", "iso", "iwq"}
	for _, suffix := range baseFileSuffix {
		srcFilename := fmt.Sprintf("%s/%s_%d.%s", targetDir, curTime, vfsmId, suffix)
		dstFilename := fmt.Sprintf("%s/%s_%d_storm.%s", targetDir, curTime, vfsmId, suffix)
		_, err := FileUtil.CopyFile(srcFilename, dstFilename)
		if err != nil {
			log.Println("[Storm] copy file error:", err)
		}
	}

	// run przm5.exe
	curOs := runtime.GOOS
	var cmd *exec.Cmd
	exeFi := "./uh.exe" // VFS Model uh path
	target := fmt.Sprintf("%s_%d_storm", curTime, vfsmId)
	switch curOs {
	case "linux":
		wineLauncher := viper.GetString("wine.launcher")
		cmd = exec.Command(wineLauncher, exeFi, target)
	case "windows":
		cmd = exec.Command(exeFi, target)
	}
	err = cmd.Run()
	if err != nil {
		log.Println("[!] vfsm.exe error: ", err)
		Res.SuccessMsg(ctx, "vfsm error")
	}

	exeFi = "./vfsm.exe" // VFS Model path
	target = fmt.Sprintf("%s_%d_storm", curTime, vfsmId)
	switch curOs {
	case "linux":
		wineLauncher := viper.GetString("wine.launcher")
		cmd = exec.Command(wineLauncher, exeFi, target)
	case "windows":
		cmd = exec.Command(exeFi, target)
	}
	err = cmd.Run()
	if err != nil {
		log.Println("[!] vfsm.exe error: ", target, err)
	}

	Res.Success(ctx, gin.H{"res": res, "soil_type": SoilTexture.Texcl}, "test")
}

func getPesticideDate(textZtsPath string) (pesticideDate time.Time, prec float64, vfsm_id int) {
	var textRes ResultDataDTO.ResData
	JsonUtil.ReadJson(textZtsPath, &textRes)
	minn := 9999

	concentration := textRes.Concentration
	for key := range concentration {
		for i, v := range concentration[key] {
			if v != 0.0 {
				minn = int(math.Min(float64(minn), float64(i)))
				break
			}
		}
	}
	if minn == -1 {
		minn = 0
	}
	resDate, _ := time.Parse("2006/1/2", textRes.Date[minn])
	return resDate, textRes.Water["PRCP"][minn], minn + 1
}

func getCN(targetDate time.Time, cnLi [][]int) (cn int) {
	for _, v := range cnLi {
		d, _ := time.Parse("2006/1/2", fmt.Sprintf("%d/%d/%d", v[2], v[1], v[0]))
		if targetDate.Before(d) {
			//if i+1 != len(cnLi) {
			//	return (v[3] + cnLi[i+1][3]) / 2
			//} else {
			//	return v[3]
			//}
			return v[3]
		}
	}
	return 0
}

func getCfactor(targetDate time.Time, record5 string) float64 {
	lineLi := strings.Split(record5, "\n")
	for i, line := range lineLi {
		c, d := parseRecord5line(line)
		if targetDate.Before(d) {
			if i+1 != len(lineLi) {
				c2, _ := parseRecord5line(lineLi[i+1])
				return (c + c2) / 2
			} else {
				return c
			}
		}
	}
	return 0
}

func parseRecord5line(line string) (c float64, d time.Time) {
	li := strings.Split(line, ",")
	c = StringUtil.Convert2Float(li[3])
	year := StringUtil.ConvertToInt(li[2])
	month := StringUtil.ConvertToInt(li[1])
	day := StringUtil.ConvertToInt(li[0])
	d, _ = time.Parse("2006/1/2", fmt.Sprintf("%d/%d/%d", year, month, day))
	return
}
