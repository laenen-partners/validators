package validators

import (
	"math"
	"testing"
)

func TestNumberInRange(t *testing.T) {
	tests := []struct {
		name  string
		value string
		min   string
		max   string
		valid bool
		code  string
	}{
		{"empty", "", "0", "100", true, ""},
		{"within range", "50", "0", "100", true, ""},
		{"at min", "0", "0", "100", true, ""},
		{"at max", "100", "0", "100", true, ""},
		{"below min", "-1", "0", "100", false, ErrNumberBelowMin},
		{"above max", "101", "0", "100", false, ErrNumberAboveMax},
		{"decimal within", "19.99", "0.01", "9999.99", true, ""},
		{"decimal precision", "0.1", "0.1", "0.1", true, ""},
		{"open min", "-999999", "", "100", true, ""},
		{"open max", "999999", "0", "", true, ""},
		{"both open", "42", "", "", true, ""},
		{"invalid number", "abc", "0", "100", false, ErrNumberFormatInvalid},
		{"negative decimals", "-0.50", "-1.00", "1.00", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NumberInRange(tt.value, tt.min, tt.max)
			if r.Valid != tt.valid {
				t.Errorf("NumberInRange(%q, %q, %q) valid=%v, want %v", tt.value, tt.min, tt.max, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}

func TestNumberInRange_ExactDecimal(t *testing.T) {
	// 0.1 + 0.2 == 0.3 must work with string-based exact arithmetic.
	r := NumberInRange("0.3", "0.3", "0.3")
	if !r.Valid {
		t.Error("NumberInRange should handle 0.3 exactly")
	}
}

func TestNumberInRangeFloat(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		min   float64
		max   float64
		valid bool
		code  string
	}{
		{"within range", 50, 0, 100, true, ""},
		{"below min", -1, 0, 100, false, ErrNumberBelowMin},
		{"above max", 101, 0, 100, false, ErrNumberAboveMax},
		{"NaN", math.NaN(), 0, 100, false, ErrNumberFormatInvalid},
		{"Inf", math.Inf(1), 0, 100, false, ErrNumberFormatInvalid},
		{"open min", -999, math.NaN(), 100, true, ""},
		{"open max", 999, 0, math.NaN(), true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NumberInRangeFloat(tt.value, tt.min, tt.max)
			if r.Valid != tt.valid {
				t.Errorf("NumberInRangeFloat(%g, %g, %g) valid=%v, want %v", tt.value, tt.min, tt.max, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
