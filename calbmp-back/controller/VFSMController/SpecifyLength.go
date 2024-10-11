package VFSMController

import (
	"calbmp-back/Params/VFSMParams"
	"calbmp-back/Res"
	"calbmp-back/service/VFSMService"
	"calbmp-back/util/FileUtil"
	"calbmp-back/util/JsonUtil"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os/exec"
	"path/filepath"
	"runtime"
)

func SpecifyLengthController(ctx *gin.Context) {
	var param VFSMParams.SpecifyLengthParam
	ctx.ShouldBindJSON(&param)
	basePath := "./inputs/"

	ikwJsonFilename := fmt.Sprintf("Ikw_%d.json", param.VfsmId)
	jsonPath := filepath.Join(basePath, param.CreatedTime, "json", ikwJsonFilename)

	var ikw VFSMService.IkwStruct
	JsonUtil.ReadJson(jsonPath, &ikw)

	for i := param.Lower; i <= param.Upper; i += param.Increment {
		ikw.VL = float64(i)
		res := ikw.JoinLine()
		//numStr := fmt.Sprintf("%d", i)
		writePath := fmt.Sprintf("%s/%s_%d_storm_len%d.%s", basePath, param.CreatedTime, param.VfsmId, i, "ikw")
		//writePath := filepath.Join(basePath, param.CreatedTime, "inputs", "length_"+param.VfsmId+"_"+numStr+".ikw")
		errWrite := FileUtil.WriteToFile(writePath, []string{res})
		if errWrite != nil {
			log.Println("[GenerateInputFile] -> write inp file, err:", errWrite)
		}

		baseFileSuffix := []string{"igr", "irn", "iro", "isd", "iso", "iwq"}
		for _, suffix := range baseFileSuffix {
			srcFilename := fmt.Sprintf("%s/%s_%d_storm.%s", basePath, param.CreatedTime, param.VfsmId, suffix)
			dstFilename := fmt.Sprintf("%s/%s_%d_storm_len%d.%s", basePath, param.CreatedTime, param.VfsmId, i, suffix)
			_, err := FileUtil.CopyFile(srcFilename, dstFilename)
			if err != nil {
				log.Println("[Storm] copy file error:", err)
			}
		}

		// run przm5.exe
		curOs := runtime.GOOS
		var cmd *exec.Cmd
		exeFi := "./vfsm.exe" // VFS Model path
		target := fmt.Sprintf("%s_%d_storm_len%d", param.CreatedTime, param.VfsmId, i)
		switch curOs {
		case "linux":
			wineLauncher := viper.GetString("wine.launcher")
			cmd = exec.Command(wineLauncher, exeFi, target)
		case "windows":
			cmd = exec.Command(exeFi, target)
		}
		err := cmd.Run()
		if err != nil {
			log.Println("[!] vfsm.exe error: ", target, err)
		}
	}

	Res.Success(ctx, gin.H{"test": ikw}, "test")

}
