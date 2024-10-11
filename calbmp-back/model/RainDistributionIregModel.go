package model

import "gorm.io/gorm"

type RainDistributionIreg struct {
	gorm.Model
	ZIP_CODE string `gorm:"column:ZIP_CODE"`
	PO_NAME  string `gorm:"column:PO_NAME"`
	NAME     string `gorm:"column:NAME"`
	NAMELSAD string `gorm:"column:NAMELSAD"`
	IREG     string `gorm:"column:IREG"`
}

func (RainDistributionIreg) TableName() string {
	return "rain_distribution_ireg_summary"
}
