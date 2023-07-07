package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func GetPresenceDailyReportName(date string) string {
	return fmt.Sprintf("ATTENDANCE_DAILY_REPORT_%s.xlsx", date)
}

func GetPresenceDailyRangeReportName(startDate, endDate string) string {
	return fmt.Sprintf("ATTENDANCE_DAILY_REPORT_%s-%s.xlsx", startDate, endDate)
}

func GetPresenceMonthlyReportName(month, year string) string {
	monthNum, _ := strconv.Atoi(month)
	month = strings.ToUpper(NumberToMonth(monthNum))
	return fmt.Sprintf("ATTENDANCE_MONTHLY_REPORT_%s-%s.xlsx", month, year)
}

func GetPresenceMonthlyNIPReportName(name, month, year string) string {
	monthNum, _ := strconv.Atoi(month)
	month = strings.ToUpper(NumberToMonth(monthNum))
	removeSpace := strings.ReplaceAll(name, " ", "_")
	upper := strings.ToUpper(removeSpace)
	return fmt.Sprintf("ATTENDANCE_MONTHLY_REPORT_%s_%s-%s.xlsx", upper, month, year)
}

func GetPresenceYearlyReportName(year string) string {
	return fmt.Sprintf("ATTENDANCE_YEARLY_REPORT_%s.xlsx", year)
}

func GetPresenceYearlyNIPReportName(name, year string) string {
	removeSpace := strings.ReplaceAll(name, " ", "_")
	upper := strings.ToUpper(removeSpace)
	return fmt.Sprintf("ATTENDANCE_YEARLY_REPORT_%s_%s.xlsx", upper, year)
}
