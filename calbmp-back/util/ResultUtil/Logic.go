package ResultUtil

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetDateSet(res map[string][]float64) []string {
	dateSet := make([]string, 0)
	// get date data set
	YearSet := res["Year"]
	MonthSet := res["Mo"]
	DaySet := res["Dy"]
	for i, v := range YearSet {
		dateStr := fmt.Sprintf("%.0f/%.0f/%.0f", v, MonthSet[i], DaySet[i])
		dateSet = append(dateSet, dateStr)
	}
	return dateSet
}

func LocateZtsPath(scenarioType, createTime, bmpId string) (ztsPath string) {
	// open specific file according to the scenario type
	// ScenarioType == baseline -> open baseline zts file
	// ScenarioType == bmp      -> open bmp zts file by BmpId
	if scenarioType == "baseline" {
		ztsPath = "./przm5place/" + createTime + "/baseline/baseline.zts"
	} else if scenarioType == "bmp" {
		ztsPath = "./przm5place/" + createTime + "/bmp/" + bmpId + "/bmp.zts"
	} else if scenarioType == "vfsm" {
		ztsPath = "./przm5place/" + createTime + "/vfsm_baseline/baseline.zts"
	}
	return
}

func openFileAsRead(path string) *os.File {
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		log.Println("[!] Open file as read mode error:", err, "path:", path)
		_ = f.Close()
		return nil
	}
	return f
}

func scientificNotation2Float64(NumStr string) float64 {
	var res float64
	if NumStr == "0.0000E+000" {
		res = 0
	} else if NumStr == "NaN" {
		res = -1 // -1 represent not a number
	} else {
		// convert Scientific counting
		res, _ = strconv.ParseFloat(NumStr, 64)
	}
	return res
}

func GetMatrix(path string) map[string][]float64 {
	f := openFileAsRead(path)

	/*
		!!! return data
		columnNameList : column names in *.zts
		ztsData     : [column name] = [data...]
	*/
	columnNameList := make([]string, 0)
	ztsData := make(map[string][]float64)

	lineNum := 0                   // current line number
	scanner := bufio.NewScanner(f) // scanner for *.zts file
	for scanner.Scan() {
		line := scanner.Text() // scan per line

		// if line number == 0, representing the title of zts file
		// if line number == 1, representing the empty line
		if lineNum == 0 || lineNum == 1 {
			//continue
		} else if lineNum == 2 {
			// line number == 2, get the column name
			columnNameList = strings.Fields(line)
			pesticideList := []string{
				"RFLX",
				"EFLX",
				"FPVL",
				"VFLX",
			}

			pesticideFlag := false
			for i, column := range columnNameList {
				for _, pesticide := range pesticideList {
					if strings.Contains(column, pesticide) {
						if !pesticideFlag {
							columnNameList[i] = column + "_TSER"
							pesticideFlag = true
						} else {
							columnNameList[i] = column + "_TCUM"
							pesticideFlag = false
						}
					}
				}
			}
			for _, column := range columnNameList {
				ztsData[column] = make([]float64, 0)
			}
		} else if lineNum > 2 {
			// split line by space
			lineData := strings.Fields(line)
			for i, columnName := range columnNameList {
				value := scientificNotation2Float64(lineData[i])
				ztsData[columnName] = append(ztsData[columnName], value)
			}
		}
		lineNum++
	}
	return ztsData
}

func GetWaterDataSet(ztsData map[string][]float64, opt []string, dataType string) (res map[string][]float64) {
	// get water data set
	res = make(map[string][]float64)
	key := ""
	for _, v := range opt {
		if v == "SWTR" {
			key = "SWTR1"
		} else {
			// valList := ztsData[v+"0"]
			key = v + "0"
		}
		src := ztsData[key]
		valList := make([]float64, 0)
		valList = append(valList, src[:]...)

		if dataType == "change" {
			// cm to acre-feet, 1cm*8.1071318210885/100=b acre-foot
			for i, val := range valList {
				valList[i] = val * 8.1071318210885 / 100.0
			}
		}
		res[v] = valList
	}
	return
}

