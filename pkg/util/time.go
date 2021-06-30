package util

import "time"

var (
	TimeLayoutDate     = "2006-01-02"
	TimeLayoutDateTime = "2006-01-02 15:04:05"
)

// MonthStart 每月的开始时间
func MonthStart() time.Time {
	y, m, _ := time.Now().Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
}

// TodayStart 当前日的开始时间
func TodayStart() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

// TodayEnd 当前日的结束时间
func TodayEnd() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 23, 59, 59, 1E9-1, time.Local)
}

// NowUnix 系统当前时间
func NowUnix() int64 {
	return time.Now().Unix()
}

// NowDate 当前日期
func NowDate() string {
	return time.Now().Format(TimeLayoutDate)
}

// NowDateTime 当前日期
func NowDateTime() string {
	return time.Now().Format(TimeLayoutDateTime)
}

// ParseDate 格式化时间
func ParseDate(dt string) (time.Time, error) {
	return time.Parse(TimeLayoutDate, dt)
}

// ParseDateTime 格式化时间(string -> time)
func ParseDateTime(dt string) (time.Time, error) {
	return time.Parse(TimeLayoutDateTime, dt)
}

// ParseStringTime 格式化时间(string -> time)
func ParseStringTime(tm, lc string) (time.Time, error) {
	loc, err := time.LoadLocation(lc)
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(TimeLayoutDateTime, tm, loc)
}

// GetHourDiff 获取相差小时
func GetHourDiff(StartTime, EndTime string) (int64, error) {
	var hour int64
	t1, err := time.ParseInLocation(TimeLayoutDateTime, StartTime, time.Local)
	t2, err := time.ParseInLocation(TimeLayoutDateTime, EndTime, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
	}
	return hour, nil
}

// GetDaysDiff 获取相差天数
func GetDaysDiff(StartTime, EndTime string) (int, error) {
	var day int
	t1, err := time.ParseInLocation(TimeLayoutDateTime, StartTime, time.Local)
	t2, err := time.ParseInLocation(TimeLayoutDateTime, EndTime, time.Local)
	if err == nil && t1.Before(t2) {
		day = int(t1.Sub(t2))
	}
	return day, err
}
// GetDaysDiff 获取相差小时
func GetDiffDay(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int(t1.Sub(t2).Hours() / 24)
}
