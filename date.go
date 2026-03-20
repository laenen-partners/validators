package validators

import (
	"fmt"
	"strconv"
	"strings"
)

// Date validates an ISO 8601 date string (YYYY-MM-DD).
func Date(value string) Result {
	if value == "" {
		return valid(nil)
	}

	parts := strings.Split(value, "-")
	if len(parts) != 3 || len(parts[0]) != 4 || len(parts[1]) != 2 || len(parts[2]) != 2 {
		return invalid(ErrDateFormatInvalid, "Date must be in YYYY-MM-DD format", "date", map[string]any{
			"value": value,
		})
	}

	year, errY := strconv.Atoi(parts[0])
	month, errM := strconv.Atoi(parts[1])
	day, errD := strconv.Atoi(parts[2])
	if errY != nil || errM != nil || errD != nil {
		return invalid(ErrDateFormatInvalid, "Date contains non-numeric components", "date", map[string]any{
			"value": value,
		})
	}

	if month < 1 || month > 12 {
		return invalid(ErrDateInvalid,
			fmt.Sprintf("Month %d is out of range (1-12)", month),
			"date", map[string]any{
				"value": value,
				"month": month,
			})
	}

	maxDay := daysInMonth(year, month)
	if day < 1 || day > maxDay {
		return invalid(ErrDateInvalid,
			fmt.Sprintf("Day %d is out of range for %04d-%02d (max %d)", day, year, month, maxDay),
			"date", map[string]any{
				"value":   value,
				"year":    year,
				"month":   month,
				"day":     day,
				"max_day": maxDay,
			})
	}

	return valid(map[string]any{
		"year":  year,
		"month": month,
		"day":   day,
	})
}

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func daysInMonth(year, month int) int {
	switch month {
	case 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}
