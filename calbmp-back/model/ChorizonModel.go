package model

type ChorizonModel struct {
	Mukey               string `gorm:"column:mukey"`
	Cokey               string `gorm:"column:cokey"`
	Chkey               string `gorm:"column:chkey"`
	Hzdept              string `gorm:"column:hzdept"`
	Hzthk               string `gorm:"column:hzthk"`
	Sandtotal_per       string `gorm:"column:sandtotal_per"`
	Claytotal_per       string `gorm:"column:claytotal_per"`
	Om                  string `gorm:"column:om"`
	Orgc_per            string `gorm:"column:orgc_per"`
	Dbthirdb_g_cm3      string `gorm:"column:dbthirdb_g_cm3"`
	Wthirdbar_cm3_cm3   string `gorm:"column:wthirdbar_cm3_cm3"`
	Wfifteenbar_cm3_cm3 string `gorm:"column:wfifteenbar_cm3_cm3"`
	Kwfact              string `gorm:"column:kwfact"`
	Slope_per           string `gorm:"column:slope_per"`
	Slopelength_m       string `gorm:"column:slopelength_m"`
	Hydgrp              string `gorm:"column:hydgrp"`
	Rootznemc           string `gorm:"column:rootznemc"`
	Muname              string `gorm:"column:muname"`
	Compname            string `gorm:"column:compname"`
	Compkind            string `gorm:"column:compkind"`
	Totalsilt_per       string `gorm:"column:totalsilt_per"`
}

func (ChorizonModel) TableName() string {
	return "chorizon_final3"
}
