package model

type Vegetation struct {
	Vegetation    string  `gorm:"column:vegetation"`
	DensityStem   int64   `gorm:"column:density_stem_s_m2"`
	GrassSpacing  float64 `gorm:"column:grass_spacing_cm"`
	MaximumHeight int64   `gorm:"column:maximum_height_cm"`
	VnModifiedN   float64 `gorm:"column:vn_modified_n_cm_s_1_3"`
	RnaNRange     float64 `gorm:"column:rna_n_range_cm_s_1_3"`
	Vn2BareSoil   float64 `gorm:"column:vn2_bare_soil_n_cm_s_1_3"`
}

func (Vegetation) TableName() string {
	return "vegetation"
}
