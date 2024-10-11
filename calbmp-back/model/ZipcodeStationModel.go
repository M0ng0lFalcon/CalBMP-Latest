package model

type ZipcodeStation struct {
	ZipCode   string  `gorm:"column:ZIP_CODE"`
	ClimateId string  `gorm:"column:climateID"`
	Log       float64 `gorm:"column:log_1"`
	Lat       float64 `gorm:"column:lat_1"`
}

func (ZipcodeStation) TableName() string {
	return "zipcode_station"
}
