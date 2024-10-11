package FileUtil

import (
	"github.com/spf13/viper"
	"log"
	"os/exec"
	"runtime"
)

func DoExecute(exeFi, arg string) {
	curOs := runtime.GOOS // 获取当前系统

	var cmd *exec.Cmd

	// 根据系统执行不同的命令
	switch curOs {
	case "linux":
		wineLauncher := viper.GetString("wine.launcher")
		cmd = exec.Command(wineLauncher, exeFi, arg)
	case "windows":
		cmd = exec.Command(exeFi, arg)
	}
	err := cmd.Run()
	if err != nil {
		log.Println("[DoExecute] run", exeFi, " error:", err)
	}
}
