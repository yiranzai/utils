package dateutil

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/modern-go/reflect2"
	"github.com/uniplaces/carbon"
)

func Test_Period(t *testing.T) {
	var c *carbon.Carbon
	var period *DateTimePeriod
	var err error
	assert.Equal(t, reflect2.IsNil(err), true)
	str := "2020-02-28 15:14:33"

	c, err = GetCarbon(str, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, c.DateTimeString(), str)

	period, err = GetWeekPeriod(str, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-02-24 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-03-01 23:59:59")

	period, err = GetMonthPeriod(str, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-02-01 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-02-29 23:59:59")

	period, err = GetYearPeriod(str, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-01-01 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-12-31 23:59:59")

	period, err = GetQuarterPeriod(str, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-01-01 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-03-31 23:59:59")

	period, err = GetDayPeriod(str, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-02-28 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-02-28 23:59:59")

	c, _ = carbon.Parse("2006-01-02 15:04:05", str, "Local")
	c, err = GetCarbon(c, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, c.DateTimeString(), str)

	period, err = GetWeekPeriod(c, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-02-24 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-03-01 23:59:59")

	period, err = GetMonthPeriod(c, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-02-01 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-02-29 23:59:59")

	period, err = GetYearPeriod(c, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-01-01 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-12-31 23:59:59")

	period, err = GetQuarterPeriod(c, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-01-01 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-03-31 23:59:59")

	period, err = GetDayPeriod(c, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-02-28 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-02-28 23:59:59")

	b, _ := carbon.Parse("2006-01-02 15:04:05", str, "Local")
	c, err = GetCarbon(*b, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, c.DateTimeString(), str)

	period, err = GetWeekPeriod(*b, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-02-24 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-03-01 23:59:59")

	period, err = GetMonthPeriod(*b, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-02-01 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-02-29 23:59:59")

	period, err = GetYearPeriod(*b, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-01-01 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-12-31 23:59:59")

	period, err = GetQuarterPeriod(*b, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-01-01 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-03-31 23:59:59")

	period, err = GetDayPeriod(*b, "Y-m-d H:i:s")
	assert.Equal(t, reflect2.IsNil(err), true)
	assert.Equal(t, period.StartAt.DateTimeString(), "2020-02-28 00:00:00")
	assert.Equal(t, period.EndAt.DateTimeString(), "2020-02-28 23:59:59")
}
