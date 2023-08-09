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
