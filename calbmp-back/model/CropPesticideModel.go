package model

type CropPesticide struct {
	Site_name                                             string `gorm:"column:site_name"`
	Product_name                                          string `gorm:"column:product_name"`
	Pounds_product_applied                                string `gorm:"column:pounds_product_applied"`
	Pounds_chemical_applied                               string `gorm:"column:pounds_chemical_applied"`
	Amount_treated                                        string `gorm:"column:amount_treated"`
	Aerial_ground_indicator                               string `gorm:"column:aerial_ground_indicator"`
	Chemical_name                                         string `gorm:"column:chemical_name"`
	Usepa_aquatic_life_benchmarks_ug_l                    string `gorm:"column:usepa_aquatic_life_benchmarks_ug_l"`
	Usepa_aquatic_life_benchmarks_ppm                     string `gorm:"column:usepa_aquatic_life_benchmarks_ppm"`
	Dpr_pccode                                            string `gorm:"column:dpr_pccode"`
	Cas_num                                               string `gorm:"column:cas_num"`
	Usetype                                               string `gorm:"column:usetype"`
	Chem_group                                            string `gorm:"column:chem_group"`
	One_dt50_foliar_days_1                                string `gorm:"column:one_dt50_foliar_days_1"`
	Dt50_foliar_days_1                                    string `gorm:"column:dt50_foliar_days_1"`
	Substance_origin                                      string `gorm:"column:substance_origin"`
	Logp_logarithm_of_octanol_water_partition_coefficient string `gorm:"column:logp_logarithm_of_octanol_water_partition_coefficient"`
	Dt50_aerobic_days                                     string `gorm:"column:dt50_aerobic_days"`
	One_dt50_aerobic_days_1                               string `gorm:"column:one_dt50_aerobic_days_1"`
	Henry_num                                             string `gorm:"column:henry_num"`
	Enthalpy_num_j_mol                                    string `gorm:"column:enthalpy_num_j_mol"`
	Constant                                              string `gorm:"column:constant"`
	Enthalpy_num_kcal_mol                                 string `gorm:"column:enthalpy_num_kcal_mol"`
	Freundlich_kf_or_kd                                   string `gorm:"column:freundlich_kf_or_kd"`
	Freundlich_1_n                                        string `gorm:"column:freundlich_1_n"`
	Freundlich_kfoc                                       string `gorm:"column:freundlich_kfoc"`
	Koc                                                   string `gorm:"column:koc"`
	Kd                                                    string `gorm:"column:kd"`
}

func (CropPesticide) TableName() string {
	return "crop_pesticide_benchmark"
}
