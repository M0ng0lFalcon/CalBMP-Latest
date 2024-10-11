package VFSMParams

import "calbmp-back/model"

type IwqParams struct {
	// set by controller
	CreatedAt     string   `json:"created_at"`
	DGPIN         float64  `json:"dgpin"`
	NDGDAY        int64    `json:"ndgday"`
	JsonPath      string   `json:"json_path"`
	DgT           []string `json:"dg_t"`
	DgTheta       []string `json:"dg_theta"`
	CropPesticide model.CropPesticide
	Chorizon      model.ChorizonModel
}
