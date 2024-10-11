package TimeUtil

import (
	"log"
	"time"
)

func GetDOY(layout, date string) int {
	dateObj, err := time.Parse(layout, date)
	if err != nil {
		log.Println("[GetDOY]:", date, "->err:", err)
	}
	return dateObj.YearDay()
}
