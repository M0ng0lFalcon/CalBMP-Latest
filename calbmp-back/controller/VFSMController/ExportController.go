package VFSMController

import (
	"calbmp-back/Res"
	"calbmp-back/util/FileUtil"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetInputs(ctx *gin.Context) {
	CreatedTime := ctx.Query("created_time")
	attachment := fmt.Sprintf("attachment; filename=VFS_input_files-%s.zip", CreatedTime)
	ctx.Header("Content-Disposition", attachment)
	root := "./inputs/"
	dest := root + CreatedTime + ".zip"
	filesToZip, err := FileUtil.GetFilesToZip(root, CreatedTime)
	if err != nil {
		fmt.Println("Error:", err)
		Res.FailMsg(ctx, "Error:"+err.Error())
		return
	}
	err = FileUtil.CreateZipFile(dest, filesToZip)
	if err != nil {
		fmt.Println("Error:", err)
		Res.FailMsg(ctx, "Error:"+err.Error())
		return
	}
	ctx.File(dest)
}

func GetOutputs(ctx *gin.Context) {
	CreatedTime := ctx.Query("created_time")
	attachment := fmt.Sprintf("attachment; filename=VFS_output_files-%s.zip", CreatedTime)
	ctx.Header("Content-Disposition", attachment)
	root := "./output/"
	dest := root + CreatedTime + ".zip"
	filesToZip, err := FileUtil.GetFilesToZip(root, CreatedTime)
	if err != nil {
		fmt.Println("Error:", err)
		Res.FailMsg(ctx, "Error:"+err.Error())
		return
	}
	err = FileUtil.CreateZipFile(dest, filesToZip)
	if err != nil {
		fmt.Println("Error:", err)
		Res.FailMsg(ctx, "Error:"+err.Error())
		return
	}
	ctx.File(dest)
}
