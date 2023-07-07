package helper

import "time"

func NumberToMonth(month int) string {
	date := time.Date(2000, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	return date.Month().String()
}
