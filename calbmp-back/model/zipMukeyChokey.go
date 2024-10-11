package model

import "gorm.io/gorm"

type ZipMukeyCokey struct {
	gorm.Model
	Fid_ca_zip string
	Objectid_2 int
	Zip_code   string
	Po_name    string
	State      string
	Compname   string
	Compkind   string
	Majcompfla string
	Otherph    string
	Geomdesc   string
	Mukey      string
	Cokey      string
	Objectid_1 int
	Musym      string
	Muname     string
}

func (ZipMukeyCokey) TableName() string {
	return "zip_mukey_cokey"
}
