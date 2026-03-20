package validators

import (
	"fmt"
	"time"
)

// DateInFuture validates that a date string (YYYY-MM-DD) is in the future.
// If maxAhead is > 0, the date must not be further ahead than maxAhead from today.
// A maxAhead of 0 means any future date is accepted.
func DateInFuture(value string, maxAhead time.Duration) Result {
	if value == "" {
		return valid(nil)
	}

	r := Date(value)
	if !r.Valid {
		return r
	}

	parsed, _ := time.Parse("2006-01-02", value)
	now := time.Now().Truncate(24 * time.Hour)

	if !parsed.After(now) {
		return invalid(ErrDateInFutureNotFuture, "Date must be in the future", "date", map[string]any{
			"value": value,
			"today": now.Format("2006-01-02"),
		})
	}

	if maxAhead > 0 {
		latest := now.Add(maxAhead)
		if parsed.After(latest) {
			return invalid(ErrDateInFutureTooFar,
				fmt.Sprintf("Date must not be more than %s in the future", formatDuration(maxAhead)),
				"date", map[string]any{
					"value":  value,
					"today":  now.Format("2006-01-02"),
					"latest": latest.Format("2006-01-02"),
				})
		}
	}

	return valid(r.Metadata)
}
