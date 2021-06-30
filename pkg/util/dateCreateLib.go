package util

import "time"

//获取本月第一天
func GetMonthFirstDay() string {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	monthFirstDay := thisMonth.Format("2006-01-02")
	return monthFirstDay

}

//获取今天年月日
func GetToday() (int, time.Month, int) {
	year, month, day := time.Now().Date()
	return year, month, day
}

//获取今天日期year,month,day,format
func GetTodayDateStr() string {
	year, month, day := time.Now().Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	DateStr := today.Format("2006-01-02")
	return DateStr
}

//获取当前日期偏移format
func GetNowTimeOffset(year int, month int, day int) string {
	currentTime := time.Now()
	newTime := currentTime.AddDate(year, month, day)
	DateStr := newTime.Format("2006-01-02")
	return DateStr
}
