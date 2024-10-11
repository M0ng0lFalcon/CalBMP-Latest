package model

type WeaData struct {
	Climateid string  `gorm:"column:climateid"`
	ZipCode   string  `gorm:"column:ZIP_CODE"`
	Log       float64 `gorm:"column:log"`
	Lat       float64 `gorm:"column:lat"`
	Year      int     `gorm:"column:Year"`
	Month     int     `gorm:"column:Month"`
	Day       int     `gorm:"column:Day"`

	Precip         string `gorm:"column:pr_cm_d"`
	Evapotranspire string `gorm:"column:Etr_cm_d"`
	Temperature    string `gorm:"column:tm_degrees"`
	Velocity       string `gorm:"column:vs_cm_s"`
	SolarRadiation string `gorm:"column:sr_La_day"`
}
