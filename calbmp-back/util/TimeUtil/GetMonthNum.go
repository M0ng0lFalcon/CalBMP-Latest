package TimeUtil

import (
	"log"
	"time"
)

func ParseTimeString(date string) time.Time {
	const TimeLayout = "2006-01-02"
	dateObj, err := time.Parse(TimeLayout, date)
	if err != nil {
		log.Println(err)
	}

	return dateObj
}
