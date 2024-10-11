package ResultParams

type TextResultParams struct {
	PesticideList []string `json:"pesticide_list"`
	CreatedTime   string   `json:"created_time"`
	ScenarioType  string   `json:"scenario_type"`
	FieldSize     float64  `json:"field_size"`
}
