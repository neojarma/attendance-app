package helper

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func ParseTime(timeStr string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		log.Println("Error parsing time:", err)
		return "", err
	}

	hour := checkDigit(strconv.Itoa(parsedTime.Hour()))
	minute := checkDigit(strconv.Itoa(parsedTime.Minute()))
	second := checkDigit(strconv.Itoa(parsedTime.Second()))

	res := fmt.Sprintf("%v:%v:%v", hour, minute, second)
	return res, nil
}

func checkDigit(timeStr string) string {
	if len(timeStr) == 1 {
		return fmt.Sprintf("0%s", timeStr)
	}

	return timeStr
}

func IsTimeValid(timeStr string) bool {
	_, err := time.Parse("15:04:05", timeStr)
	return err == nil
}
