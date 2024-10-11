package VFSMService

import (
	"calbmp-back/Params/VFSMParams"
	"calbmp-back/model"
	"calbmp-back/util/JsonUtil"
	"calbmp-back/util/StringUtil"
	"fmt"
	"strings"
)

type IrnStruct struct {
	// variable
	// line 1
	NRAIN int64
	RPEAK float64
	// line 2
	RAIN []string

	// parameter
	Prec      float64
	Pfstation model.Pfstation
	param     VFSMParams.IrnParams

	// result
	resLi []string
	res   string
}

func (_this *IrnStruct) init(irnParam VFSMParams.IrnParams) {
	// default value
	_this.NRAIN = 4

	// parameters
	_this.param = irnParam
	_this.Prec = irnParam.Prec
	_this.Pfstation = irnParam.Pfstation
}

func (_this *IrnStruct) line1() (line string) {
	// NRAIN RPEAK
	_this.RPEAK = _this.Pfstation.Rainfall_intensity_mm_s * 2 / 1000
	rpeak := StringUtil.Float2ScientificNotation(_this.RPEAK)

	line = StringUtil.Var2Line(_this.NRAIN, rpeak)
	_this.resLi = append(_this.resLi, line)

	return
}

func (_this *IrnStruct) line2() (line string) {
	var tmp1 string
	var tmp2, tmp3 float64
	// line 1
	tmp1 = ".0000E+00"
	tmp2 = _this.Pfstation.Rainfall_intensity_mm_s / 1000
	_this.RAIN = append(_this.RAIN, fmt.Sprintf("%s %.4E", tmp1, tmp2))

	// line2
	tmp2 = _this.Prec / _this.Pfstation.Rainfall_intensity_mm_s / 2.67
	tmp3 = _this.Pfstation.Rainfall_intensity_mm_s * 2 / 1000
	_this.RAIN = append(_this.RAIN, fmt.Sprintf("%.4E %.4E", tmp2, tmp3))

	// line 3
	tmp2 = _this.Prec * 10 / _this.Pfstation.Rainfall_intensity_mm_s
	tmp3 = _this.Pfstation.Rainfall_intensity_mm_s / 1000
	_this.RAIN = append(_this.RAIN, fmt.Sprintf("%.4E %.4E", tmp2, tmp3))

	// line 4
	tmp1 = ".8640E+05"
	tmp2 = 0
	_this.RAIN = append(_this.RAIN, fmt.Sprintf("%s %.4E", tmp1, tmp2))

	line = strings.Join(_this.RAIN, "\n")
	_this.resLi = append(_this.resLi, line)
	return
}

func (_this *IrnStruct) joinLine() (res string) {
	_this.line1()
	_this.line2()

	_this.res = strings.Join(_this.resLi, "\n")
	return _this.res
}

func (_this *IrnStruct) toJson() {
	JsonUtil.WriteJson(_this.param.JsonPath, _this)
}

func GenerateIrnFun(irnParam VFSMParams.IrnParams) string {
	var irn IrnStruct
	irn.init(irnParam)
	res := irn.joinLine()
	irn.toJson()

	//log.Println("irn success")
	return res
}
