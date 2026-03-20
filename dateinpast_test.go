package validators

import (
	"testing"
	"time"
)

func TestDateInPast(t *testing.T) {
	today := time.Now().Truncate(24 * time.Hour)
	yesterday := today.AddDate(0, 0, -1).Format("2006-01-02")
	tomorrow := today.AddDate(0, 0, 1).Format("2006-01-02")
	todayStr := today.Format("2006-01-02")
	longAgo := "1900-01-01"
	oneYearAgo := today.AddDate(-1, 0, -1).Format("2006-01-02")

	tests := []struct {
		name   string
		value  string
		maxAge time.Duration
		valid  bool
		code   string
	}{
		{"empty", "", 0, true, ""},
		{"yesterday no limit", yesterday, 0, true, ""},
		{"long ago no limit", longAgo, 0, true, ""},
		{"today is not past", todayStr, 0, false, ErrDateInPastNotPast},
		{"tomorrow is not past", tomorrow, 0, false, ErrDateInPastNotPast},
		{"yesterday within 30 days", yesterday, 30 * 24 * time.Hour, true, ""},
		{"over a year ago with 1 year limit", oneYearAgo, 365 * 24 * time.Hour, false, ErrDateInPastTooFar},
		{"bad format", "not-a-date", 0, false, ErrDateFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := DateInPast(tt.value, tt.maxAge)
			if r.Valid != tt.valid {
				t.Errorf("DateInPast(%q, %v) valid=%v, want %v", tt.value, tt.maxAge, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
