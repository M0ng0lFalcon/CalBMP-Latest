package VFSMService

import (
	"calbmp-back/Params/VFSMParams"
	"calbmp-back/util/StringUtil"
	"strings"
)

func GenerateIgrFun(rec VFSMParams.IgrParams) string {
	// 2.2    0.012	15.0   .04    0
	res := make([]string, 0)

	vegModel := rec.VegModel

	SS := vegModel.GrassSpacing
	VN := vegModel.VnModifiedN
	H := int64(0)

	if rec.MaximumHeight == -1 {
		H = vegModel.MaximumHeight
	} else {
		H = rec.MaximumHeight
	}
	VN2 := vegModel.Vn2BareSoil
	ICO := 0

	tempLine := StringUtil.Var2Line(SS, VN, H, VN2, ICO)
	res = append(res, tempLine)

	//log.Println("igr success")
	return strings.Join(res, "\n")
}
