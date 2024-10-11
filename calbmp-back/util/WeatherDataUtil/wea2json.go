package WeatherDataUtil

import (
	"bufio"
	"calbmp-back/util/FileUtil"
	"calbmp-back/util/JsonUtil"
	"fmt"
	"log"
	"strings"
)

func Wea2json(weaPath string, jsonPath string) {
	weaFi := FileUtil.OpenFileAsRead(weaPath)

	scanner := bufio.NewScanner(weaFi)
	res := make(map[string]map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		li := strings.Split(line, ",")
		// write year
		year := fmt.Sprintf("%s/%s/%s", li[2], li[0], li[1])
		prec := fmt.Sprintf("%s", li[3])
		Evapotranspire := fmt.Sprintf("%s", li[4])
		Temperature := fmt.Sprintf("%s", li[5])
		Velocity := fmt.Sprintf("%s", li[6])
		SolarRadiation := fmt.Sprintf("%s", li[7])
		res[year] = map[string]string{
			"prec":           prec,
			"Evapotranspire": Evapotranspire,
			"Temperature":    Temperature,
			"Velocity":       Velocity,
			"SolarRadiation": SolarRadiation,
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	JsonUtil.WriteJson(jsonPath, res)
}
