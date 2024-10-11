package TimeUtil

import (
	"time"
)

func GetDaysBetween2Date(format, date1Str, date2Str string) (int64, error) {
	// 将字符串转化为Time格式
	date1, err := time.ParseInLocation(format, date1Str, time.Local)
	if err != nil {
		return 0, err
	}
	// 将字符串转化为Time格式
	date2, err := time.ParseInLocation(format, date2Str, time.Local)
	if err != nil {
		return 0, err
	}
	//计算相差天数
	return int64(date2.Sub(date1).Hours() / 24), nil
}
