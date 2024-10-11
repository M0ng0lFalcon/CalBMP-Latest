package StringUtil

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func FormatDate(dateStr string) string {
	return strings.Replace(dateStr, "-", ",", -1)
}

// ReLineCmd : second step
func ReLineCmd(str string) string {
	str = strings.Replace(str, "{", "", -1)
	str = strings.Replace(str, "}", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, "get", "", -1)
	str = strings.Replace(str, "Record", "", -1)
	return str
}

func DeleteFrontZero(ss string) string {
	res := ""
	idx := 0
	for i, v := range ss {
		if v != '0' {
			idx = i
			break
		}
	}
	res = ss[idx:]
	return res
}

func ConvertToInt(NumStr string) int {
	NumStr = strings.Trim(NumStr, " ")
	Num, err := strconv.Atoi(NumStr)
	if err != nil {
		log.Println("[!] Convert", NumStr, " Error :", err)
	}
	return Num
}

func Convert2Float(NumStr string) float64 {
	NumStr = strings.Trim(NumStr, " ")
	Num, err := strconv.ParseFloat(NumStr, 64)
	if err != nil {
		//log.Println("[!] Convert: ", NumStr, "; Error :", err)
		return 0.0
	}
	return Num
}

func Float2ScientificNotation(FloatNum float64) string {
	res := fmt.Sprintf("%.2e", FloatNum)

	return res
}

// Var2Line : multi var make line
func Var2Line(valLi ...interface{}) (res string) {
	strLi := make([]string, 0)
	for _, val := range valLi {
		tempLine := ""
		switch val.(type) {
		case string:
			tempLine = val.(string)
		case int:
			tempLine = fmt.Sprintf("%d", val.(int))
		case int8:
			tempLine = fmt.Sprintf("%d", val.(int8))
		case int16:
			tempLine = fmt.Sprintf("%d", val.(int16))
		case int32:
			tempLine = fmt.Sprintf("%d", val.(int32))
		case int64:
			tempLine = fmt.Sprintf("%d", val.(int64))
		case uint:
			tempLine = fmt.Sprintf("%d", val.(uint))
		case uint8:
			tempLine = fmt.Sprintf("%d", val.(uint8))
		case uint16:
			tempLine = fmt.Sprintf("%d", val.(uint16))
		case uint32:
			tempLine = fmt.Sprintf("%d", val.(uint32))
		case uint64:
			tempLine = fmt.Sprintf("%d", val.(uint64))
		// float number
		case float32:
			tempLine = fmt.Sprintf("%f", val.(float32))
		case float64:
			tempLine = fmt.Sprintf("%f", val.(float64))
		}
		strLi = append(strLi, tempLine)
	}
	res = strings.Join(strLi, " ")
	return res
}
