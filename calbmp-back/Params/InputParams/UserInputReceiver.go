package InputParams

type UserInputStepReceiver struct {
	// public params
	ZipCode   string  `json:"ZipCode"`
	ClimateId string  `json:"climateId"`
	FieldSize string  `json:"FieldSize"`
	MuKey     string  `json:"mukey"`
	CoKey     string  `json:"cokey"`
	Log       float64 `json:"log"`
	Lat       float64 `json:"lat"`

	// user input step 1
	WeaFilePath string  `json:"wea_file_path"`
	ZtsFilePath string  `json:"zts_file_path"`
	County      string  `json:"county"`
	CompName    string  `json:"comp_name"`
	Muname      string  `json:"muname"`
	KnowSlope   bool    `json:"know_slope"`
	Slope       float64 `json:"slope"`

	// user input step 2
	Crop                 string    `json:"Crop"`
	Emergence            string    `json:"EmergenceDate"`
	Maturity             string    `json:"MaturityDate"`
	Harvest              string    `json:"HarvestDate"`
	IrrigationType       int       `json:"IrrigationType"`
	IrrigationDate       []string  `json:"irrigation_date"`
	IrrigationAmount     []float64 `json:"irrigation_amount"`
	Pesticide            []string  `json:"Pesticide"`
	CntPesticide         int       `json:"cntPesticide"`
	PesticideSet         []string  `json:"pesticideSet"`
	ApplicationEquipment []float64 `json:"ApplicationEquipment"`
	ApplicationMethod    []int     `json:"ApplicationMethod"`
	Depth                []float64 `json:"depth"`
	NoOfApp              int       `json:"NOofApp"`
	Date                 []string  `json:"date"`
	Amount               []string  `json:"amount"`
	IfBMP                bool      `json:"ifBmp"`
	IfVfsm               bool      `json:"if_vfsm"`
	UserDefinedIrrgRate  float64   `json:"user_defined_irrg_rate"`

	// set by controller
	Username string
	CurTime  string
}
