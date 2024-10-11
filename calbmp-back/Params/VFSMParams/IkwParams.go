package VFSMParams

import "calbmp-back/model"

type IkwParams struct {
	Vegetation string  `json:"vegetation"`
	FWidth     float64 `json:"f_width"`
	VL         float64 `json:"vl"`

	// set by controller
	CropName string `json:"crop_name"` // get from inp file
	Username string `json:"username"`
	VfsmID   int64  `json:"vfsm_id"`
	JsonPath string `json:"json_path"`
	VegModel model.Vegetation
}
