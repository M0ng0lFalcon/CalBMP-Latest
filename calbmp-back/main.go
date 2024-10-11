package main

import (
	"calbmp-back/Database"
	"calbmp-back/RedisUtil"
	"calbmp-back/Router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	r := gin.Default()
	// init router
	r = Router.CollectRoute(r)
	// init database
	Database.InitDB()
	// init redis clint
	RedisUtil.InitRedisConnector()
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
