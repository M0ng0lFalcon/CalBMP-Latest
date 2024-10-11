package VFSMService

import (
	"calbmp-back/Params/VFSMParams"
	"calbmp-back/RedisUtil"
	"calbmp-back/util/StringUtil"
	"fmt"
	"strings"
)

func GenerateIsoFun(rec VFSMParams.IsoParams) string {
	res := make([]string, 0)
	template := ""
	tempLine := ""

	soilTexture := rec.SoilTexture

	// VKS  SAV OS OI SM SCHK
	template = "%f %f %f %f %.2f %d"
	VKS := fmt.Sprintf(
		"%.4E",
		StringUtil.Convert2Float(soilTexture.Vks_m_s_10_6),
	)
	SAV := fmt.Sprintf(
		"%.4E",
		StringUtil.Convert2Float(soilTexture.Vks_m_s_10_6),
	)
	OS := soilTexture.Os_m3_m3
	OI := RedisUtil.GetFloatVal(rec.Username + "_W0") // todo: get from przm input file, record 15(W0)
	SM := 0.0
	SCHK := 1
	//tempLine = fmt.Sprintf(template, VKS, SAV, OS, OI, SM, SCHK)
	tempLine = StringUtil.Var2Line(VKS, SAV, OS, OI, SM, SCHK)
	res = append(res, tempLine)

	if rec.WTD != -1 {
		// WTD (WTD=water table depth(m)
		template = "%f"
		WTD := rec.WTD
		tempLine = fmt.Sprintf(template, WTD)
		res = append(res, tempLine)

		// ITHETATYPE  PAR(I)
		template = "%d %f %f %f %f"
		ITHETATYPE := 1
		ParI := []float64{0.15, 13.46, 1.52, 0.348} // TODO: just for test
		tempLine = fmt.Sprintf(template, ITHETATYPE, ParI[0], ParI[1], ParI[2], ParI[3])
		res = append(res, tempLine)
	}

	//log.Println("iso success")
	return strings.Join(res, "\n")
}
