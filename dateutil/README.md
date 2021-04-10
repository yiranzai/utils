### Dateutil

时间和日期处理相关

### Features

##### `GetMonthsBetweenDates(sDate, eDate uint) []string` // 返回时间区间内的所有月份

`GetMonthsBetweenDates(sDate, eDate uint) []string` // 返回时间区间内的所有月份

##### `FormatByFeYMD(date, format string) string` // 格式化一个时间

`FormatByFeYMD(date, format string) string` // 格式化一个时间

##### `FormatByPhpYMD(date, format string) string` // 格式化一个时间

`FormatByPhpYMD(date, format string) string` // 格式化一个时间

##### `NewCarbonByFeYMD(date, format string) (parse *carbon.Carbon, layout string, err error)` // 生成一个 Carbon 时间

`NewCarbonByFeYMD(date, format string) (parse *carbon.Carbon, layout string, err error)` // 生成一个 Carbon 时间

##### `NewCarbonByPhpYMD(date, format string) (parse *carbon.Carbon, layout string, err error)` // 生成一个 Carbon 时间

`NewCarbonByPhpYMD(date, format string) (parse *carbon.Carbon, layout string, err error)` // 生成一个 Carbon 时间

##### `TimeFormatByFeYMD(date time.Time, format string) string` // 格式化一个时间

`TimeFormatByFeYMD(date time.Time, format string) string` // 格式化一个时间

##### `TimeFormatByPhpYMD(date time.Time, format string) string` // 格式化一个时间

`TimeFormatByPhpYMD(date time.Time, format string) string` // 格式化一个时间

##### `FeYmd2layout(format string) string` // format 转 layout

`FeYmd2layout(format string) string` // format 转 layout

##### `PhpYmd2layout(format string) string` // format 转 layout

`PhpYmd2layout(format string) string` // format 转 layout

##### `GetNextPeriodTime(period string) time.Time` // 获取下一个范围的初始时间

`GetNextPeriodTime(period string) time.Time` // 获取下一个范围的初始时间

##### `GetNextMonth() time.Time` // 获取下月一号时间

`GetNextMonth() time.Time` // 获取下月一号时间

##### `GetNextWeek() time.Time` // 获取下周周一时间

`GetNextWeek() time.Time` // 获取下周周一时间

##### `GetCarbon(v interface{}, format string) (*carbon.Carbon, error)` // 生成 Carbon 对象

`GetCarbon(v interface{}, format string) (*carbon.Carbon, error)` // 生成 Carbon 对象

##### `GetWeekPeriod(v interface{}, format string) (*DateTimePeriod, error)` // 获取某个日期所处的 Week 范围

`GetWeekPeriod(v interface{}, format string) (*DateTimePeriod, error)` // 获取某个日期所处的 Week 范围

##### `GetMonthPeriod(v interface{}, format string) (*DateTimePeriod, error)` // 获取某个日期所处的 Month 范围

`GetMonthPeriod(v interface{}, format string) (*DateTimePeriod, error)` // 获取某个日期所处的 Month 范围

##### `GetYearPeriod(v interface{}, format string) (*DateTimePeriod, error)` // 获取某个日期所处的 Year 范围

`GetYearPeriod(v interface{}, format string) (*DateTimePeriod, error)` // 获取某个日期所处的 Year 范围

##### `GetQuarterPeriod(v interface{}, format string) (*DateTimePeriod, error)` // 获取某个日期所处的 Quarter 范围

`GetQuarterPeriod(v interface{}, format string) (*DateTimePeriod, error)` // 获取某个日期所处的 Quarter 范围

##### `GetDayPeriod(v interface{}, format string) (*DateTimePeriod, error)` // 获取某个日期所处的 Day 范围

`GetDayPeriod(v interface{}, format string) (*DateTimePeriod, error)` // 获取某个日期所处的 Day 范围
