package dateutil

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/modern-go/reflect2"
	"github.com/uniplaces/carbon"
)

//DateTimePeriod 表示一个时间范围
type DateTimePeriod struct {
	StartAt *carbon.Carbon `json:"start_at"`
	EndAt   *carbon.Carbon `json:"end_at"`
}

//GetMonthsBetweenDates 返回时间区间内的所有月份
func GetMonthsBetweenDates(sDate, eDate uint) []string {
	var months []string
	timeFormatTpl := "200601"
	date, err := time.Parse(timeFormatTpl, strconv.Itoa(int(sDate)))
	if err != nil {
		return months
	}
	date2, err := time.Parse(timeFormatTpl, strconv.Itoa(int(eDate)))
	if err != nil {
		return months
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return months
	}

	for {
		dateStr := date.Format(timeFormatTpl)
		date2Str := date2.Format(timeFormatTpl)
		months = append(months, dateStr)
		if dateStr >= date2Str {
			break
		}
		date = date.AddDate(0, 1, 0)
	}
	return months
}

//FormatByFeYMD 格式化一个时间
func FormatByFeYMD(date, format string) string {
	parse, layout, _ := NewCarbonByFeYMD(date, format)
	return parse.Format(layout)
}

//FormatByPhpYMD 格式化一个时间
func FormatByPhpYMD(date, format string) string {
	parse, layout, _ := NewCarbonByPhpYMD(date, format)
	return parse.Format(layout)
}

//NewCarbonByFeYMD 生成一个Carbon时间
func NewCarbonByFeYMD(date, format string) (parse *carbon.Carbon, layout string, err error) {
	layout = FeYmd2layout(format)
	parse, err = carbon.Parse(layout, date, "Local")
	return parse, layout, err
}

//NewCarbonByPhpYMD 生成一个Carbon时间
func NewCarbonByPhpYMD(date, format string) (parse *carbon.Carbon, layout string, err error) {
	layout = PhpYmd2layout(format)
	parse, err = carbon.Parse(layout, date, "Local")
	return parse, layout, err
}

//TimeFormatByFeYMD 格式化一个时间
func TimeFormatByFeYMD(date time.Time, format string) string {
	return date.Format(FeYmd2layout(format))
}

//TimeFormatByPhpYMD 格式化一个时间
func TimeFormatByPhpYMD(date time.Time, format string) string {
	return date.Format(PhpYmd2layout(format))
}

//FeYmd2layout format转layout
func FeYmd2layout(format string) string {
	layout := strings.Replace(format, "YYYY", "2006", 1)
	layout = strings.Replace(layout, "YY", "06", 1)
	layout = strings.Replace(layout, "MM", "01", 1)
	layout = strings.Replace(layout, "M", "1", 1)
	layout = strings.Replace(layout, "DD", "02", 1)
	layout = strings.Replace(layout, "D", "2", 1)
	layout = strings.Replace(layout, "HH", "15", 1)
	layout = strings.Replace(layout, "H", "03", 1)
	layout = strings.Replace(layout, "ii", "04", 1)
	layout = strings.Replace(layout, "ss", "05", 1)
	return layout
}

//PhpYmd2layout format转layout
func PhpYmd2layout(format string) string {
	layout := strings.Replace(format, "Y", "2006", 1)
	layout = strings.Replace(layout, "y", "06", 1)
	layout = strings.Replace(layout, "M", "Jan", 1)
	layout = strings.Replace(layout, "m", "01", 1)
	layout = strings.Replace(layout, "F", "January", 1)
	layout = strings.Replace(layout, "n", "1", 1)
	layout = strings.Replace(layout, "l", "Monday", 1)
	layout = strings.Replace(layout, "D", "Mon", 1)
	layout = strings.Replace(layout, "d", "02", 1)
	layout = strings.Replace(layout, "j", "2", 1)
	layout = strings.Replace(layout, "H", "15", 1)
	layout = strings.Replace(layout, "h", "03", 1)
	layout = strings.Replace(layout, "i", "04", 1)
	layout = strings.Replace(layout, "s", "05", 1)
	layout = strings.Replace(layout, "P", "PM", 1)
	layout = strings.Replace(layout, "p", "pm", 1)
	return layout
}

//GetNextPeriodTime 获取下一个范围的初始时间
func GetNextPeriodTime(period string) time.Time {
	switch period {
	case "monthly":
		return GetNextMonth()
	case "weekly":
		return GetNextWeek()
	}
	return time.Now()
}

//GetNextMonth 获取下月一号时间
func GetNextMonth() time.Time {
	now := time.Now()
	monthTemplate := "200601"
	nextMonthStr := now.AddDate(0, 1, 0).Format(monthTemplate)
	nextMonthTime, _ := time.Parse(monthTemplate, nextMonthStr)
	return nextMonthTime
}

//GetNextWeek 获取下周周一时间
func GetNextWeek() time.Time {
	now := time.Now()
	gap := int(now.Weekday() - time.Monday)
	offset := 7 - gap
	nextMonday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return nextMonday
}

//GetCarbon 生成Carbon对象
func GetCarbon(v interface{}, format string) (*carbon.Carbon, error) {
	var c *carbon.Carbon
	if reflect2.IsNil(v) {
		c = carbon.Now()
	} else if t, ok := v.(*carbon.Carbon); ok {
		c = t
	} else if t, ok := v.(carbon.Carbon); ok {
		c = &t
	} else if t, ok := v.(time.Time); ok {
		c = carbon.NewCarbon(t)
	} else if t, ok := v.(string); ok {
		c, _, _ = NewCarbonByPhpYMD(t, format)
	} else {
		return nil, fmt.Errorf("无法识别参数")
	}
	return c, nil
}

//GetWeekPeriod 获取某个日期所处的Week范围
func GetWeekPeriod(v interface{}, format string) (*DateTimePeriod, error) {
	c, err := GetCarbon(v, format)
	if err != nil {
		return nil, err
	}
	return &DateTimePeriod{
		StartAt: c.StartOfWeek(),
		EndAt:   c.EndOfWeek(),
	}, nil
}

//GetMonthPeriod 获取某个日期所处的Month范围
func GetMonthPeriod(v interface{}, format string) (*DateTimePeriod, error) {
	c, err := GetCarbon(v, format)
	if err != nil {
		return nil, err
	}
	return &DateTimePeriod{
		StartAt: c.StartOfMonth(),
		EndAt:   c.EndOfMonth(),
	}, nil
}

//GetYearPeriod 获取某个日期所处的Year范围
func GetYearPeriod(v interface{}, format string) (*DateTimePeriod, error) {
	c, err := GetCarbon(v, format)
	if err != nil {
		return nil, err
	}
	return &DateTimePeriod{
		StartAt: c.StartOfYear(),
		EndAt:   c.EndOfYear(),
	}, nil
}

//GetQuarterPeriod 获取某个日期所处的Quarter范围
func GetQuarterPeriod(v interface{}, format string) (*DateTimePeriod, error) {
	c, err := GetCarbon(v, format)
	if err != nil {
		return nil, err
	}
	return &DateTimePeriod{
		StartAt: c.StartOfQuarter(),
		EndAt:   c.EndOfQuarter(),
	}, nil
}

//GetDayPeriod 获取某个日期所处的Day范围
func GetDayPeriod(v interface{}, format string) (*DateTimePeriod, error) {
	c, err := GetCarbon(v, format)
	if err != nil {
		return nil, err
	}
	return &DateTimePeriod{
		StartAt: c.StartOfDay(),
		EndAt:   c.EndOfDay(),
	}, nil
}
