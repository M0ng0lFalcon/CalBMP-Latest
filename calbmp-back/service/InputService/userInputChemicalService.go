package InputService

import (
	"calbmp-back/Params/InputParams"
	"calbmp-back/Repository/ChorizonRepository"
	"calbmp-back/Repository/CropPesticideFinalRepository"
	"calbmp-back/util/StringUtil"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetRecordC1(rec InputParams.UserInputStepReceiver) string {
	res := ""
	// request params
	noOfApp := rec.NoOfApp
	ifBmp := rec.IfBMP
	// Number of applications
	cntPesticide := rec.CntPesticide

	var ifBmpStr string
	if ifBmp == false {
		ifBmpStr = "False"
	} else {
		ifBmpStr = "True"
	}

	temp := "%d, %d, %s, 5, 2, 0, 0"
	res = fmt.Sprintf(temp, noOfApp, cntPesticide, ifBmpStr)
	return res
}

func GetRecordC2(rec InputParams.UserInputStepReceiver) string {
	res := ""

	// get some params
	Date := rec.Date
	NoOfApp := rec.NoOfApp
	cam := rec.ApplicationMethod
	depthLi := rec.Depth
	depi := make([]float64, len(cam))
	Amount := rec.Amount
	ApplicationEquipment := rec.ApplicationEquipment

	// set depi
	for i, v := range cam {
		if v >= 1 && v <= 3 {
			depi[i] = 0
		} else {
			depi[i] = depthLi[i]
		}
	}

	/**
	0.99 = Ground-applied Sprayer
	0.95 = Ariel sprayer
	1    = Others
	*/
	drft := make([]float64, len(ApplicationEquipment))
	for i, sqjRate := range ApplicationEquipment {
		if sqjRate == 0.95 {
			drft[i] = 0.05
		} else if sqjRate == 1 {
			drft[i] = 0.0
		} else {
			drft[i] = 0.01
		}
	}

	// 当cam=7时，tband参数有意义，drft即tband
	for i, v := range cam {
		if v == 7 {
			drft[i] = 0.8
		}
	}

	// date(dd, mm, yy), cam, depi, amount, sqjRate, drft
	temp := "%s,%d,%f,%f,%f,%f"

	for i := 0; i < NoOfApp; i++ {
		if i != 0 {
			res += "\n"
		}
		// format date
		DateLi := strings.Split(Date[i], "-")
		Year := DateLi[0]
		Month := DateLi[1]
		Day := DateLi[2]
		dateVar := fmt.Sprintf("%s,%s,%s",
			StringUtil.DeleteFrontZero(Day),
			StringUtil.DeleteFrontZero(Month),
			Year)

		// 转换单位
		amountFloat, _ := strconv.ParseFloat(Amount[i], 10)
		amountFloat *= 1.12085116

		res += fmt.Sprintf(temp,
			dateVar,                 // date 1,2,3
			cam[i],                  // cam 4
			depi[i],                 // depi 5
			amountFloat,             // amount 6
			ApplicationEquipment[i], // sqjRate 7
			drft[i],                 // drft 8
		)
	}

	return res
}

func GetRecordC3(rec InputParams.UserInputStepReceiver) string {
	res := ""

	// get params
	pesticide := rec.PesticideSet

	for i, ChemicalName := range pesticide {
		if i != 0 {
			res += ","
		}
		// get crop pesticide to find out logP
		CropPesticide := CropPesticideFinalRepository.FindCropPesticideByChemicalName(ChemicalName)
		logP := StringUtil.Convert2Float(CropPesticide.Logp_logarithm_of_octanol_water_partition_coefficient)

		// calculate uptkf by formula
		uptkf := 0.789 * math.Exp(-(((logP - 1.78) * (logP - 1.78)) / 2.44))
		// convert float to string
		uptkfStr := strconv.FormatFloat(uptkf, 'g', 20, 64)
		// concat res
		res += uptkfStr
	}

	return res
}

func GetRecordC4(rec InputParams.UserInputStepReceiver) string {
	res := ""

	temp := "0, %s, 0.5"

	// get params
	pesticide := rec.Pesticide
	cam := rec.ApplicationMethod

	cnt := 0
	for i, v := range cam {
		if v == 2 || v == 3 {
			if cnt != 0 {
				res += "\n"
			}
			// get crop pesticide by pesticide name
			cp := CropPesticideFinalRepository.FindCropPesticideByChemicalName(pesticide[i])

			line := fmt.Sprintf(temp, cp.One_dt50_foliar_days_1)
			res += line

			cnt++
		}
	}

	if cnt > 0 {
		res = "***Record C4 (Chem #1)\n" + res
	} else {
		res = "***Record C4 (Chem #1)"
	}
	return res
}

func GetRecordC5(rec InputParams.UserInputStepReceiver) string {
	res := "***Record C5"

	// params
	pesticideCnt := rec.CntPesticide
	cam := rec.ApplicationMethod

	flag := false

	if pesticideCnt > 1 {
		for _, v := range cam {
			if v == 2 {
				flag = true
			}
		}
	}

	if flag {
		res += "\n1,1,1"
	}
	return res
}

func GetRecordC6(rec InputParams.UserInputStepReceiver) string {
	res := ""
	temp := "4300,%s,%s,"
	pesticide := rec.PesticideSet
	//cropName := rec.Crop
	for i, ChemicalName := range pesticide {
		if i != 0 {
			res += "\n"
		}

		cp := CropPesticideFinalRepository.FindCropPesticideByChemicalName(ChemicalName)

		//fmt.Println("cp:", cp.Henry_num, cp.Enthalpy_num_kcal_mol)

		res += fmt.Sprintf(temp, cp.Henry_num, cp.Enthalpy_num_kcal_mol)
	}
	return res
}

// GetRecordC7 : get record 7 value
func GetRecordC7(rec InputParams.UserInputStepReceiver) string {
	res := ""

	// get params
	MuKey := rec.MuKey
	CoKey := rec.CoKey
	Pesticide := rec.Pesticide

	cpp := ChorizonRepository.FindChorizonListByMukeyAndCokey(MuKey, CoKey)
	// len of horizon
	cnt := len(cpp)

	KfKdList := make([]string, 0)
	for i := 0; i < len(Pesticide); i++ {
		cp := CropPesticideFinalRepository.FindCropPesticideByChemicalName(Pesticide[i])

		kfKd := StringUtil.Convert2Float(cp.Freundlich_kf_or_kd)
		if kfKd == 0 {
			// find max orgC
			orgC := -1.0
			for j := 0; j < cnt; j++ {
				if orgC < StringUtil.Convert2Float(cpp[j].Orgc_per) {
					orgC = StringUtil.Convert2Float(cpp[j].Orgc_per)
				}
			}
			// calculate kf_kd
			kfKd = (StringUtil.Convert2Float(cp.Freundlich_kfoc) * orgC) / 100.0
		}
		// convert to string
		KfKdStr := fmt.Sprintf("%f", kfKd)
		// append to list
		KfKdList = append(KfKdList, KfKdStr)
	}

	line := strings.Join(KfKdList, ",")
	for i := 0; i < cnt; i++ {
		if i != 0 {
			res += "\n"
		}
		res += line
	}
	return res
}

func GetRecordC7A(rec InputParams.UserInputStepReceiver) string {
	res := ""

	// get params
	pesticideList := rec.PesticideSet
	muKey := rec.MuKey
	coKey := rec.CoKey

	cpp := ChorizonRepository.FindChorizonListByMukeyAndCokey(muKey, coKey)
	cnt := len(cpp)

	FreundlichList := make([]string, 0)
	for _, v := range pesticideList {
		cp := CropPesticideFinalRepository.FindCropPesticideByChemicalName(v)
		Freundlich := fmt.Sprintf("%s", cp.Freundlich_1_n)
		FreundlichList = append(FreundlichList, Freundlich)
	}
	line := strings.Join(FreundlichList, ",")
	for i := 0; i < cnt; i++ {
		if i != 0 {
			res += "\n"
		}
		res += line
	}

	return res
}

func GetRecordC7B(rec InputParams.UserInputStepReceiver) string {
	res := ""

	// get params
	cntPesticide := rec.CntPesticide
	Mukey := rec.MuKey
	Cokey := rec.CoKey

	cpp := ChorizonRepository.FindChorizonListByMukeyAndCokey(Mukey, Cokey)
	cnt := len(cpp)

	temp := ""
	for i := 0; i < cntPesticide; i++ {
		temp += "1,"
	}

	for i := 0; i < cnt; i++ {
		if i != 0 {
			res += "\n"
		}
		res += temp
	}

	return res
}

func GetRecordC7C(rec InputParams.UserInputStepReceiver) string {
	return GetRecordC7B(rec)
}

func GetRecordC7D() string {
	return "1.0e-12,"
}

func GetRecordC7E(rec InputParams.UserInputStepReceiver) string {
	res := ""
	// get params
	cntPesticide := rec.CntPesticide

	for i := 0; i < cntPesticide; i++ {
		res += "1,"
	}
	return res
}

func GetRecordC8(rec InputParams.UserInputStepReceiver) string {
	res := ""

	// get params
	Mukey := rec.MuKey
	Cokey := rec.CoKey
	pesticide := rec.PesticideSet

	cpp := ChorizonRepository.FindChorizonListByMukeyAndCokey(Mukey, Cokey)
	cnt := len(cpp)

	temp := ""
	for _, v := range pesticide {
		cp := CropPesticideFinalRepository.FindCropPesticideByChemicalName(v)
		OneDt50Aerobic := cp.One_dt50_aerobic_days_1
		temp += OneDt50Aerobic + "," + OneDt50Aerobic + ",0,"
	}

	for i := 0; i < cnt; i++ {
		if i != 0 {
			res += "\n"
		}
		res += temp
	}

	return res
}

func GetRecordC9(rec InputParams.UserInputStepReceiver) string {
	res := ""
	Mukey := rec.MuKey
	Cokey := rec.CoKey

	cpp := ChorizonRepository.FindChorizonListByMukeyAndCokey(Mukey, Cokey)
	cnt := len(cpp)

	for i := 0; i < cnt; i++ {
		if i != 0 {
			res += "\n"
		}
		res += "0,0,0,0,0,0"
	}
	return res
}
