package utils

import (
	"time"
)

/**
 * @Author:lingwang
 * @File :日期转换方法
 * @Description:
 * @Version: 1.0.0
 * @Date :2021/12/27 11:36
 */
const (
	TIMEFORMAT       = "20060102150405"
	NORMALTIMEFORMAT = "2006-01-02 15:04:05"
	DAYTIMEFORMAT    = "2006-01-02"
	HOURFORMAR       = "15:03:04"
)

/** 获取当前时间 */
func GetTime() time.Time {
	return time.Now()
}

/** 格式化时间 */
func GetTimeString(time time.Time) string {
	return time.Format(TIMEFORMAT)
}

/** 格式化时间 */
func GetTimeToFormat(time time.Time, format string) string {
	return time.Format(format)
}

/** 时间戳->秒 */
func GetTimeUnix(time time.Time) int64 {
	return time.Unix()
}

/** 时间戳->毫秒 */
func GetTimeMills(time time.Time) int64 {
	return time.Unix() / 1e6
}

/** 时间戳转时间 */
func GetTimeByInt(t int64) time.Time {
	return time.Unix(t, 0)
}

/** 字符串转时间 */
func GetTimeByString(timeString string) (time.Time, error) {
	if timeString == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(TIMEFORMAT, timeString, time.Local)
}

/** 字符串转时间-自定义 */
func GetTimeByStringToFormat(timeStirng string, format string) (time.Time, error) {
	if timeStirng == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(format, timeStirng, time.Local)
}

/** 比较两个时间大小 */
func CompareTime(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

/** 判断两个时间是否相等 */
func EqualTime(t1, t2 time.Time) bool {
	return t1.Equal(t2)
}

/** n小时后的时间字符串 */
func GetNextHourTime(t string, n int64) string {
	t2, _ := time.ParseInLocation(TIMEFORMAT, t, time.Local)
	t1 := t2.Add(time.Hour * time.Duration(n))
	return GetTimeString(t1)
}

/** n天、月、年后的时间字符串 */
func GetNextTime(t string, n, m, y int) string {
	t2, _ := time.ParseInLocation(TIMEFORMAT, t, time.Local)
	t1 := t2.AddDate(n, m, y)
	return GetTimeString(t1)
}

/** 计算两个时间查N小时 */
func GetHourDiffer(start_time, end_time string) float32 {
	var hour float32
	t1, err1 := time.ParseInLocation(TIMEFORMAT, start_time, time.Local)
	t2, err2 := time.ParseInLocation(TIMEFORMAT, end_time, time.Local)
	if err1 == nil && err2 == nil && (!EqualTime(t1, t2)) {
		diff := GetTimeUnix(t1) - GetTimeUnix(t2)
		hour = float32(diff) / 3600
		return hour
	}
	return hour

}

/** 判断整点 */
func CheckHours(t time.Time) bool {
	_, m, s := t.Clock()
	if m == s && m == 0 && s == 0 {
		return true
	}
	return false
}
