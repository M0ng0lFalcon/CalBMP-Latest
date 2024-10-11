package Database

import (
	"calbmp-back/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	host := viper.GetString("datasource.host")
	user := viper.GetString("datasource.user")
	password := viper.GetString("datasource.password")
	dbname := viper.GetString("datasource.dbname")
	port := viper.GetString("datasource.port")
	sslMode := viper.GetString("datasource.sslMode")
	TimeZone := viper.GetString("datasource.TimeZone")

	ConnectTemplate := "host=%v user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
	args := fmt.Sprintf(ConnectTemplate,
		host,
		user,
		password,
		dbname,
		port,
		sslMode,
		TimeZone)
	db, err := gorm.Open(postgres.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	// auto migrate
	ManualAutoMigrateFun(db)

	DB = db
}

func ManualAutoMigrateFun(db *gorm.DB) {
	// auto migrate database
	_ = db.AutoMigrate(&model.User{})
	_ = db.AutoMigrate(&model.History{})
}

func GetDB() *gorm.DB {
	return DB
}
