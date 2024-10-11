package ResultUtil

import "strconv"

// GetDataByIndex : get data by index from basic data list
func GetDataByIndex(
	dataLi [][]string,
	indexLi []int,
) (res [][]float64) {
	// 1. check concentration mode
	if len(indexLi) > 0 && indexLi[0] == -1 {
		return GetConcentrationData(dataLi)
	}

	// 2. else : get data by normal mode
	// traverse index list
	for _, idx := range indexLi {
		// get string data by index
		basicNow := GetStringColumnData(idx, dataLi)

		// validate data : is NaN exist in basicNow
		flag := validateNaN(basicNow)
		// if not exist
		if flag {
			res = append(res, ConvertValueOfColumn(basicNow))
		} else { // if exist
			continue
		}
	}
	return res
}

func GetStringColumnData(idx int, dataLi [][]string) (columnStr []string) {
	for _, line := range dataLi {
		columnStr = append(columnStr, line[idx])
	}
	return columnStr
}

func ConvertToFloat64(NumStr string) float64 {
	var res float64
	if NumStr == "0.0000E+000" {
		res = 0
	} else {
		// convert Scientific counting
		res, _ = strconv.ParseFloat(NumStr, 64)
	}

	return res
}

// validateNaN : validate is NaN is existed in data list
// NaN = 0
func validateNaN(nums []string) bool {
	res := true
	for _, v := range nums {
		if v == "NaN" {
			res = false
			break
		}
	}
	return res
}

func ConvertValueOfColumn(basicColumn []string) (columnData []float64) {
	for _, numStr := range basicColumn {
		num := ConvertToFloat64(numStr)
		columnData = append(columnData, num)
	}
	return columnData
}

// ! -------------- concentration mode --------------

func GetConcentrationData(dataLi [][]string) (concentrationData [][]float64) {
	RFLXIdx := 7
	RUNFIdx := 5
	// get str column data
	RFLXStrData := GetStringColumnData(RFLXIdx, dataLi)
	RUNFStrData := GetStringColumnData(RUNFIdx, dataLi)
	// get float column data
	RFLXFloatData := ConvertValueOfColumn(RFLXStrData)
	RUNFFloatData := ConvertValueOfColumn(RUNFStrData)

	res := make([]float64, 0)
	var tmp float64
	for i := range RFLXFloatData {
		if RUNFFloatData[i] != 0 {
			tmp = RFLXFloatData[i] / RUNFFloatData[i] * 1e6
		} else {
			tmp = RFLXFloatData[i] * 1e6
		}
		res = append(res, tmp)
	}

	concentrationData = append(concentrationData, res)
	return concentrationData
}
