package Record5DTO

import "time"

type Record5DTO struct {
	UnitId        int
	EmergenceDate time.Time
	HarvestDate   time.Time
	USLECList     []float64
	NList         []float64
	CNValue       int
	BareValue     int
	ResidueValue  float64
	CropName      string
}
