package ResultDataDTO

type ResData struct {
	Date          []string             `json:"date"`
	Water         map[string][]float64 `json:"water"`
	Pesticide     map[string][]float64 `json:"pesticide"`
	Sediment      map[string][]float64 `json:"sediment"`
	Concentration map[string][]float64 `json:"concentration"`
	Benchmark     map[string]float64   `json:"benchmark"`
}
