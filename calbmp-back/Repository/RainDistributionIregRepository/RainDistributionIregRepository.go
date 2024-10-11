package RainDistributionIregRepository

import (
	"calbmp-back/Database"
	"calbmp-back/model"
)

func FindRainDistributionIregByZipCode(ZipCode string) model.RainDistributionIreg {
	var RainDistributionIreg model.RainDistributionIreg
	db := Database.GetDB()

	db.Where("\"ZIP_CODE\" = ?", ZipCode).First(&RainDistributionIreg)

	return RainDistributionIreg
}
