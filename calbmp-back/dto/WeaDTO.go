package dto

type WeaDTO struct {
	WeaItems []WeaItem
}

type WeaItem struct {
	Data           string `json:"Data"`
	Evapotranspire string `json:"Evapotranspire"`
	SolarRadiation string `json:"SolarRadiation"`
	Temperature    string `json:"Temperature"`
	Velocity       string `json:"Velocity"`
	Prec           string `json:"Prec"`
}
