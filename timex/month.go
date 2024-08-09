package timex

import (
	"time"
)

func Month(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func CurrentMonth() time.Time {
	return Month(time.Now())
}

// ThisMonth 返回指定时间的月份的第一天
func ThisMonth(ts ...time.Time) time.Time {
	if len(ts) <= 0 {
		return CurrentMonth()
	}
	return Month(ts[0])
}

// LastMonth 返回指定时间的上一个月的第一天
func LastMonth(ts ...time.Time) time.Time {
	return ThisMonth(ts...).AddDate(0, -1, 0)
}

// NextMonth 返回指定时间的下个月的第一天
func NextMonth(ts ...time.Time) time.Time {
	return ThisMonth(ts...).AddDate(0, 1, 0)
}

// MonthOfLastDay 返回指定时间的月份的最后一天
func MonthOfLastDay(ts ...time.Time) time.Time {
	return NextMonth(ts...).AddDate(0, 0, -1)
}
