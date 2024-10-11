package VegetationRepository

import (
	"calbmp-back/Database"
	"calbmp-back/model"
)

func FindByVegetation(vegetation string) (item model.Vegetation) {
	db := Database.GetDB()
	db.Where("vegetation=?", vegetation).First(&item)
	return item
}

func FindAllList() (vegetation []string) {
	var itemLi []model.Vegetation
	db := Database.GetDB()
	db.Select("vegetation").Find(&itemLi).Scan(&vegetation)
	return
}
