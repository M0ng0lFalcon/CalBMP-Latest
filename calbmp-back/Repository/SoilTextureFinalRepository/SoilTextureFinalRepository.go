package SoilTextureFinalRepository

import (
	"calbmp-back/Database"
	"calbmp-back/model"
)

func FindByMukeyAndCokey(Mukey string, Cokey string) model.Soil_texture_final {
	db := Database.GetDB()
	var soilTexture model.Soil_texture_final

	db.Where(
		"cokey=? AND mukey=? AND hydgrp IS NOT NULL",
		Cokey,
		Mukey,
	).Order("hzdept").First(&soilTexture)

	return soilTexture
}
