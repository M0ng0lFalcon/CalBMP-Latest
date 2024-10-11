package InputFileGenerateUtil

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ChangeWeatherData(IrrigationDate []string, Amount []float64, path string) {
	dst := strings.Replace(path, ".wea", ".bak", -1)
	errRename := os.Rename(path, dst)
	if errRename != nil {
		log.Println("[ChangeWeatherData] org wea file rename file:", errRename)
		return
	}

	file, err := os.OpenFile(dst, os.O_RDWR, 066)
	if err != nil {
		log.Println("[ChangeWeatherData] :", err)
		return
	}
	dstFIle, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		log.Println("[ChangeWeatherData] :", err)
		return
	}

	reader := bufio.NewReader(file)
	curIdx := 0
	for {
		// read per line
		line, errRead := reader.ReadString('\n')
		if errRead != nil {
			if errRead == io.EOF {
				break
			} else {
				log.Println("[ChangeWeatherData] :", err)
				return
			}
		}

		// check date
		splitList := strings.Split(line, ",")
		orgDate := fmt.Sprintf("%s-%s-%s", splitList[2], splitList[0], splitList[1])
		if curIdx != len(IrrigationDate) && orgDate == IrrigationDate[curIdx] {
			log.Println("[change weather data]")
			precipitation, errParse := strconv.ParseFloat(splitList[3], 32)
			if errParse != nil {
				return
			}
			//precipitation += Amount[curIdx] * 2.54
			precipitation += Amount[curIdx]
			splitList[3] = fmt.Sprintf("%v", precipitation)

			newLine := strings.Join(splitList, ",")
			_, errWrite := dstFIle.WriteString(newLine)
			if errWrite != nil {
				log.Println("[ChangeWeatherData]: err write:", errWrite)
				return
			}
			curIdx += 1
		} else {
			_, errWrite := dstFIle.WriteString(line)
			if errWrite != nil {
				log.Println("[ChangeWeatherData]: err write:", errWrite)
				return
			}
		}
	}
	errRemove := os.Remove(dst)
	if errRemove != nil {
		log.Println("[ChangeWeatherData] remove file:", errRemove)
		return
	}
}
