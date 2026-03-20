package validators

import (
	"fmt"
	"time"
)

// DateInPast validates that a date string (YYYY-MM-DD) is in the past.
// If maxAge is > 0, the date must not be further back than maxAge from today.
// A maxAge of 0 means any past date is accepted.
func DateInPast(value string, maxAge time.Duration) Result {
	if value == "" {
		return valid(nil)
	}

	// First validate format using the existing Date validator.
	r := Date(value)
	if !r.Valid {
		return r
	}

	parsed, _ := time.Parse("2006-01-02", value)
	now := time.Now().Truncate(24 * time.Hour)

	if !parsed.Before(now) {
		return invalid(ErrDateInPastNotPast, "Date must be in the past", "date", map[string]any{
			"value": value,
			"today": now.Format("2006-01-02"),
		})
	}

	if maxAge > 0 {
		earliest := now.Add(-maxAge)
		if parsed.Before(earliest) {
			return invalid(ErrDateInPastTooFar,
				fmt.Sprintf("Date must not be more than %s in the past", formatDuration(maxAge)),
				"date", map[string]any{
					"value":    value,
					"today":    now.Format("2006-01-02"),
					"earliest": earliest.Format("2006-01-02"),
				})
		}
	}

	return valid(r.Metadata)
}

func formatDuration(d time.Duration) string {
	days := int(d.Hours() / 24)
	if days >= 365 {
		years := days / 365
		if years == 1 {
			return "1 year"
		}
		return fmt.Sprintf("%d years", years)
	}
	if days >= 30 {
		months := days / 30
		if months == 1 {
			return "1 month"
		}
		return fmt.Sprintf("%d months", months)
	}
	if days == 1 {
		return "1 day"
	}
	return fmt.Sprintf("%d days", days)
}
