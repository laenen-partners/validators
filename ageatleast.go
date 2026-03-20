package validators

import (
	"fmt"
	"time"
)

// AgeAtLeast validates that a birth date (YYYY-MM-DD) represents an age of at least minYears.
// Uses calendar year calculation (not duration), handling leap years correctly.
func AgeAtLeast(birthDate string, minYears int) Result {
	if birthDate == "" {
		return valid(nil)
	}

	r := Date(birthDate)
	if !r.Valid {
		return r
	}

	parsed, _ := time.Parse("2006-01-02", birthDate)
	now := time.Now()

	age := now.Year() - parsed.Year()
	// Subtract 1 if birthday hasn't occurred yet this year.
	if now.Month() < parsed.Month() || (now.Month() == parsed.Month() && now.Day() < parsed.Day()) {
		age--
	}

	if age < minYears {
		return invalid(ErrAgeAtLeastTooYoung,
			fmt.Sprintf("Must be at least %d years old, got %d", minYears, age),
			"date", map[string]any{
				"value":    birthDate,
				"age":      age,
				"min_age":  minYears,
			})
	}

	return valid(map[string]any{
		"age":   age,
		"year":  r.Metadata["year"],
		"month": r.Metadata["month"],
		"day":   r.Metadata["day"],
	})
}