func GetPesticideDataSet(ztsData map[string][]float64, pesticideList, opt []string, fieldSize float64, dataType string) (res map[string][]float64) {
	// get pesticide data set
	res = make(map[string][]float64)
	for i, pesticide := range pesticideList {
		for _, v := range opt {
			if v == "VFLX" {
				// generate key
				fpvlTserKey := fmt.Sprintf("%s%d_TSER", "FPVL", i+1)
				fpvlTcumKey := fmt.Sprintf("%s%d_TCUM", "FPVL", i+1)
				vflxTserKey := fmt.Sprintf("%s%d_TSER", "VFLX", i+1)
				vflxTcumKey := fmt.Sprintf("%s%d_TCUM", "VFLX", i+1)
				vflxResTserKey := fmt.Sprintf("%s_%s_%d_TSER", "VFLX", pesticide, i+1)
				vflxResTcumKey := fmt.Sprintf("%s_%s_%d_TCUM", "VFLX", pesticide, i+1)

				// get data
				fpvlTserData_src := ztsData[fpvlTserKey]
				fpvlTcumData_src := ztsData[fpvlTcumKey]
				vflxTserData_src := ztsData[vflxTserKey]
				vflxTcumData_src := ztsData[vflxTcumKey]

				// deep copy
				fpvlTserData := make([]float64, 0)
				fpvlTcumData := make([]float64, 0)
				vflxTserData := make([]float64, 0)
				vflxTcumData := make([]float64, 0)
				fpvlTserData = append(fpvlTserData, fpvlTserData_src[:]...)
				fpvlTcumData = append(fpvlTcumData, fpvlTcumData_src[:]...)
				vflxTserData = append(vflxTserData, vflxTserData_src[:]...)
				vflxTcumData = append(vflxTcumData, vflxTcumData_src[:]...)

				// merge data
				for j := 0; j < len(fpvlTserData); j++ {
					vflxTserData[j] = fpvlTserData[j] + vflxTserData[j]
					vflxTcumData[j] = fpvlTcumData[j] + vflxTcumData[j]
					if dataType == "change" {
						vflxTserData[j] = vflxTserData[j] * 1033.18411
						vflxTcumData[j] = vflxTcumData[j] * 1033.18411
					}

				}
				// add to res
				res[vflxResTserKey] = ztsData[vflxTserKey]
				res[vflxResTcumKey] = ztsData[vflxTcumKey]
			} else {
				keyTser := fmt.Sprintf("%s%d_TSER", v, i+1)
				keyTcum := fmt.Sprintf("%s%d_TCUM", v, i+1)
				ResTserKey := fmt.Sprintf("%s_%s_%d_TSER", v, pesticide, i+1)
				ResTcumKey := fmt.Sprintf("%s_%s_%d_TCUM", v, pesticide, i+1)

				// deep copy
				srcTser := ztsData[keyTser]
				srcTcum := ztsData[keyTcum]
				tempTser := make([]float64, 0)
				tempTcum := make([]float64, 0)
				tempTser = append(tempTser, srcTser[:]...)
				tempTcum = append(tempTcum, srcTcum[:]...)

				if dataType == "change" {
					// change unit
					for j, tempVal := range tempTser {
						tempTser[j] = tempVal * 1033.18411
					}
					for j, tempVal := range tempTcum {
						tempTcum[j] = tempVal * 1033.18411
					}
				}

				// add to res
				res[ResTserKey] = tempTser
				res[ResTcumKey] = tempTcum
			}
		}
	}
	return
}

func GetSedimentDataSet(ztsData map[string][]float64, opt []string, dataType string) (res map[string][]float64) {
	res = make(map[string][]float64)
	for _, v := range opt {
		src := ztsData[v+"0"]
		valList := make([]float64, 0)
		valList = append(valList, src[:]...)

		if dataType == "change" {
			// t to lb, 1t = 2205lb
			for i, val := range valList {
				valList[i] = val * 2205
			}
		}

		res[v] = valList
	}
	return
}

func GetConcentrationDataSet(ztsData map[string][]float64, pesticideList []string, dataType string) (res map[string][]float64) {
	res = make(map[string][]float64)
	runfData := ztsData["RUNF0"]
	rflxDataList := make([][]float64, 0)
	for i := 0; i < len(pesticideList); i++ {
		key := fmt.Sprintf("RFLX%d_TSER", i+1)
		temp := ztsData[key]
		rflxDataList = append(rflxDataList, temp)
	}
	var tmp float64
	for i, pesticide := range pesticideList {
		tempDataLi := make([]float64, 0)
		for j := range runfData {
			if rflxDataList[i][j] == -1 {
				tmp = -1
			} else if runfData[j] > 0 {
				tmp = rflxDataList[i][j] / runfData[j] * 1e6
			} else {
				tmp = rflxDataList[i][j] * 1e6
			}
			tmp *= 1000 // ppm to ug/L, 1ppm = 1000 ug/L
			tempDataLi = append(tempDataLi, tmp)
		}
		key := fmt.Sprintf("%s_%d", pesticide, i+1)
		res[key] = tempDataLi
	}
	return
}
