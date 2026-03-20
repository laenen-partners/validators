package validators

import (
	"testing"
	"time"
)

func TestDateInFuture(t *testing.T) {
	today := time.Now().Truncate(24 * time.Hour)
	yesterday := today.AddDate(0, 0, -1).Format("2006-01-02")
	tomorrow := today.AddDate(0, 0, 1).Format("2006-01-02")
	todayStr := today.Format("2006-01-02")
	farFuture := today.AddDate(5, 0, 1).Format("2006-01-02")
	nextWeek := today.AddDate(0, 0, 7).Format("2006-01-02")

	tests := []struct {
		name     string
		value    string
		maxAhead time.Duration
		valid    bool
		code     string
	}{
		{"empty", "", 0, true, ""},
		{"tomorrow no limit", tomorrow, 0, true, ""},
		{"far future no limit", farFuture, 0, true, ""},
		{"today is not future", todayStr, 0, false, ErrDateInFutureNotFuture},
		{"yesterday is not future", yesterday, 0, false, ErrDateInFutureNotFuture},
		{"next week within 30 days", nextWeek, 30 * 24 * time.Hour, true, ""},
		{"far future with 1 year limit", farFuture, 365 * 24 * time.Hour, false, ErrDateInFutureTooFar},
		{"bad format", "nope", 0, false, ErrDateFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := DateInFuture(tt.value, tt.maxAhead)
			if r.Valid != tt.valid {
				t.Errorf("DateInFuture(%q, %v) valid=%v, want %v", tt.value, tt.maxAhead, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
