package ZipMukeyCokeyRepository

import (
	"calbmp-back/Database"
	"calbmp-back/dto/ZipMukeyCokeyDTO"
	"calbmp-back/model"
)

func FindZipCodeByMukeyCokey(mukey, cokey string) (zipCode string) {
	db := Database.GetDB()
	db.Select("\"ZIP_CODE\"").Where("mukey=? AND cokey=?", mukey, cokey).Scan(&zipCode)
	return
}

func GetCompNameMukeyCokeyByZipCode(ZipCode string) []ZipMukeyCokeyDTO.CompnameMukeyCokeyDTO {
	db := Database.GetDB()
	var ZipMukeyCokeyList []model.ZipMukeyCokey
	CompnameMukeyCokey := make([]ZipMukeyCokeyDTO.CompnameMukeyCokeyDTO, 0)

	db.Select("compname", "mukey", "cokey", "muname").Where(
		"zip_code=? AND compkind='Series'",
		ZipCode,
	).Order("muname").Find(&ZipMukeyCokeyList).Scan(&CompnameMukeyCokey)

	return CompnameMukeyCokey
}
