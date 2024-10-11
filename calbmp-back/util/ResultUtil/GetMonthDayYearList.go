package ResultUtil

import "fmt"

func GetMonthDayYearList(dataLi [][]string) []string {
	MonthDayYearList := make([]string, 0)
	for _, v := range dataLi {
		tmpDate := fmt.Sprintf("%s/%s/%s", v[1], v[2], v[0])
		MonthDayYearList = append(MonthDayYearList, tmpDate)
	}

	return MonthDayYearList
}
