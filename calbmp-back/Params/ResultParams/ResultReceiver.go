package ResultParams

type ResultReceiver struct {
	// ScenarioType : baseline or bmp
	ScenarioType  string   `json:"scenario_type"`
	BmpId         string   `json:"bmp_id"`
	Water         []string `json:"water"`
	Pesticide     []string `json:"pesticide"`
	Sediment      []string `json:"sediment"`
	Concentration []string `json:"concentration"`
	PesticideList []string `json:"pesticide_list"`
	CreatedTime   string   `json:"created_time"`
	FieldSize     float64  `json:"field_size"`
}
