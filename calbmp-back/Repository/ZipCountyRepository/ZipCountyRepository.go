package ZipCountyRepository

import (
	"calbmp-back/Database"
	"calbmp-back/model"
	"fmt"
)

func GetCountyNames() []string {
	db := Database.GetDB()
	var ZipCountyList []model.ZipCounty
	CountyNames := make([]string, 0)

	db.Select("name").Distinct("name").Order("name").Find(&ZipCountyList).Scan(&CountyNames)

	return CountyNames
}

func GetZipCodes(CountyName string) []string {
	db := Database.GetDB()
	ZipCodes := make([]string, 0)

	sqlTemp := "select distinct zip_code" +
		" from zip_county" +
		" where name = '%s' " +
		"  and zip_code in ( " +
		"    select zip_code " +
		"    from zip_mukey_cokey " +
		"    where zip_code = zip_county.zip_code " +
		"      and compname is not null " +
		") order by zip_code;"
	sql := fmt.Sprintf(sqlTemp, CountyName)

	db.Raw(sql).Scan(&ZipCodes)

	return ZipCodes
}
