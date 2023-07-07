package helper

import (
	"errors"
	"time"
)

func GenerateDateRange(startDate, endDate string) ([]string, error) {
	layout := "2006-01-02"
	start, err := time.Parse(layout, startDate)
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(layout, endDate)
	if err != nil {
		return nil, err
	}

	var dates []string
	for current := start; current.Before(end) || current.Equal(end); current = current.AddDate(0, 0, 1) {
		dates = append(dates, current.Format(layout))
	}

	if len(dates) == 0 {
		return nil, errors.New("invalid date range format")
	}

	return dates, nil
}
