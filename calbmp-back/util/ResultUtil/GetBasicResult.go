package ResultUtil

import (
	"bufio"
	"calbmp-back/util/FileUtil"
	"strings"
)

// GetBasicResult : read data from zts file
func GetBasicResult(
	ScenarioType string,
	BmpId string,
	CreatedTime string,
) ([]string, [][]string) {
	// open file
	var resultFilePath string
	// open specific file according to the scenario type
	// ScenarioType == baseline -> open baseline zts file
	// ScenarioType == bmp      -> open bmp zts file by BmpId
	if ScenarioType == "baseline" {
		resultFilePath = "./przm5place/" + CreatedTime + "/baseline/baseline.zts"
	} else {
		resultFilePath = "./przm5place/" + CreatedTime + "/bmp/" + BmpId + "/bmp.zts"
	}
	// return file obj according to the file path
	resFile := FileUtil.OpenFileAsRead(resultFilePath)

	// !!! read file line by line
	columnNames := make([]string, 0)
	// !!! struct of dataLi, return data:
	// dataLi[line num of data][column index]
	dataLi := make([][]string, 0)
	lineNum := 0
	scanner := bufio.NewScanner(resFile)
	for scanner.Scan() {
		// if line number == 0, representing the title of zts file
		// if line number == 1, representing the empty line
		if lineNum == 0 || lineNum == 1 {
			// continue when line number < 2
		} else if lineNum == 2 {
			// line number == 2, get the column name
			columnNames = append(columnNames, strings.Fields(scanner.Text())...)
		} else if lineNum > 2 {
			// line number > 2 : get data line by line
			tmpData := strings.Fields(scanner.Text())
			dataLi = append(dataLi, tmpData)
		}
		lineNum++
	}

	return columnNames, dataLi
}
