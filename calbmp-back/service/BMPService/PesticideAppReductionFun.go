package BMPService

import (
	"calbmp-back/Params/BmpParams"
	"calbmp-back/Params/InputParams"
	"calbmp-back/util/StringUtil"
	"fmt"
	"strconv"
	"strings"
)

func PesticideAppReduction(
	rec InputParams.UserInputStepReceiver,
	PAR BmpParams.PesticideAppReductionParams,
) string {
	res := ""

	// get some params
	Date := rec.Date
	NoOfApp := rec.NoOfApp
	cam := rec.ApplicationMethod
	depi := make([]int, len(cam))
	Amount := rec.Amount
	ApplicationEquipment := rec.ApplicationEquipment

	// set depi
	for i, v := range cam {
		if v >= 1 && v <= 3 {
			depi[i] = 0
		} else {
			depi[i] = cam[i]
		}
	}

	// drft
	drft := make([]float64, len(ApplicationEquipment))
	for i, sqjRate := range ApplicationEquipment {
		if sqjRate == 0.95 {
			drft[i] = 0.05
		} else if sqjRate == 1 {
			drft[i] = 0.0
		} else {
			drft[i] = 0.01
		}
	}

	// date cam depi amount sqjRate drft
	temp := "%s,%d,%d,%f,%f,%f"

	for i := 0; i < NoOfApp; i++ {
		if i != 0 {
			res += "\n"
		}
		// format date
		DateLi := strings.Split(Date[i], "-")
		Year := DateLi[0]
		Month := DateLi[1]
		Day := DateLi[2]
		dateVar := fmt.Sprintf("%s,%s,%s",
			StringUtil.DeleteFrontZero(Day),
			StringUtil.DeleteFrontZero(Month),
			Year)

		amountFloat, _ := strconv.ParseFloat(Amount[i], 10)
		amountFloat *= 1.12085116

		res += fmt.Sprintf(temp,
			dateVar,
			cam[i],
			depi[i],
			amountFloat,
			ApplicationEquipment[i]*(1.0-(float64(PAR.Rate)/100)),
			drft[i],
		)
	}

	return res
}
