package validators

import "testing"

func TestDateRange(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		minDate string
		maxDate string
		valid   bool
		code    string
	}{
		{"empty", "", "2020-01-01", "2025-12-31", true, ""},
		{"within range", "2023-06-15", "2020-01-01", "2025-12-31", true, ""},
		{"at min boundary", "2020-01-01", "2020-01-01", "2025-12-31", true, ""},
		{"at max boundary", "2025-12-31", "2020-01-01", "2025-12-31", true, ""},
		{"before min", "2019-12-31", "2020-01-01", "2025-12-31", false, ErrDateRangeBeforeMin},
		{"after max", "2026-01-01", "2020-01-01", "2025-12-31", false, ErrDateRangeAfterMax},
		{"open min", "1900-01-01", "", "2025-12-31", true, ""},
		{"open max", "2099-01-01", "2020-01-01", "", true, ""},
		{"both open", "2023-06-15", "", "", true, ""},
		{"bad format", "nope", "2020-01-01", "2025-12-31", false, ErrDateFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := DateRange(tt.value, tt.minDate, tt.maxDate)
			if r.Valid != tt.valid {
				t.Errorf("DateRange(%q, %q, %q) valid=%v, want %v", tt.value, tt.minDate, tt.maxDate, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
