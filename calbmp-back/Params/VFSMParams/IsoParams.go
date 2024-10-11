package VFSMParams

import "calbmp-back/model"

type IsoParams struct {
	WTD float64 `json:"wtd"`

	// set by controller
	Username    string `json:"username"`
	Mukey       string `json:"mukey"`
	Cokey       string `json:"cokey"`
	JsonPath    string `json:"json_path"`
	SoilTexture model.Soil_texture_final
}
