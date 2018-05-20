package datetime

import (
	"fmt"
	"time"
)

const (
	Y_M                 = "2006-01"
	Y_M_D               = "2006-01-02"
	Y_M_D_2             = "2006年01月02日"
	Y_M_D_H_I_S         = "2006-01-02 15:04:05"
	Y_M_D_H_I_S_CST     = "2006-01-02 15:04:05 +0800 CST" //2016-12-04 15:39:06 +0800 CST
	Y_M_D_H_I_S_RFC3339 = "2006-01-02T15:04:05Z07:00"
	Y_M_D_H_I_S_2       = "2006年01月02日 15:04:05"
	H_I_S               = "15:04:05"
)

type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", t.DateTime())
	return []byte(stamp), nil
}

//日期时间
func (t DateTime) DateTime() string {
	return t.Format(Y_M_D_H_I_S)
}

//日期
func (t DateTime) Date() string {
	return t.Format(Y_M_D)
}

//时间
func (t DateTime) Time() string {
	return t.Format(H_I_S)
}
func (t DateTime) Format(layout string) string {
	return Format(t, layout)
}

//格式
func Format(str interface{}, layout string) string {
	date, err := FormatTimeStruct(str, layout)
	if err != nil {
		return ""
	}
	if date.IsZero() {
		return ""
	}

	return date.Format(layout)
}

//格式
func FormatTimeStruct(str interface{}, layout string) (time.Time, error) {
	var date time.Time
	var err error
	//判断变量类型
	switch str.(type) {
	case time.Time:
		date = str.(time.Time)
	case string:
		//如果是字符串则转换成 标准日期时间格式
		date, err = time.Parse(layout, str.(string))
		if err != nil {
			return time.Time{}, err
		}
	}

	return date, nil
}

//格式
func FormatTimeStructLocation(str interface{}, layout string) (time.Time, error) {
	var date time.Time
	var err error
	//判断变量类型
	switch str.(type) {
	case time.Time:
		date = str.(time.Time)
	case string:
		local, _ := time.LoadLocation("Local")
		//如果是字符串则转换成 标准日期时间格式
		date, err = time.ParseInLocation(layout, str.(string), local)
		if err != nil {
			return time.Time{}, err
		}
	}

	return date, nil
}

//当前日期时间
func Now() string {
	return time.Now().Format(Y_M_D_H_I_S)
}

//当前日期
func Date() string {
	return time.Now().Format(Y_M_D)
}

//当前时间
func Time() string {
	return time.Now().Format(H_I_S)
}

//当前年月
func YearMonth() string {
	return time.Now().Format(Y_M)
}

//时间格式化
func TimeFormatByYmdHms(year int, month time.Month, day, hour, min, sec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, 0, time.Local)
}

//当前年月日
func NowYearMonthDay() (year int, month time.Month, day int) {
	return time.Now().Date()
}

//当前时分秒
func NowHourMinSec() (hour, min, sec int) {
	return time.Now().Clock()
}

//输出当前日期是星期几
func Weekday() string {
	return time.Now().Weekday().String()
}

func ISOWeekStart(t time.Time) time.Time {
	wd := t.Weekday()
	if wd == time.Monday {
		return t
	}
	offset := int(time.Monday - wd)
	if offset > 0 {
		offset -= 7
	}
	return t.AddDate(0, 0, offset)
}

//指定日期是年中的第几天
func GetDaysInYear(t string) int {
	now, _ := FormatTimeStruct(t, Y_M_D)
	total := 0
	arr := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, month, d := now.Date()
	m := int(month)
	for i := 0; i < m-1; i++ {
		total = total + arr[i]
	}
	if (y%400 == 0 || (y%4 == 0 && y%100 != 0)) && m > 2 {
		total = total + d + 1

	} else {
		total = total + d
	}
	return total;
}

//年中的第几天
func GetDaysInYearByThisYear() int {
	return GetDaysInYear(Date())
}
