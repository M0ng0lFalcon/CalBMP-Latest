package VFSMParams

import "calbmp-back/model"

type IgrParams struct {
	Vegetation    string `json:"vegetation"`
	MaximumHeight int64  `json:"maximum_height"`

	// set by controller
	JsonPath string `json:"json_path"`
	VegModel model.Vegetation
}
