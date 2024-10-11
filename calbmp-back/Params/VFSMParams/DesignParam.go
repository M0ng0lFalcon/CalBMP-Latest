package VFSMParams

type DesignParam struct {
	// storm
	// by controller
	P             float64
	CN            int
	A             float64
	L             float64
	Y             float64
	SoilType      string
	K             float64
	C             float64
	P2            float64
	SoilOrgMatter string
	// from user
	StormType int `json:"storm_type"`
	D         int `json:"d"`
}
