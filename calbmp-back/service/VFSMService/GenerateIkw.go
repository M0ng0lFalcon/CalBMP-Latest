package VFSMService

import (
	"calbmp-back/Params/VFSMParams"
	"calbmp-back/RedisUtil"
	"calbmp-back/model"
	"calbmp-back/util/JsonUtil"
	"calbmp-back/util/StringUtil"
	"fmt"
	"strings"
)

type IkwStruct struct {
	// variables
	// line 1
	Label string
	// line 2
	FWIDTH float64
	// line 3
	VL      float64
	N       int64
	THETAW  float64
	CR      float64
	MAXITER int64
	NPOL    int64
	IELOUT  int64
	KPG     int64
	// line 4
	NPROP int64
	// line 5
	SX  float64
	RNA float64
	SOA float64
	// line 6
	IWQ int64

	// parameters
	param    VFSMParams.IkwParams
	vegModel model.Vegetation
	slope    float64
	// result
	resLi []string
	res   string
}

func (_this *IkwStruct) init(rec VFSMParams.IkwParams) {
	_this.param = rec
	_this.vegModel = rec.VegModel
	_this.slope = RedisUtil.GetFloatVal(rec.Username + "_slope")

	_this.FWIDTH = rec.FWidth
	_this.VL = rec.VL

	_this.N = 57
	_this.THETAW = 0.5
	_this.CR = 0.8
	_this.MAXITER = 350
	_this.NPOL = 3
	_this.IELOUT = 1
	_this.KPG = 1

	_this.NPROP = 1

	_this.IWQ = 1
}

func (_this *IkwStruct) line1() {
	// LABEL: A label (max, 50 characters) to identify the program run
	header := "%s, vfs%d"
	_this.Label = fmt.Sprintf(header, _this.param.CropName, _this.param.VfsmID)

	_this.resLi = append(_this.resLi, _this.Label)
}

func (_this *IkwStruct) line2() {
	// FWIDTH: width of the strip (m)
	line := StringUtil.Var2Line(_this.FWIDTH)
	_this.resLi = append(_this.resLi, line)
}

func (_this *IkwStruct) line3() {
	// VL  N  THETAW  CR  MAXTIER  NPOL  IELOUT KPG  (VL=Length of the VFS (m)
	template := "%f %d %f %f %d %d %d %d"
	line := fmt.Sprintf(template,
		_this.param.VL,
		_this.N,
		_this.THETAW,
		_this.CR,
		_this.MAXITER,
		_this.NPOL,
		_this.IELOUT,
		_this.KPG,
	)
	_this.resLi = append(_this.resLi, line)
}

func (_this *IkwStruct) line4() {
	// NPROP=number of segments with different surface properties (slope or roughness)
	template := "%d"
	line := fmt.Sprintf(template, _this.NPROP)
	_this.resLi = append(_this.resLi, line)
}

func (_this *IkwStruct) line5() {
	// SX  RNA  SOA
	template := "%f %f %f"
	SX := _this.param.VL / float64(_this.NPROP)
	RNA := _this.vegModel.RnaNRange
	SOA := _this.slope * 0.01

	line := fmt.Sprintf(template, SX, RNA, SOA)
	_this.resLi = append(_this.resLi, line)
}

func (_this *IkwStruct) line6() {
	// IWQ: water quality/transport problem selection flag (0 is do not run problem, 1 run problem which means .iwq file required
	line := StringUtil.Var2Line(_this.IWQ)
	_this.resLi = append(_this.resLi, line)
}

func (_this *IkwStruct) JoinLine() (res string) {
	_this.line1()
	_this.line2()
	_this.line3()
	_this.line4()
	_this.line5()
	_this.line6()

	_this.res = strings.Join(_this.resLi, "\n")
	return _this.res
}

func (_this *IkwStruct) toJson() {
	JsonUtil.WriteJson(_this.param.JsonPath, _this)
}

func GenerateIkwFun(rec VFSMParams.IkwParams) string {
	var ikw IkwStruct
	ikw.init(rec)
	res := ikw.JoinLine()
	ikw.toJson()

	//log.Println("ikw success")
	return res
}
