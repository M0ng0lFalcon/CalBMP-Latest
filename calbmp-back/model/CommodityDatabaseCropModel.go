package model

import "gorm.io/gorm"

type CommodityDatabaseCrop struct {
	gorm.Model
	SITE_NAME           string  `gorm:"column:site_name"`
	Classes             string  `gorm:"column:classes"`
	Scenarios_residue   string  `gorm:"column:scenarios_residue"`
	Covertype           string  `gorm:"column:covertype"`
	Treatment           string  `gorm:"column:treatment"`
	HydrologicCondition string  `gorm:"column:hydrologiccondition"`
	HSG_A               string  `gorm:"column:hsg_a"`
	HSG_B               string  `gorm:"column:hsg_b"`
	HSG_C               string  `gorm:"column:hsg_c"`
	HSG_D               string  `gorm:"column:hsg_d"`
	Rootdepth_cm        string  `gorm:"column:rootdepth_cm"`
	Canopycover         string  `gorm:"column:canopycover"`
	Height_cm           string  `gorm:"column:height_cm"`
	Holdup_cm           string  `gorm:"column:holdup_cm"`
	Emer_min            string  `gorm:"column:emer_min"`
	Emer_max            string  `gorm:"column:emer_max"`
	Mature_min          string  `gorm:"column:mature_min"`
	Mature_max          string  `gorm:"column:mature_max"`
	Harvest_min         string  `gorm:"column:harvest_min"`
	Harvest_max         string  `gorm:"column:harvest_max"`
	USLECmn             float64 `gorm:"column:uslecmn"`
}

func (CommodityDatabaseCrop) TableName() string {
	return "CommodityDatabaseCrop"
}
