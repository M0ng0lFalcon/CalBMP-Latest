package dto

type EventBasedOutputsStruct struct {
	RunoffDate    []string    `json:"runoff_date"`
	ScenarioValue [][]float64 `json:"scenario_value"`
	Benchmark     string      `json:"benchmark"`
}
