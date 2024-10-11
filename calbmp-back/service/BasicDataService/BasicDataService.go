package BasicDataService

import (
	"calbmp-back/Repository/ChorizonRepository"
	"calbmp-back/Repository/CommodityDatabaseCropRepository"
	"calbmp-back/Repository/CropPesticideFinalRepository"
	"calbmp-back/Repository/ZipCodeStationRepository"
	"calbmp-back/Repository/ZipCountyRepository"
	"calbmp-back/Repository/ZipMukeyCokeyRepository"
	"calbmp-back/dto/ChorizonDTO"
	"calbmp-back/dto/ZipMukeyCokeyDTO"
	"calbmp-back/model"
)

func GetAllSoilFun() []ChorizonDTO.ChorizonBasicDTO {
	soilList := ChorizonRepository.FindAll()
	return soilList
}

func GetZipCodeBySoil(mukey, cokey string) (zipCode string) {
	zipCode = ZipMukeyCokeyRepository.FindZipCodeByMukeyCokey(mukey, cokey)
	return
}

func GetCountyNames() []string {
	CountyNames := ZipCountyRepository.GetCountyNames()

	return CountyNames
}

func GetZipCodes(CountyName string) []string {
	ZipCodes := ZipCountyRepository.GetZipCodes(CountyName)

	return ZipCodes
}

func GetCompnamesMukeyCokey(ZipCode string) []ZipMukeyCokeyDTO.CompnameMukeyCokeyDTO {
	CompnameMukeyCokeyList := ZipMukeyCokeyRepository.GetCompNameMukeyCokeyByZipCode(ZipCode)

	return CompnameMukeyCokeyList
}

func GetStation(ZipCode string) model.ZipcodeStation {
	ZipCodeStation := ZipCodeStationRepository.GetStationDataByZipCode(ZipCode)

	return ZipCodeStation
}

func GetCropNames() []string {
	CropNames := CommodityDatabaseCropRepository.GetAllCrops()

	return CropNames
}

func GetPesticides() []string {
	Pesticides := CropPesticideFinalRepository.GetAllPesticideName()

	return Pesticides
}
