package CropPesticideFinalRepository

import (
	"calbmp-back/Database"
	"calbmp-back/model"
)

func FindByCropName(cropName string) model.CropPesticide {
	var item model.CropPesticide
	db := Database.GetDB()
	db.Where("")
	return item
}

func GetAllPesticideName() []string {
	db := Database.GetDB()
	var CropPesticide model.CropPesticide

	Pesticides := make([]string, 0)

	param := "chemical_name"

	db.Distinct(param).Select(param).Where(param + " IS NOT NULL").Order(param).Find(&CropPesticide).Scan(&Pesticides)

	return Pesticides
}

func FindCropPesticideByChemicalName(ChemicalName string) model.CropPesticide {
	db := Database.GetDB()
	var CropPesticide model.CropPesticide

	db.Where("chemical_name=?", ChemicalName).First(&CropPesticide)

	return CropPesticide
}

func FindBenchMarkValueByChemicalName(ChemicalName string) (res float64) {
	db := Database.GetDB()
	//var CropPesticide model.CropPesticide

	sql := "SELECT usepa_aquatic_life_benchmarks_ppm\nFROM \"crop_pesticide_benchmark\"\nWHERE chemical_name = ?\nLIMIT 1"

	db.Raw(sql, ChemicalName).Scan(&res)

	return res
}
