package model

type Soil_texture_final struct {
	Objectid     string `gorm:"column:objectid"`
	Mukey        string `gorm:"column:mukey"`
	Cokey        string `gorm:"column:cokey"`
	Chkey        string `gorm:"column:chkey"`
	Hzdept       string `gorm:"column:hzdept"`
	Hzthk        string `gorm:"column:hzthk"`
	Sandtotal    string `gorm:"column:sandtotal"`
	Claytotal    string `gorm:"column:claytotal"`
	Om           string `gorm:"column:om"`
	Orgc         string `gorm:"column:orgc"`
	Dbthirdb_g   string `gorm:"column:dbthirdb_g"`
	Wthirdbar    string `gorm:"column:wthirdbar"`
	Wthirdbar1   string `gorm:"column:wthirdbar1"`
	Wfifteen     string `gorm:"column:wfifteen"`
	Wfifteenba   string `gorm:"column:wfifteenba"`
	Kwfact       string `gorm:"column:kwfact"`
	Slope        string `gorm:"column:slope"`
	Slopelengt   string `gorm:"column:slopelengt"`
	Slopelen_1   string `gorm:"column:slopelen_1"`
	Drainagecl   string `gorm:"column:drainagecl"`
	Hydgrp       string `gorm:"column:hydgrp"`
	Rootznemc    string `gorm:"column:rootznemc"`
	Muname       string `gorm:"column:muname"`
	Oid          string `gorm:"column:oid"`
	Objectid_1   string `gorm:"column:objectid_1"`
	Texture      string `gorm:"column:texture"`
	Stratextsf   string `gorm:"column:stratextsf"`
	Rvindicato   string `gorm:"column:rvindicato"`
	Texdesc      string `gorm:"column:texdesc"`
	Chkey_1      string `gorm:"column:chkey_1"`
	Chtgkey      string `gorm:"column:chtgkey"`
	Oid_1        string `gorm:"column:oid_1"`
	Objectid_2   string `gorm:"column:objectid_2"`
	Texcl        string `gorm:"column:texcl"`
	Lieutex      string `gorm:"column:lieutex"`
	Chtgkey_1    string `gorm:"column:chtgkey_1"`
	Chtkey       string `gorm:"column:chtkey"`
	Texture2     string `gorm:"column:texture2"`
	Vks_m_s_10_6 string `gorm:"column:vks_m_s_10-6"`
	Sav_m        string `gorm:"column:sav_m"`
	Os_m3_m3     string `gorm:"column:os_m3_m3"`
}

func (Soil_texture_final) TableName() string {
	return "soil_texture_final"
}
