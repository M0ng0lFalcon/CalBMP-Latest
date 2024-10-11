package CommodityDatabaseUSLECRepository

import (
	"calbmp-back/Database"
	"calbmp-back/model"
	"strconv"
)

func FindUSLECValuesByCropClass(CropClass string) []float64 {
	db := Database.GetDB()

	// return value
	USLECList := make([]float64, 24)

	// temp value
	USLECStrList := make([]string, 24)

	var CDUSLECList []model.CommodityDatabaseUSLEC
	db.Where("\"Classes\" = ?", CropClass).Find(&CDUSLECList)

	// 1. get uslec value by string
	// item struct: CN, N, USLEC
	for _, v := range CDUSLECList {
		if v.Parameters == "USLEC" {
			USLECStrList[0] = v.C1
			USLECStrList[1] = v.C2
			USLECStrList[2] = v.C3
			USLECStrList[3] = v.C4
			USLECStrList[4] = v.C5
			USLECStrList[5] = v.C6
			USLECStrList[6] = v.C7
			USLECStrList[7] = v.C8
			USLECStrList[8] = v.C9
			USLECStrList[9] = v.C10
			USLECStrList[10] = v.C11
			USLECStrList[11] = v.C12
			USLECStrList[12] = v.C13
			USLECStrList[13] = v.C14
			USLECStrList[14] = v.C15
			USLECStrList[15] = v.C16
			USLECStrList[16] = v.C17
			USLECStrList[17] = v.C18
			USLECStrList[18] = v.C19
			USLECStrList[19] = v.C20
			USLECStrList[20] = v.C21
			USLECStrList[21] = v.C22
			USLECStrList[22] = v.C23
			USLECStrList[23] = v.C24

			break
		}
	}

	// 2. convert string to float
	for i, v := range USLECStrList {
		// uslec = c
		C, _ := strconv.ParseFloat("0"+v, 64)
		USLECList[i] = C
	}

	return USLECList
}

func FindNValueByClass(CropClass string) []float64 {
	db := Database.GetDB()

	// return value
	NList := make([]float64, 24)

	// temp value
	NStrList := make([]string, 24)

	// get data orm
	var CDUSLECList []model.CommodityDatabaseUSLEC
	db.Where("\"Classes\" = ?", CropClass).Find(&CDUSLECList)

	// 1. get N value by string
	// item struct: CN, N, USLEC
	for _, v := range CDUSLECList {
		if v.Parameters == "n" {
			NStrList[0] = v.C1
			NStrList[1] = v.C2
			NStrList[2] = v.C3
			NStrList[3] = v.C4
			NStrList[4] = v.C5
			NStrList[5] = v.C6
			NStrList[6] = v.C7
			NStrList[7] = v.C8
			NStrList[8] = v.C9
			NStrList[9] = v.C10
			NStrList[10] = v.C11
			NStrList[11] = v.C12
			NStrList[12] = v.C13
			NStrList[13] = v.C14
			NStrList[14] = v.C15
			NStrList[15] = v.C16
			NStrList[16] = v.C17
			NStrList[17] = v.C18
			NStrList[18] = v.C19
			NStrList[19] = v.C20
			NStrList[20] = v.C21
			NStrList[21] = v.C22
			NStrList[22] = v.C23
			NStrList[23] = v.C24

			break
		}
	}

	// 2. convert string to float
	for i, v := range NStrList {
		N, _ := strconv.ParseFloat("0"+v, 64)
		NList[i] = N
	}

	return NList
}
