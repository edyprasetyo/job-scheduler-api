package tools

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func generalDateFormatToGoDateFormat(dateFormat string) string {
	goDateFormat := strings.ReplaceAll(dateFormat, "dd", "02")
	goDateFormat = strings.ReplaceAll(goDateFormat, "MM", "01")
	goDateFormat = strings.ReplaceAll(goDateFormat, "yyyy", "2006")
	goDateFormat = strings.ReplaceAll(goDateFormat, "HH", "15")
	goDateFormat = strings.ReplaceAll(goDateFormat, "mm", "04")
	goDateFormat = strings.ReplaceAll(goDateFormat, "ss", "05")
	return goDateFormat
}

func DateToString(date time.Time, dateFormat string) string {
	goDateFormat := generalDateFormatToGoDateFormat(dateFormat)
	return date.Format(goDateFormat)
}

func GenerateUUID() string {
	return uuid.New().String()
}

func StringToDate(dateStr string, dateFormat string) *time.Time {
	goDateFormat := generalDateFormatToGoDateFormat(dateFormat)
	date, err := time.Parse(goDateFormat, dateStr)
	if err != nil {
		return nil
	}
	return &date
}

func IsDateBefore(date1 time.Time, date2 time.Time) bool {
	t1 := time.Date(date1.Year(), date1.Month(), date1.Day(), date1.Hour(), date1.Minute(), date1.Second(), date1.Nanosecond(), time.UTC)
	t2 := time.Date(date2.Year(), date2.Month(), date2.Day(), date2.Hour(), date2.Minute(), date2.Second(), date2.Nanosecond(), time.UTC)
	return t1.Before(t2)
}

func DatediffInSeconds(date1 time.Time, date2 time.Time) int {
	t1 := time.Date(date1.Year(), date1.Month(), date1.Day(), date1.Hour(), date1.Minute(), date1.Second(), date1.Nanosecond(), time.UTC)
	t2 := time.Date(date2.Year(), date2.Month(), date2.Day(), date2.Hour(), date2.Minute(), date2.Second(), date2.Nanosecond(), time.UTC)
	return int(t1.Sub(t2).Seconds())
}
