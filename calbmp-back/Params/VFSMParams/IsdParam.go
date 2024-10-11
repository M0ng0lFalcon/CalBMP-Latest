package VFSMParams

import "calbmp-back/model"

type IsdParams struct {
	// set by controller
	Mukey    string  `json:"mukey"`
	Cokey    string  `json:"cokey"`
	CI       float64 `json:"ci"`
	JsonPath string  `json:"json_path"`
	Chorizon model.ChorizonModel
}
