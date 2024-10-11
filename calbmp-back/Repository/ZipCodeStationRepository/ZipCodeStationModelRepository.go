package ZipCodeStationRepository

import (
	"calbmp-back/Database"
	"calbmp-back/model"
)

func GetStationDataByZipCode(ZipCode string) model.ZipcodeStation {
	db := Database.GetDB()
	var ZipCodeStation model.ZipcodeStation
	db.Where("\"ZIP_CODE\"=?", ZipCode).First(&ZipCodeStation)

	return ZipCodeStation
}
