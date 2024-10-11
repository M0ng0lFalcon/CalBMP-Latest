package ChorizonRepository

import (
	"calbmp-back/Database"
	"calbmp-back/dto/ChorizonDTO"
	"calbmp-back/model"
)

func FindAll() []ChorizonDTO.ChorizonBasicDTO {
	var temp []model.ChorizonModel
	res := make([]ChorizonDTO.ChorizonBasicDTO, 0)
	db := Database.GetDB()
	db.Model(&model.ChorizonModel{}).Select(
		[]string{"id", "mukey", "cokey", "chkey", "muname"},
	).Find(&temp).Scan(&res)
	return res
}

func FindChorizonByMukeyAndCokeyAndHzdept(Mukey string, Cokey string, Hzdept string) model.ChorizonModel {
	var Chorizon model.ChorizonModel
	db := Database.GetDB()
	db.Where("mukey = ? AND cokey = ? AND hzdept = ?", Mukey, Cokey, Hzdept).First(&Chorizon)
	return Chorizon
}

func FindChorizonListByMukeyAndCokey(Mukey string, Cokey string) []model.ChorizonModel {
	var ChorizonList []model.ChorizonModel
	db := Database.GetDB()

	db.Where("mukey = ? AND cokey = ?", Mukey, Cokey).Order("hzdept").Find(&ChorizonList)

	return ChorizonList
}

func FindHSGTypeByCompNameAndMukeyAndCokey(Mukey string, Cokey string) string {
	db := Database.GetDB()
	var Chorizon model.ChorizonModel
	var HSGType string

	db.Select(
		"hydgrp",
	).Where(
		"cokey=? AND mukey=? AND hydgrp IS NOT NULL",
		Cokey,
		Mukey,
	).First(&Chorizon).Scan(&HSGType)

	return HSGType
}

func FindMuname(keyword string) (res []string) {
	db := Database.GetDB()
	whereStr := "compkind='Series' AND slope_per <> '0' AND slopelength_m <> '0' AND muname LIKE ?"
	db.Model(&model.ChorizonModel{}).Debug().Distinct("muname").Where(whereStr, "%"+keyword+"%").Select("muname").Scan(&res)
	return res
}

func FindCompname(muname string) (res []string) {
	db := Database.GetDB()
	db.Model(&model.ChorizonModel{}).Distinct("compname").Where("muname=? AND slope_per <> '0' AND slopelength_m <> '0'", muname).Select("compname").Scan(&res)
	return res
}

func FindByMunameCompname(muname, compname string) (mukey, cokey string) {
	db := Database.GetDB()
	var item model.ChorizonModel
	db.Model(&model.ChorizonModel{}).Where("muname=? AND compname=? AND slope_per <> '0' AND slopelength_m <> '0'", muname, compname).First(&item)
	mukey = item.Mukey
	cokey = item.Cokey
	return
}
