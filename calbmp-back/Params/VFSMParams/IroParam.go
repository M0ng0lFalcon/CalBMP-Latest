package VFSMParams

type IroParams struct {
	SWIDTH  float64 `json:"swidth"`
	SLENGTH float64 `json:"slength"`

	// set by controller
	IsRain   bool    `json:"is_rain"`
	Prec     float64 `json:"prec"` // get fom zts file
	JsonPath string  `json:"json_path"`
	IrnJson  string
}
