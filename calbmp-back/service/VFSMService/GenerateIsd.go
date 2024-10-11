package VFSMService

import (
	"calbmp-back/Params/VFSMParams"
	"calbmp-back/model"
	"calbmp-back/util/JsonUtil"
	"calbmp-back/util/StringUtil"
	"strings"
)

type IsdStruct struct {
	// line 1
	NPART  int64
	COARSE float64
	CI     float64
	POR    float64

	// line 2
	// if npart=7
	DP float64
	SG float64
	// if npart = 8
	Silt_frac float64
	Itillage  int64

	// variable
	param    VFSMParams.IsdParams
	Chorizon model.ChorizonModel
	Mukey    string
	Cokey    string

	// result
	resLi []string
	res   string
}

func (_this *IsdStruct) init(param VFSMParams.IsdParams) {
	_this.Mukey = param.Mukey
	_this.Cokey = param.Cokey
	_this.CI = param.CI
	_this.param = param

	_this.Chorizon = param.Chorizon

	_this.NPART = 8
	_this.POR = 0.434
	_this.Itillage = 0
}

func (_this *IsdStruct) line1() {
	_this.COARSE = StringUtil.Convert2Float(_this.Chorizon.Sandtotal_per) / 100.0

	line := StringUtil.Var2Line(_this.NPART, _this.COARSE, _this.CI, _this.POR)
	_this.resLi = append(_this.resLi, line)
}

func (_this *IsdStruct) line2() {
	_this.Silt_frac = StringUtil.Convert2Float(_this.Chorizon.Totalsilt_per) / 100.0
	_this.Itillage = 0
	line := StringUtil.Var2Line(_this.Silt_frac, _this.Itillage)
	_this.resLi = append(_this.resLi, line)
}

func (_this *IsdStruct) joinLine() (res string) {
	_this.line1()
	_this.line2()

	_this.res = strings.Join(_this.resLi, "\n")
	return _this.res
}

func (_this *IsdStruct) toJson() {
	JsonUtil.WriteJson(_this.param.JsonPath, _this)
}

func GenerateIsdFun(param VFSMParams.IsdParams) string {
	var isd IsdStruct
	isd.init(param)
	res := isd.joinLine()
	isd.toJson()

	//log.Println("isd success")
	return res
}
