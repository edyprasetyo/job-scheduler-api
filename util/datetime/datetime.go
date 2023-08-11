package datetime

import (
	"strings"
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

func generalDateFormatToGoDateFormat(dateFormat string) string {
	goDateFormat := strings.ReplaceAll(dateFormat, "dd", "02")
	goDateFormat = strings.ReplaceAll(goDateFormat, "MM", "01")
	goDateFormat = strings.ReplaceAll(goDateFormat, "yyyy", "2006")
	goDateFormat = strings.ReplaceAll(goDateFormat, "HH", "15")
	goDateFormat = strings.ReplaceAll(goDateFormat, "mm", "04")
	goDateFormat = strings.ReplaceAll(goDateFormat, "ss", "05")
	return goDateFormat
}

func ToString(date time.Time, dateFormat string) string {
	goDateFormat := generalDateFormatToGoDateFormat(dateFormat)
	return date.Format(goDateFormat)
}

func FromString(dateStr string, dateFormat string) *time.Time {
	goDateFormat := generalDateFormatToGoDateFormat(dateFormat)
	date, err := time.Parse(goDateFormat, dateStr)
	if err != nil {
		return nil
	}
	return &date
}
