package CRRMUtil

import (
	"calbmp-back/Params/BmpParams"
	"calbmp-back/util/TimeUtil"
	"sort"
)

func SortDateByEmergence(CropInfos []BmpParams.CropInfoParams) {
	sort.SliceStable(CropInfos, func(i, j int) bool {
		EmergenceDateString1 := CropInfos[i].EmergenceDate
		EmergenceDateString2 := CropInfos[j].EmergenceDate

		EmergenceDate1 := TimeUtil.ParseTimeString(EmergenceDateString1)
		EmergenceDate2 := TimeUtil.ParseTimeString(EmergenceDateString2)

		return EmergenceDate1.Before(EmergenceDate2)
	})
}
