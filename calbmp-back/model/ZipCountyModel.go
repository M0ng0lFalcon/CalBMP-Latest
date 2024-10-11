package model

import "gorm.io/gorm"

type ZipCounty struct {
	gorm.Model
	FID_CA_zip int     `gorm:"column:fid_ca_zip"`
	OBJECTID   int     `gorm:"column:objectid"`
	ZIP_CODE   string  `gorm:"column:zip_code"`
	PO_NAME    string  `gorm:"column:po_name"`
	STATE      string  `gorm:"column:state"`
	POPULATION int     `gorm:"column:population"`
	POP_SQMI   float64 `gorm:"column:pop_sqmi"`
	SQMI       float64 `gorm:"column:sqmi"`
	Shape_Leng float64 `gorm:"column:shape_leng"`
	Shape_Area float64 `gorm:"column:shape_area"`
	FID_CA_Cou int     `gorm:"column:fid_ca_cou"`
	STATEFP    string  `gorm:"column:statefp"`
	COUNTYFP   string  `gorm:"column:countyfp"`
	COUNTYNS   string  `gorm:"column:countyns"`
	GEOID      string  `gorm:"column:geoid"`
	NAME       string  `gorm:"column:name"`
	NAMELSAD   string  `gorm:"column:namelsad"`
	LSAD       string  `gorm:"column:lsad"`
	CLASSFP    string  `gorm:"column:classfp"`
	MTFCC      string  `gorm:"column:mtfcc"`
	CSAFP      string  `gorm:"column:csafp"`
	CBSAFP     string  `gorm:"column:cbsafp"`
	METDIVFP   string  `gorm:"column:metdivfp"`
	FUNCSTAT   string  `gorm:"column:funcstat"`
	ALAND      string  `gorm:"column:aland"`
	AWATER     string  `gorm:"column:awater"`
	INTPTLAT   string  `gorm:"column:intptlat"`
	INTPTLON   string  `gorm:"column:intptlon"`
}

func (ZipCounty) TableName() string {
	return "zip_county"
}
