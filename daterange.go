package validators

import (
	"fmt"
	"time"
)

// DateRange validates that a date string (YYYY-MM-DD) falls between min and max (inclusive).
// Either min or max may be empty to leave that bound open.
func DateRange(value, minDate, maxDate string) Result {
	if value == "" {
		return valid(nil)
	}

	r := Date(value)
	if !r.Valid {
		return r
	}

	parsed, _ := time.Parse("2006-01-02", value)

	if minDate != "" {
		minR := Date(minDate)
		if !minR.Valid {
			return invalid(ErrDateFormatInvalid, "Invalid min date format", "date", map[string]any{
				"min_date": minDate,
			})
		}
		minParsed, _ := time.Parse("2006-01-02", minDate)
		if parsed.Before(minParsed) {
			return invalid(ErrDateRangeBeforeMin,
				fmt.Sprintf("Date must not be before %s", minDate),
				"date", map[string]any{
					"value":    value,
					"min_date": minDate,
				})
		}
	}

	if maxDate != "" {
		maxR := Date(maxDate)
		if !maxR.Valid {
			return invalid(ErrDateFormatInvalid, "Invalid max date format", "date", map[string]any{
				"max_date": maxDate,
			})
		}
		maxParsed, _ := time.Parse("2006-01-02", maxDate)
		if parsed.After(maxParsed) {
			return invalid(ErrDateRangeAfterMax,
				fmt.Sprintf("Date must not be after %s", maxDate),
				"date", map[string]any{
					"value":    value,
					"max_date": maxDate,
				})
		}
	}

	return valid(r.Metadata)
}
