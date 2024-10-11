package VFSMService

import (
	"calbmp-back/Params/VFSMParams"
	"calbmp-back/util/StringUtil"
	"strings"
)

type InpStruct struct {
	// line 1
	P         float64
	CN        int
	A         float64
	StormType int
	D         int
	L         float64
	Y         float64
	// line 2
	SoilType string
	// line 3
	K  float64
	C  float64
	P2 float64
	Dp float64
	// line 4
	IREOTY int
	// line 5
	SoilOrgMatter string

	// result
	resLi []string
	res   string
}

func (_this *InpStruct) init(param VFSMParams.DesignParam) {
	_this.IREOTY = 1
	_this.Dp = -1
	_this.P = param.P * 10
	_this.CN = param.CN
	_this.A = param.A
	_this.L = param.L
	_this.Y = param.Y
	_this.SoilType = param.SoilType
	_this.SoilOrgMatter = param.SoilOrgMatter
	_this.StormType = param.StormType // from user
	_this.D = param.D                 // from user
	_this.K = param.K

	_this.C = param.C
	_this.P2 = param.P2
}

func (_this *InpStruct) line1() {
	// P, CN, A, storm type, D, L, Y
	line := StringUtil.Var2Line(_this.P, _this.CN, _this.A, _this.StormType, _this.D, _this.L, _this.Y) + "\n"
	_this.resLi = append(_this.resLi, line)
}

func (_this *InpStruct) line2() {
	// Soil type
	_this.resLi = append(_this.resLi, _this.SoilType)
}

func (_this *InpStruct) line3() {
	// K, C, P, Dp
	line := StringUtil.Var2Line(_this.K, _this.C, _this.P, _this.Dp)
	_this.resLi = append(_this.resLi, line)
}

func (_this *InpStruct) line4() {
	// IREOTY
	line := StringUtil.Var2Line(_this.IREOTY)
	_this.resLi = append(_this.resLi, line)
}

func (_this *InpStruct) line5() {
	// soil org matter
	line := StringUtil.Var2Line(_this.SoilOrgMatter)
	_this.resLi = append(_this.resLi, line)
}

func (_this *InpStruct) joinLine() (res string) {
	_this.line1()
	_this.line2()
	_this.line3()
	_this.line4()
	_this.line5()

	_this.res = strings.Join(_this.resLi, "\n")
	return _this.res
}

func GenerateInp(param VFSMParams.DesignParam) string {
	var inp InpStruct
	inp.init(param)
	res := inp.joinLine()
	return res
}
