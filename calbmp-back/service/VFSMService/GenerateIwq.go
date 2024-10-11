package VFSMService

import (
	"calbmp-back/Params/InputParams"
	"calbmp-back/Params/VFSMParams"
	"calbmp-back/model"
	"calbmp-back/util/JsonUtil"
	"calbmp-back/util/StringUtil"
	"fmt"
	"strings"
)

type IwqStruct struct {
	// line 1
	IWQPRO int64
	// line 2
	IKD  int64
	VKOC float64
	VKD  float64
	OCP  float64
	// line 3
	CCP float64
	// line 4
	IDG int64
	// line 5
	NDGDAY   int64
	DGHALF   string
	FC       float64
	DGPIN    float64
	DGML     float64
	DgLD     float64
	Dgmeres0 int64
	// line 6
	DgT []string
	// line 7
	DgTheta []string
	// line 8
	IMOB int64

	// params
	param         VFSMParams.IwqParams
	inp1          InputParams.UserInputStepReceiver
	inp2          InputParams.UserInputStepReceiver
	cropPesticide model.CropPesticide
	chorizonFinal model.ChorizonModel

	// result
	resLi []string
	res   string
}

func (_this *IwqStruct) init(param VFSMParams.IwqParams) {
	_this.IWQPRO = 3
	_this.IDG = 1

	_this.cropPesticide = param.CropPesticide
	_this.chorizonFinal = param.Chorizon

	_this.param = param
}

func (_this *IwqStruct) line1() (line string) {
	template := "%d"
	line = fmt.Sprintf(template, _this.IWQPRO)
	_this.resLi = append(_this.resLi, line)
	return
}

func (_this *IwqStruct) line2() (line string) {
	template := "%d %f %f"
	_this.IKD = 1
	_this.VKOC = StringUtil.Convert2Float(_this.cropPesticide.Koc)
	_this.VKD = StringUtil.Convert2Float(_this.cropPesticide.Kd)
	_this.OCP = StringUtil.Convert2Float(_this.chorizonFinal.Orgc_per)

	line = fmt.Sprintf(template, _this.IKD, _this.VKOC, _this.OCP)
	_this.resLi = append(_this.resLi, line)
	return
}

func (_this *IwqStruct) line3() (line string) {
	template := "%f"
	_this.CCP = StringUtil.Convert2Float(_this.chorizonFinal.Claytotal_per)
	line = fmt.Sprintf(template, _this.CCP)
	_this.resLi = append(_this.resLi, line)
	return
}

func (_this *IwqStruct) line4() (line string) {
	template := "%d"
	line = fmt.Sprintf(template, _this.IDG)
	_this.resLi = append(_this.resLi, line)
	return
}
func (_this *IwqStruct) line5() (line string) {
	_this.NDGDAY = _this.param.NDGDAY
	_this.DGHALF = _this.cropPesticide.Dt50_aerobic_days
	_this.FC = StringUtil.Convert2Float(_this.chorizonFinal.Wthirdbar_cm3_cm3)
	_this.DGPIN = _this.param.DGPIN
	_this.DGML = 2
	_this.DgLD = 0.05
	_this.Dgmeres0 = 0
	line = StringUtil.Var2Line(
		_this.NDGDAY,
		_this.DGHALF,
		_this.FC,
		_this.DGPIN,
		_this.DGML,
		_this.DgLD,
		_this.Dgmeres0,
	)
	_this.resLi = append(_this.resLi, line)
	return
}
func (_this *IwqStruct) line6() (line string) {
	_this.DgT = _this.param.DgT
	temp := make([]interface{}, len(_this.DgT))
	for i, v := range _this.DgT {
		temp[i] = v
	}
	line = StringUtil.Var2Line(temp...)
	_this.resLi = append(_this.resLi, line)
	return
}
func (_this *IwqStruct) line7() (line string) {
	//_this.DgTheta = make([]float64, 0)
	//for i := 0; i < int(_this.NDGDAY); i++ {
	//	_this.DgTheta = append(_this.DgTheta, 0.274)
	//}
	_this.DgTheta = _this.param.DgTheta
	temp := make([]interface{}, len(_this.DgTheta))
	for i, v := range _this.DgTheta {
		var thetaTemp float64
		_, _ = fmt.Sscanf(v, "%e", &thetaTemp)
		temp[i] = thetaTemp / StringUtil.Convert2Float(_this.chorizonFinal.Hzthk)
	}
	line = StringUtil.Var2Line(temp...)
	_this.resLi = append(_this.resLi, line)
	return
}
func (_this *IwqStruct) line8() (line string) {
	_this.IMOB = 1
	line = StringUtil.Var2Line(_this.IMOB)
	_this.resLi = append(_this.resLi, line)
	return
}

func (_this *IwqStruct) joinLine() (res string) {
	_this.line1()
	_this.line2()
	_this.line3()
	_this.line4()
	_this.line5()
	_this.line6()
	_this.line7()
	_this.line8()

	_this.res = strings.Join(_this.resLi, "\n")
	return _this.res
}

func (_this *IwqStruct) toJson() {
	JsonUtil.WriteJson(_this.param.JsonPath, _this)
}

func GenerateIwqFun(param VFSMParams.IwqParams) string {
	var iwq IwqStruct
	iwq.init(param)
	res := iwq.joinLine()
	iwq.toJson()

	//log.Println("iwq success")
	return res
}
