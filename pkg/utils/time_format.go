/*
** description("").
** copyright('tuoyun,www.tuoyun.net').
** author("fg,Gordon@tuoyun.net").
** time(2021/2/22 11:52).
 */
package utils

import (
	"fmt"
	"strconv"
	"time"
)

const (
	TimeOffset = 8 * 3600  //8 hour offset
	HalfOffset = 12 * 3600 //Half-day hourly offset
)

// Get the current timestamp by Second
func GetCurrentTimestampBySecond() int64 {
	return time.Now().Unix()
}

// Convert timestamp to time.Time type
func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}

// Convert nano timestamp to time.Time type
func UnixNanoSecondToTime(nanoSecond int64) time.Time {
	return time.Unix(0, nanoSecond)
}
func UnixMillSecondToTime(millSecond int64) time.Time {
	return time.Unix(0, millSecond*1e6)
}

// Get the current timestamp by Nano
func GetCurrentTimestampByNano() int64 {
	return time.Now().UnixNano()
}

// Get the current timestamp by Mill
func GetCurrentTimestampByMill() int64 {
	return time.Now().UnixNano() / 1e6
}

// Get the timestamp at 0 o'clock of the day
func GetCurDayZeroTimestamp() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	return t.Unix() - TimeOffset
}

// Get the timestamp at 12 o'clock on the day
func GetCurDayHalfTimestamp() int64 {
	return GetCurDayZeroTimestamp() + HalfOffset

}

// Get the formatted time at 0 o'clock of the day, the format is "2006-01-02_00-00-00"
func GetCurDayZeroTimeFormat() string {
	return time.Unix(GetCurDayZeroTimestamp(), 0).Format("2006-01-02_15-04-05")
}

// Get the formatted time at 12 o'clock of the day, the format is "2006-01-02_12-00-00"
func GetCurDayHalfTimeFormat() string {
	return time.Unix(GetCurDayZeroTimestamp()+HalfOffset, 0).Format("2006-01-02_15-04-05")
}
func GetTimeStampByFormat(datetime string) string {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.Unix()
	return strconv.FormatInt(timestamp, 10)
}

func TimeStringFormatTimeUnix(timeFormat string, timeSrc string) int64 {
	tm, _ := time.Parse(timeFormat, timeSrc)
	return tm.Unix()
}

func TimeStringToTime(timeString string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", timeString)
	return t, err
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02")
}

func GetDateTimeBeginTimeAndEndTime() (time.Time, time.Time) {
	//1.获取当前时区
	loc, _ := time.LoadLocation("Local")
	loc = loc

	//2.今日日期字符串
	date := time.Now().Format("2006-01-02")

	//3.拼接成当天0点时间字符串
	startDate := date + " 00:00:00.000"
	//得到0点日期 2021-04-24 00:00:00 +0800 CST
	startTime, _ := time.Parse("2006-01-02 15:04:05.000", startDate)
	fmt.Println(startTime)
	//4.拼接成当天23点时间字符串
	endDate := date + " 23:59:59.999"
	//得到23点日期 2021-04-24 23:59:59 +0800 CST
	endTime, _ := time.Parse("2006-01-02 15:04:05.000", endDate)
	fmt.Println(endTime)

	return startTime, endTime
}

func GetDateTimeBeginTimeAndEndTimeByInputTime(time2 time.Time) (time.Time, time.Time) {

	//2.今日日期字符串
	date := time2.Format("2006-01-02")

	//3.拼接成当天0点时间字符串
	startDate := date + " 00:00:00.000"
	//得到0点日期 2021-04-24 00:00:00 +0800 CST
	startTime, _ := time.Parse("2006-01-02 15:04:05.000", startDate)
	fmt.Println(startTime)
	//4.拼接成当天23点时间字符串
	endDate := date + " 23:59:59.999"
	//得到23点日期 2021-04-24 23:59:59 +0800 CST
	endTime, _ := time.Parse("2006-01-02 15:04:05.000", endDate)
	fmt.Println(endTime)

	return startTime, endTime
}

// 取当前时间到其他时间的差值
func SubDemo(ts string) (error, time.Duration) {
	now := time.Now()
	_, err := time.Parse("2006-01-02 15:04:05", ts)
	if err != nil {
		fmt.Printf("parse string err:%v\n", err)
		return err, 0
	}
	// 按照东八区的时区格式解析一个字符串
	tlocal, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("Parse a string according to the time zone format of Dongba district err:%v\n", err)
		return err, 0
	}
	// 按照指定的时区解析时间
	t, err := time.ParseInLocation("2006-01-02 15:04:05", ts, tlocal)
	if err != nil {
		fmt.Printf("Resolve the time according to the specified time zone:%v\n", err)
		return err, 0
	}
	// 计算时间的差值
	reverseTime := now.Sub(t)
	return nil, reverseTime
}
