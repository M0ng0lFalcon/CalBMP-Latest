package WeatherDataUtil

import (
	"calbmp-back/Database"
	"calbmp-back/Params/InputParams"
	"calbmp-back/model"
	"calbmp-back/util/FileUtil"
	"fmt"
	"log"
	"strings"
)

func GenerateWeaData(rec InputParams.UserInputStepReceiver, filePath string) {
	// get params
	startYear := strings.Split(rec.Emergence, "-")[0]
	stopYear := strings.Split(rec.Harvest, "-")[0]
	ZipCode := rec.ZipCode
	Log := rec.Log
	Lat := rec.Lat

	if startYear == stopYear {
		res := getWeaDataByYear(startYear, ZipCode, Log, Lat)
		err := FileUtil.WriteToFile(filePath, res)
		if err != nil {
			log.Println("wea file write err:", err)
		}
	} else {
		res1 := getWeaDataByYear(startYear, ZipCode, Log, Lat)
		res2 := getWeaDataByYear(stopYear, ZipCode, Log, Lat)
		res1 = append(res1, res2...)
		_ = FileUtil.WriteToFile(filePath, res1)
	}

}

func getWeaDataByYear(Year string, ZipCode string, Log float64, Lat float64) []string {
	var res []string
	sqlTemp := "select * from \"weaData%s\" where \"ZIP_CODE\" = '%s' and log=%f and lat=%f order by \"Year\",\"Month\",\"Day\""
	weaSql := fmt.Sprintf(sqlTemp, Year, ZipCode, Log, Lat)
	//fmt.Println("[S] query wea data sql: ", weaSql)
	db := Database.GetDB()
	var weaData []model.WeaData
	db.Raw(weaSql).Scan(&weaData)

	temp := "%02d,%02d,%d,%s,%s,%s,%s,%s\n"

	// weather data
	for _, v := range weaData {
		tmp := fmt.Sprintf(temp,
			v.Month,
			v.Day,
			v.Year,
			v.Precip,
			v.Evapotranspire,
			v.Temperature,
			v.Velocity,
			v.SolarRadiation,
		)
		res = append(res, tmp)
	}
	return res
}
