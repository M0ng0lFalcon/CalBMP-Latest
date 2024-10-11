package model

import "gorm.io/gorm"

type SoilEvaAnetd struct {
	gorm.Model
	ZIP_CODE string  `gorm:"column:ZIP_CODE"`
	ANETD_cm float32 `gorm:"column:ANETD_cm"`
}

func (SoilEvaAnetd) TableName() string {
	return "soil_eva_anetd_final"
}
