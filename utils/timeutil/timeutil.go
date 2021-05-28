package timeutil

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	FormatDateTime = "2006-01-02 15:04:05"
	FormatDate     = "2006-01-02"
	FormatTime     = "15:04:05"
)

var (
	ErrParseTime = errors.New("解析时间出错")
)

// IsSameDateStr 判断是否为同一天
func IsSameDateStr(dateStr, format string, flagDate time.Time) (res bool, err error) {
	date, err := time.Parse(format, dateStr)
	if err != nil {
		return
	}
	f := "20060102"
	res = date.Format(f) == flagDate.Format(f)
	return
}
func IsSameDate(date1, date2 time.Time) bool {
	f := "20060102"
	return date1.Format(f) == date2.Format(f)
}

const (
	SecondsOfHour  = time.Second * 60 * 60
	SecondsOfDay   = SecondsOfHour * 24
	SecondsOfWeek  = SecondsOfDay * 7
	SecondsOfMonth = SecondsOfDay * 30

	ReqDateToday     = "today"
	ReqDateYesterday = "yesterday"
	ReDateSeven      = "sevendays"
	ReqDateThisMonth = "thisMonth"
	ReqDateLastMonth = "lastMonth"
)

// Datetime 时间格式化
type Datetime time.Time

func (t Datetime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// Date 日期格式化
type Date time.Time

func (t Date) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02"))
	return []byte(stamp), nil
}

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("15:04"))
	return []byte(stamp), nil
}

// BeginDay 获取某天的开始结束时间
func BeginDay(date time.Time) time.Time {
	timeStr := date.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	return t
}

// EndDay 获取某天的最后结束时间
func EndDay(date time.Time) time.Time {
	timeStr := date.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	return t
}

// BeginDateOfMonth 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func BeginDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return ZeroTime(d)
}

// EndDateOfMonth 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func EndDateOfMonth(d time.Time) time.Time {
	return BeginDateOfMonth(d).AddDate(0, 1, -1)
}

// ZeroTime 获取某一天的0点时间
func ZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 毫秒转时间
func MsToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	tm := time.Unix(0, msInt*int64(time.Millisecond))

	fmt.Println(tm.Format("2006-02-01 15:04:05.000"))

	return tm, nil
}

// IsNowInTimeRange 当前时间是否在指定范围内
//参数为时间字符串，格式为"时:分:秒"
func IsNowInTimeRange(startTimeStr, endTimeStr string) bool {
	//当前时间
	now := time.Now()
	//统一日期
	format := now.Format("2006-01-02")
	//转换为time类型需要的格式
	layout := "2006-01-02 15:04:05"
	//将开始时间拼接“年-月-日 ”转换为time类型
	timeStart, _ := time.ParseInLocation(layout, format+" "+startTimeStr, time.Local)
	//将结束时间拼接“年-月-日 ”转换为time类型
	timeEnd, _ := time.ParseInLocation(layout, format+" "+endTimeStr, time.Local)
	//使用time的Before和After方法，判断当前时间是否在参数的时间范围
	return now.Before(timeEnd) && now.After(timeStart)
}

// IsToday 时间格式为：202101040001
func IsToday(timeStr string) (ok bool) {
	s := timeStr[0 : len(timeStr)-4]
	if parse, err := time.ParseInLocation("20060102", s, time.Local); err != nil {
		panic(err)
	} else {
		today := BeginDay(time.Now())
		return parse == today
	}
}

// GetBeginAndEndDate 根据类型返回开始时间和结束时间
func GetBeginAndEndDate(dateType string) (begin time.Time, end time.Time, err error) {
	now := time.Now()
	switch dateType {
	case ReqDateToday, "": //默认今天
		begin = BeginDay(now)
		end = EndDay(now)
	case ReqDateYesterday:
		now = now.AddDate(0, 0, -1)
		begin = BeginDay(now)
		end = EndDay(now)
	case ReDateSeven: //7天
		now = now.AddDate(0, 0, -7)
		begin = BeginDay(now)
		end = EndDay(time.Now())
	case ReqDateThisMonth:
		begin = BeginDateOfMonth(now)
		end = EndDateOfMonth(now)
	case ReqDateLastMonth:
		now = now.AddDate(0, -1, 0)
		begin = BeginDateOfMonth(now)
		end = EndDateOfMonth(now)
	default:
		err = errors.New(fmt.Sprintf("不支持的类型:%v", dateType))
	}
	return
}
