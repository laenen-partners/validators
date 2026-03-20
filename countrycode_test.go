package validators

import "testing"

func TestCountryCode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid alpha-2", "BE", true, ""},
		{"valid alpha-2 lowercase", "be", true, ""},
		{"valid alpha-3", "BEL", true, ""},
		{"valid alpha-3 lowercase", "bel", true, ""},
		{"valid US", "US", true, ""},
		{"valid USA", "USA", true, ""},
		{"unknown alpha-2", "XX", false, ErrCountryCodeUnknown},
		{"unknown alpha-3", "XXX", false, ErrCountryCodeUnknown},
		{"too short", "B", false, ErrCountryCodeFormatInvalid},
		{"too long", "BELG", false, ErrCountryCodeFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CountryCode(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("CountryCode(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
