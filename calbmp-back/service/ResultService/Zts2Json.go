package ResultService

import (
	"bufio"
	"calbmp-back/util/FileUtil"
	"calbmp-back/util/JsonUtil"
	"fmt"
	"strings"
)

func Zts2Json(ztsPath, outPath string) {
	ztsFi := FileUtil.OpenFileAsRead(ztsPath)

	ztsMap := make(map[string]map[string]string)
	lineNum := 0
	scanner := bufio.NewScanner(ztsFi)
	fieldNames := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		// if line number == 0, representing the title of zts file
		// if line number == 1, representing the empty line
		// if line number == 2, field names
		if lineNum < 2 {
			// continue when line number <= 2
		} else if lineNum == 2 {
			fieldNames = append(fieldNames, strings.Fields(scanner.Text())[3:]...)
		} else if lineNum >= 3 {
			// line number >= 3 : get data line by line
			li := strings.Fields(line)
			tempMap := make(map[string]string)
			date := fmt.Sprintf("%s/%s/%s", li[0], li[1], li[2])
			for i, v := range fieldNames {
				tempMap[v] = li[i+3]
			}
			ztsMap[date] = tempMap
		}
		lineNum++
	}
	JsonUtil.WriteJson(outPath, ztsMap)
}
