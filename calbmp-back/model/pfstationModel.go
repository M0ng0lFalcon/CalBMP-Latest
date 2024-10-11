package model

type Pfstation struct {
	Fid                        int64   `gorm:"column:fid"`
	Shape                      string  `gorm:"column:shape"`
	Zip_code                   string  `gorm:"column:zip_code"`
	Climate_id                 float64 `gorm:"column:climate_id"`
	Log_1                      float64 `gorm:"column:log_1"`
	Lat_1                      float64 `gorm:"column:lat_1"`
	Rastervalu                 float64 `gorm:"column:rastervalu"`
	Rainfall_intensity_mm_hour float64 `gorm:"column:rainfall_intensity_mm_hour"`
	Rainfall_intensity_mm_s    float64 `gorm:"column:rainfall_intensity_mm_s"`
}

func (Pfstation) TableName() string {
	return "pfstation"
}
