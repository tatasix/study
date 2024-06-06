package mytest

import (
	"fmt"
	"time"
)

const (
	_ = iota
	RightsTimesPeriodDay
	RightsTimesPeriodZero
	RightsTimesPeriodMonth
	RightsTimesPeriodYear
	RightsTimesPeriodForever
)

func generateTimeRange(now time.Time, period int) (start time.Time, end time.Time) {
	switch period {
	case RightsTimesPeriodDay:
		start = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		end = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())
	case RightsTimesPeriodZero:
		start = time.Time{}
		end = time.Time{}
	case RightsTimesPeriodMonth:
		start = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		nextMonth := now.AddDate(0, 1, 0)
		end = time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, -1, now.Location())
	case RightsTimesPeriodYear:
		start = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		nextYear := now.AddDate(1, 0, 0)
		end = time.Date(nextYear.Year(), 1, 1, 0, 0, 0, -1, now.Location())
	case RightsTimesPeriodForever:
		start = now
		end = time.Time{}
	}
	return start, end
}

func Handle2() {
	now := time.Now()
	period := RightsTimesPeriodMonth
	start, end := generateTimeRange(now, period)
	fmt.Printf("start=%s, end=%s\n", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
}
