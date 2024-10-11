package BmpParams

import "calbmp-back/Params/InputParams"

type BmpBasicParams struct {
	BmpId       string                            `json:"bmp_id"`
	BmpOpts     []string                          `json:"bmp_opts"`
	Step1Params InputParams.UserInputStepReceiver `json:"step_1_params"`
	Step2Params InputParams.UserInputStepReceiver `json:"step_2_params"`

	// bmp optional params
	PesticideAppReduction            PesticideAppReductionParams            `json:"pesticide_app_reduction"`
	StripCropping                    StripCroppingParams                    `json:"strip_cropping"`
	ParallelTerracing                ParallelTerracingParams                `json:"parallel_terracing"`
	CoverCrops                       CoverCropsParams                       `json:"cover_crops"`
	CropRotationAndResidueManagement CropRotationAndResidueManagementParams `json:"crop_rotation_and_residue_management"`
}

// pesticide application reduction

type PesticideAppReductionParams struct {
	Rate int `json:"rate"`
}

// strip cropping

type StripCroppingParams struct {
	Crop    string  `json:"crop"`
	Rate    float64 `json:"rate"`
	PreCrop string  `json:"pre_crop"`
}

// parallel terracing

type ParallelTerracingParams struct {
	Type string `json:"type"`
}

// cover crop

type CoverCropsParams struct {
	CoverCrop          string `json:"cover_crop"`
	CoverCropEmergence string `json:"cover_crop_emergence"`
	CoverCropHarvest   string `json:"cover_crop_harvest"`
}

// crop rotation and residue management

type CropInfoParams struct {
	CropName      string  `json:"crop_name"`
	EmergenceDate string  `json:"emergence_date"`
	HarvestDate   string  `json:"harvest_date"`
	ResidueValue  float64 `json:"residue_value"`
}

type CropRotationAndResidueManagementParams struct {
	CropInfos []CropInfoParams `json:"crop_infos"`
}
