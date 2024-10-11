package InputFileGenerateUtil

import (
	"calbmp-back/Params/InputParams"
	"fmt"
	"strconv"
	"strings"
)

// GenerateDateListWithCN :
// return day, month, year, cn, cn flag;
// cn flag : 0 -> bare value; 1 -> cn value
func GenerateDateListWithCN(
	rec InputParams.UserInputStepReceiver,
	CNValueStr string,
	BareValueStr string,
) [][]int {
	// struct:
	// 0. index of values
	// 1. day, month, year, cn, cn flag;
	DateLiWithCN := make([][]int, 0)

	// convert CNValue and BareValue to int
	CNValue, _ := strconv.Atoi(CNValueStr)
	BareValue, _ := strconv.Atoi(BareValueStr)

	// get basic date value
	Emergence := rec.Emergence
	Harvest := rec.Harvest

	// init date value
	// 1. split date string
	EmergenceLi := strings.Split(Emergence, "-")
	HarvestLi := strings.Split(Harvest, "-")

	// 2. get year, month and day of emergence
	StartYear, _ := strconv.Atoi(EmergenceLi[0])
	EmergenceMonth, _ := strconv.Atoi(EmergenceLi[1])
	EmergenceDay, _ := strconv.Atoi(EmergenceLi[2])

	// 3. get year, month and day of harvest
	StopYear, _ := strconv.Atoi(HarvestLi[0])
	HarvestMonth, _ := strconv.Atoi(HarvestLi[1])
	HarvestDay, _ := strconv.Atoi(HarvestLi[2])

	// generate date list
	/*
		flag = true  -> day = 16
		flag = false -> day = 1
	*/
	flag := true   // control day
	CNFlag := true // cn value or bare soil
	// cnt control day of a year,
	// a year have 2 days
	cnt := 0

	curDay := EmergenceDay
	curMonth := EmergenceMonth
	curYear := StartYear

	for i := 0; i < 24; i++ {
		// Check the CNFlag and decide which CN to use
		// CN Value:
		// CNFlag = true  -> CNValue
		// CNFlag = false -> BareValue
		if CNFlag == true {
			DateLiWithCN = append(
				DateLiWithCN,
				[]int{
					curDay,
					curMonth,
					curYear,
					CNValue,
					1,
				},
			)
		} else {
			DateLiWithCN = append(
				DateLiWithCN,
				[]int{
					curDay,
					curMonth,
					curYear,
					BareValue,
					0,
				},
			)
		}

		// check day
		if flag == true {
			curDay = 16
		} else {
			curDay = 1
		}

		// Change day from 1 to 16
		// and then increase month by one
		if cnt >= 1 {
			curMonth = curMonth%12 + 1

			// If curMonth equals 1,
			// represent for next year
			if curMonth == 1 {
				curYear++
			}
			// reset cnt to 0
			cnt = 0
		} else {
			cnt++
		}

		// check harvest date to control CNFlag
		if curYear == StopYear {
			if curMonth > HarvestMonth {
				CNFlag = false
			} else if curMonth == HarvestMonth && curDay >= HarvestDay {
				CNFlag = false
			}
		} else if curYear > StopYear {
			CNFlag = false
		}

		flag = !flag
	}

	return DateLiWithCN
}

func MakeRecord5line(
	day,
	month,
	year,
	cn int,
	c,
	n float64,
) string {
	// day, month, year, c = uslec, n, cn
	temp := "%2d, %2d, %4d,   %.5f, %.5f, %d"

	line := fmt.Sprintf(temp, day, month, year, c, n, cn)

	return line
}
