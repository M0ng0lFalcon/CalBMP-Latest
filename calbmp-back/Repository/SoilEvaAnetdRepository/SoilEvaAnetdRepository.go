package SoilEvaAnetdRepository

import (
	"calbmp-back/Database"
	"calbmp-back/model"
)

func FindSoilEvaAnetdByZipCode(ZipCode string) model.SoilEvaAnetd {
	var SoilEvaAnetd model.SoilEvaAnetd
	db := Database.GetDB()
	db.Where("\"ZIP_CODE\" = ?", ZipCode).Order("\"ANETD_cm\"").First(&SoilEvaAnetd)
	return SoilEvaAnetd
}
