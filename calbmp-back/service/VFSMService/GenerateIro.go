package VFSMService

import (
	"calbmp-back/Params/VFSMParams"
	"calbmp-back/util/JsonUtil"
	"fmt"
	"strings"
)

type IroStruct struct {
	// variables
	// line 1
	SWIDTH  float64 // 田块宽度（m）
	SLENGTH float64 // 田块长度（m）
	// line 2
	NBCROFF  int64
	BCROPEAK float64
	// line 3
	RAIN []string

	// parameters
	param  VFSMParams.IroParams
	IsRain bool

	// result
	resLi []string
	res   string
}

func (_this *IroStruct) init(param VFSMParams.IroParams) {
	_this.IsRain = param.IsRain
	_this.param = param
	_this.NBCROFF = 4
	_this.SWIDTH = _this.param.SWIDTH
	_this.SLENGTH = _this.param.SLENGTH
}

func (_this *IroStruct) line1() (line string) {
	template := "%f %f"

	line = fmt.Sprintf(template, _this.SWIDTH, _this.SLENGTH)
	_this.resLi = append(_this.resLi, line)
	return line
}

func (_this *IroStruct) line2() (line string) {
	var info IrnStruct
	JsonUtil.ReadJson(_this.param.IrnJson, &info) // todo: 取最新json

	if _this.IsRain {
		_this.BCROPEAK = info.RPEAK
	} else {
		// todo: get from przm5 result , record 9
		_this.BCROPEAK = info.RPEAK
	}

	template := "%d %.4E"
	line = fmt.Sprintf(template, _this.NBCROFF, _this.BCROPEAK)
	_this.resLi = append(_this.resLi, line)
	return line
}

func (_this *IroStruct) line3() (line string) {
	if _this.IsRain {
		var info IrnStruct
		JsonUtil.ReadJson(_this.param.IrnJson, &info)
		line = strings.Join(info.RAIN, "\n")
	} else {
		resLi := make([]string, 0)

		// 需要获取到的变量
		BCROPEAK := _this.BCROPEAK
		baseRecord9Val := 1.0

		var tmp1 string
		var tmp2, tmp3 float64
		// line 1
		tmp1 = ".0000E+00"
		tmp2 = BCROPEAK / 2
		resLi = append(resLi, fmt.Sprintf("%s %.4E", tmp1, tmp2))

		// line 2
		tmp2 = _this.param.Prec / baseRecord9Val / 2.67
		tmp3 = BCROPEAK
		resLi = append(resLi, fmt.Sprintf("%f %.4E", tmp2, tmp3))

		// line 3
		tmp2 = _this.param.Prec / baseRecord9Val
		tmp3 = BCROPEAK / 2
		resLi = append(resLi, fmt.Sprintf("%f %.4E", tmp2, tmp3))

		// line 4
		tmp1 = ".8640E+05"
		tmp2 = 0
		resLi = append(resLi, fmt.Sprintf("%s %.4E", tmp1, tmp2))

		line = strings.Join(resLi, "\n")
	}
	_this.resLi = append(_this.resLi, line)
	return line
}

func (_this *IroStruct) joinLine() (res string) {
	_this.line1()
	_this.line2()
	_this.line3()

	_this.res = strings.Join(_this.resLi, "\n")
	return _this.res
}

func (_this *IroStruct) toJson() {
	JsonUtil.WriteJson(_this.param.JsonPath, _this)
}

func GenerateIroFun(param VFSMParams.IroParams) string {
	var iro IroStruct
	iro.init(param)
	res := iro.joinLine()
	iro.toJson()

	//log.Println("iro success")
	return res
}
