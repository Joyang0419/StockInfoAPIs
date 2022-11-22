package helpers

import (
	"time"
)

func FitDateFormat(dateFormat string, dateStrs ...string) bool {
	for _, dateStr := range dateStrs {
		_, err := time.Parse(dateFormat, dateStr)
		if err != nil {
			return false
		}
	}
	return true
}
