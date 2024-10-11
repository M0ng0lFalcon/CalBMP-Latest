package BMPService

import (
	"calbmp-back/Params/InputParams"
	"calbmp-back/util/GlobalValueUtil"
	"calbmp-back/util/InputFileGenerateUtil"
	"fmt"
)

func ChangeRecordCNValue(
	MinusNum int,
	gv *GlobalValueUtil.GlobalVar,
	rec InputParams.UserInputStepReceiver,
) string {
	res := ""

	// get params
	USLECList := gv.USLEC
	NList := gv.N
	CNValueStr := gv.CNValue
	BareValueStr := gv.BareSoil

	// day, month, year, c = uslec, n, cn
	temp := "%2d, %2d, %4d,   %.5f, %.5f, %d"

	// day, month, year : int
	DateLiWithCN := InputFileGenerateUtil.GenerateDateListWithCN(rec, CNValueStr, BareValueStr)
	for i := 0; i < 24; i++ {
		if i != 0 {
			res += "\n"
		}
		CNValue := DateLiWithCN[i][3]
		CNFlag := DateLiWithCN[i][4]
		if CNFlag == 1 {
			CNValue -= MinusNum
		}
		res += fmt.Sprintf(temp,
			DateLiWithCN[i][0],
			DateLiWithCN[i][1],
			DateLiWithCN[i][2],
			USLECList[i],
			NList[i],
			CNValue,
		)
	}

	return res
}
