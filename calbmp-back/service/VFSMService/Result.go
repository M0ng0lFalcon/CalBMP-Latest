package VFSMService

import (
	"calbmp-back/util/FileUtil"
	"calbmp-back/util/StringUtil"
	"fmt"
	"log"
	"strings"
)

// ParseOsm
//
//	@Description: 解析Osm文件
//	@param CreatedTime 项目创建时间
//	@param MinId 最小的VFS id
//	@param MaxId 最大的VFS id
//	@return osm 解析结果
func ParseOsm(CreatedTime string, MinId, MaxId int) (osm map[string][]float64) {
	// 初始化
	outputDir := "./output"
	osm = make(map[string][]float64)
	osm["runoff"] = make([]float64, 0)   // 单位：m3
	osm["sediment"] = make([]float64, 0) // 单位：g

	// 逐文件进行解析
	for i := MinId; i <= MaxId; i++ {
		// 构建文件名
		filename := fmt.Sprintf("%s/%s_%d.osm", outputDir, CreatedTime, i)

		// 读取文件
		var lines []string
		err := FileUtil.ReadFile(filename, &lines)
		if err != nil {
			log.Println(err)
		}

		// 逐行解析
		for _, line := range lines {
			// 利用空白进行分割
			li := strings.Fields(line)

			// 解析
			if strings.HasPrefix(line, "Volume from outflow") { // runoff
				temp := StringUtil.Convert2Float(li[len(li)-2])
				osm["runoff"] = append(osm["runoff"], temp)
			} else if strings.HasPrefix(line, "Sediment outflow") { // sediment
				temp := StringUtil.Convert2Float(li[len(li)-1])
				osm["sediment"] = append(osm["sediment"], temp)
			}
		}
	}
	return
}

func ParseOwq(CreatedTime string, MinId, MaxId int) (owq map[string][]float64) {
	outputDir := "./output"
	owq = make(map[string][]float64)
	owq["overall"] = make([]float64, 0)
	owq["PesticideInput"] = make([]float64, 0)
	owq["PesticideOutput"] = make([]float64, 0)
	owq["PesticideOutflowInLiquidPhase"] = make([]float64, 0)
	owq["PesticideOutflowInSolidPhase"] = make([]float64, 0)
	owq["PesticideTrappedInVFS"] = make([]float64, 0)
	for i := MinId; i <= MaxId; i++ {
		filename := fmt.Sprintf("%s/%s_%d.owq", outputDir, CreatedTime, i)
		var lines []string
		err := FileUtil.ReadFile(filename, &lines)
		if err != nil {
			log.Println(err)
		}

		flagLi := []bool{false, false, false, false, false}

		for _, line := range lines {
			li := strings.Fields(line)
			if strings.HasSuffix(line, "Pesticide input (mi)") && !flagLi[0] {
				temp := StringUtil.Convert2Float(li[0])
				owq["PesticideInput"] = append(owq["PesticideInput"], temp)
				flagLi[0] = true
			} else if strings.HasSuffix(line, "Pesticide output (mo)") && !flagLi[1] {
				temp := StringUtil.Convert2Float(li[0])
				owq["PesticideOutput"] = append(owq["PesticideOutput"], temp)
				flagLi[1] = true
			} else if strings.HasSuffix(line, "Pesticide outflow in liquid phase (mod)") && !flagLi[2] {
				temp := StringUtil.Convert2Float(li[0])
				owq["PesticideOutflowInLiquidPhase"] = append(owq["PesticideOutflowInLiquidPhase"], temp)
				flagLi[2] = true
			} else if strings.HasSuffix(line, "Pesticide outflow in solid phase (mop)") && !flagLi[3] {
				temp := StringUtil.Convert2Float(li[0])
				owq["PesticideOutflowInSolidPhase"] = append(owq["PesticideOutflowInSolidPhase"], temp)
				flagLi[3] = true
			} else if strings.HasSuffix(line, "Pesticide trapped in VFS (mf)") && !flagLi[4] {
				temp := StringUtil.Convert2Float(li[0])
				owq["PesticideTrappedInVFS"] = append(owq["PesticideTrappedInVFS"], temp)
				flagLi[4] = true
			}
		}
	}

	pesticideOutput := 0.0
	pesticideInput := 0.0
	for _, v := range owq["PesticideInput"] {
		pesticideInput += v
	}
	for _, v := range owq["PesticideOutput"] {
		pesticideOutput += v
	}
	overall := 0.0
	if pesticideInput != 0 {
		overall = (pesticideInput - pesticideOutput) / pesticideInput
	} else {
		overall = -1
	}
	owq["overall"] = append(owq["overall"], overall)
	return
}
