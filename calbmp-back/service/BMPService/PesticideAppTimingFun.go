package BMPService

import (
	"calbmp-back/Params/InputParams"
	"fmt"
)

func PesticideAppTiming(rec InputParams.UserInputStepReceiver) string {
	res := ""
	// request params
	noOfApp := rec.NoOfApp
	cntPesticide := rec.CntPesticide

	temp := "%d, %d, True, 5, 2, 3, 0"
	res = fmt.Sprintf(temp, noOfApp, cntPesticide)
	return res
}
