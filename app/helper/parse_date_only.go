package helper

import (
	"strings"
	"time"
)

func ParseDateOnly(date string) string {
	return strings.Split(date, "T")[0]
}

func IsDateValid(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
