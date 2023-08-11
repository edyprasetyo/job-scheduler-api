package datetime

import (
	"time"
)

func Now() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.UTC)
}

func Before(date1 time.Time, date2 time.Time) bool {
	return date1.Before(date2)
}

func After(date1 time.Time, date2 time.Time) bool {
	return date1.After(date2)
}

func DiffInSeconds(date1 time.Time, date2 time.Time) int {
	return int(date2.Sub(date1).Seconds())
}

func DiffInMinutes(date1 time.Time, date2 time.Time) int {
	return int(date2.Sub(date1).Minutes())
}

func DiffInHours(date1 time.Time, date2 time.Time) int {
	return int(date2.Sub(date1).Hours())
}

func DiffInDays(date1 time.Time, date2 time.Time) int {
	return int(date2.Sub(date1).Hours() / 24)
}
