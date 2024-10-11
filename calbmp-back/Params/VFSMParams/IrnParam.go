package VFSMParams

import "calbmp-back/model"

type IrnParams struct {
	// set by controller
	ZipCode   string  `json:"zip_code"` // get from inp
	Prec      float64 `json:"prec"`     // get from zts
	JsonPath  string  `json:"json_path"`
	Pfstation model.Pfstation
}
