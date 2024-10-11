package PfstationRepository

import (
	"calbmp-back/Database"
	"calbmp-back/model"
)

func FindByZipCode(zipCode string) (item model.Pfstation) {
	Database.DB.Where("zip_code=?", zipCode).First(&item)
	return
}
