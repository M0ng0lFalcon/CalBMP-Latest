package HistoryParams

type HistoryRec struct {
	ProjectName string `json:"project_name"`
	Muname      string `json:"muname"`
	ZipCode     string `json:"zip_code"`
	County      string `json:"county"`
	CompName    string `json:"comp_name"`
	Username    string `json:"username"`
	Step1       string `json:"step1"`
	Step2       string `json:"step2"`
	CreatedTime string `json:"created_time"`
	EchartList  int    `json:"echart_list"`
}
