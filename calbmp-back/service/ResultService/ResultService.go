package ResultService

import (
	"archive/zip"
	"calbmp-back/Params/ResultParams"
	"calbmp-back/Repository/CropPesticideFinalRepository"
	"calbmp-back/dto/ResultDataDTO"
	"calbmp-back/util/ResultUtil"
	"calbmp-back/util/StringUtil"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

// GetDataByName : get common visualization data
func GetDataByName(rec ResultParams.ResultReceiver, dataType string) ResultDataDTO.ResData {
	// locate zts file
	scenarioType := rec.ScenarioType
	BmpId := rec.BmpId
	CreatedTime := rec.CreatedTime

	// init result data
	var resData ResultDataDTO.ResData
	resData.Date = make([]string, 0)
	resData.Water = make(map[string][]float64)
	resData.Pesticide = make(map[string][]float64)
	resData.Sediment = make(map[string][]float64)
	resData.Concentration = make(map[string][]float64)

	// get path of zts file
	ztsPath := ResultUtil.LocateZtsPath(scenarioType, CreatedTime, BmpId)
	// get matrix
	ztsData := ResultUtil.GetMatrix(ztsPath)
	// get date data
	resData.Date = ResultUtil.GetDateSet(ztsData)
	// get water
	resData.Water = ResultUtil.GetWaterDataSet(ztsData, rec.Water, dataType)
	// get pesticide data set
	resData.Pesticide = ResultUtil.GetPesticideDataSet(ztsData, rec.PesticideList, rec.Pesticide, rec.FieldSize, dataType)
	// get sediment data set
	resData.Sediment = ResultUtil.GetSedimentDataSet(ztsData, rec.Sediment, dataType)
	// get concentration data set
	resData.Concentration = ResultUtil.GetConcentrationDataSet(ztsData, rec.PesticideList, dataType)

	// get benchmark value
	resData.Benchmark = make(map[string]float64)
	for i, v := range rec.PesticideList {
		BenchmarkValue := CropPesticideFinalRepository.FindBenchMarkValueByChemicalName(v)
		//resData.Benchmark = append(resData.Benchmark, BenchmarkValue)
		key := fmt.Sprintf("%s_%d", v, i+1)
		BenchmarkValue *= 1000 // ppm to ug/L
		resData.Benchmark[key] = BenchmarkValue
	}

	return resData
}

// GetComparisonData : get comparison data
func GetComparisonData(
	BmpCnt int,
	HarvestData, CreatedTime string,
) ([]float64, []float64, []float64) {
	// get baseline data and global column list
	columnNames, dataLi := ResultUtil.GetBasicResult("baseline", "0", CreatedTime)

	// get month/day/year list
	MonthDayYearList := make([]string, 0)
	for _, v := range dataLi {
		tmpDate := fmt.Sprintf("%s/%s/%s", v[1], v[2], v[0])
		MonthDayYearList = append(MonthDayYearList, tmpDate)
	}

	// get format(month/day/year) of harvest
	dateSplit := strings.Split(HarvestData, "-")
	HYear := StringUtil.DeleteFrontZero(dateSplit[0])
	HMonth := StringUtil.DeleteFrontZero(dateSplit[1])
	HDay := StringUtil.DeleteFrontZero(dateSplit[2])
	HarvestString := fmt.Sprintf("%s/%s/%s", HMonth, HDay, HYear)

	// get index of harvest date
	HIndex := 0
	for i, v := range MonthDayYearList {
		if HarvestString == v {
			HIndex = i
			break
		}
	}

	// get index of runoff and erosion
	Params := []string{"RFLX", "EFLX", "VFLX"}
	Indexes := ResultUtil.ParseName(Params, columnNames)

	// get baseline target data
	BRunoff := ResultUtil.ConvertToFloat64(dataLi[HIndex][Indexes[1]])
	BErosion := ResultUtil.ConvertToFloat64(dataLi[HIndex][Indexes[3]])
	BVolatilization := ResultUtil.ConvertToFloat64(dataLi[HIndex][Indexes[5]])

	// get basic data by bmp id
	ComparisonRunoff := make([]float64, 0)
	ComparisonErosion := make([]float64, 0)
	ComparisonVolatilization := make([]float64, 0)
	for i := 1; i <= BmpCnt; i++ {
		BmpId := fmt.Sprintf("%d", i)
		_, bmpDataLi := ResultUtil.GetBasicResult("bmp", BmpId, CreatedTime)

		CRunoff := ResultUtil.ConvertToFloat64(bmpDataLi[HIndex][Indexes[1]])
		CErosion := ResultUtil.ConvertToFloat64(bmpDataLi[HIndex][Indexes[3]])
		CVolatilization := ResultUtil.ConvertToFloat64(bmpDataLi[HIndex][Indexes[5]])

		// formula : abs(bmp的累积的harvest date的值 - baseline的累积的harvest date的值) /  baseline的累积的值

		var CMPRRunoff float64
		var CMPRErosion float64
		var CMPRVolatilization float64
		if BRunoff != 0 {
			CMPRRunoff = math.Abs(CRunoff-BRunoff) / BRunoff * 100
		} else {
			CMPRRunoff = 0
		}
		if BErosion != 0 {
			CMPRErosion = math.Abs(CErosion-BErosion) / BErosion * 100
		} else {
			CMPRErosion = 0
		}
		if BVolatilization != 0 {
			CMPRVolatilization = math.Abs(CVolatilization-BVolatilization) / BVolatilization * 100
		} else {
			CMPRVolatilization = 0
		}

		ComparisonRunoff = append(ComparisonRunoff, CMPRRunoff)
		ComparisonErosion = append(ComparisonErosion, CMPRErosion)
		ComparisonVolatilization = append(ComparisonVolatilization, CMPRVolatilization)
	}

	return ComparisonRunoff, ComparisonErosion, ComparisonVolatilization
}

func GetTextResultFun(params ResultParams.TextResultParams, dataType string) ResultDataDTO.ResData {
	rec := ResultParams.ResultReceiver{
		ScenarioType:  params.ScenarioType,
		BmpId:         "0",
		Water:         []string{"RUNF", "PRCP", "IRRG", "SWTR"},
		Pesticide:     []string{"RFLX", "EFLX", "VFLX"},
		Sediment:      []string{"ESLS"},
		Concentration: []string{"concentration"},
		PesticideList: params.PesticideList,
		CreatedTime:   params.CreatedTime,
		FieldSize:     params.FieldSize,
	}

	// get basic result
	basicRes := GetDataByName(rec, dataType)

	// get runf result data
	basicRunf := basicRes.Water["RUNF"]

	// get units by runf != 0
	unitCnt := 0
	unitIdxList := make([]int, 0)
	for i, v := range basicRunf {
		if v != 0 {
			unitCnt++
			unitIdxList = append(unitIdxList, i)
		}
	}

	// init res data
	res := ResultDataDTO.ResData{
		Date:          make([]string, unitCnt),
		Water:         make(map[string][]float64),
		Pesticide:     make(map[string][]float64),
		Sediment:      make(map[string][]float64),
		Concentration: make(map[string][]float64),
		Benchmark:     make(map[string]float64),
	}

	// get data by index
	for i, idx := range unitIdxList {
		res.Date[i] = basicRes.Date[idx]
		for key, element := range basicRes.Water {
			res.Water[key] = append(res.Water[key], element[idx])
		}
		for key, element := range basicRes.Pesticide {
			res.Pesticide[key] = append(res.Pesticide[key], element[idx])
		}
		for key, element := range basicRes.Sediment {
			res.Sediment[key] = append(res.Sediment[key], element[idx])
		}
		for key, element := range basicRes.Concentration {
			res.Concentration[key] = append(res.Concentration[key], element[idx])
		}
	}
	res.Benchmark = basicRes.Benchmark

	return res
}

func GetInputFilesFun(CreatedTime string) (files []map[string]string) {
	files = make([]map[string]string, 0)

	basePath := "./przm5place/"
	targetPath := basePath + CreatedTime

	errWalk := filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
		if info.Name() != "down.zip" && !info.IsDir() {
			temp := make(map[string]string)
			temp["filename"] = path
			files = append(files, temp)
		}
		return nil
	})
	if errWalk != nil {
		return nil
	}

	return
}

func ZipInputFileFun(param ResultParams.ZipInputFileParam) {
	basePath := "./przm5place/"
	zipFilePath := basePath + param.CreatedTime + "/" + "down.zip"
	fzip, _ := os.Create(zipFilePath)
	zipWriter := zip.NewWriter(fzip)

	for _, v := range param.FileList {
		w, errCreate := zipWriter.Create(v)
		if errCreate != nil {
			log.Println(errCreate)
		}

		f, errOpen := os.Open(v)
		if errOpen != nil {
			log.Println(errOpen)
		}
		if _, errCopy := io.Copy(w, f); errCopy != nil {
			log.Println(errCopy)
		}
	}
	zipWriter.Close()
}
