package validators

import "testing"

func TestCurrencyCode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid EUR", "EUR", true, ""},
		{"valid USD", "USD", true, ""},
		{"valid lowercase", "eur", true, ""},
		{"valid GBP", "GBP", true, ""},
		{"valid JPY", "JPY", true, ""},
		{"unknown", "XXX", false, ErrCurrencyCodeUnknown},
		{"too short", "EU", false, ErrCurrencyCodeFormatInvalid},
		{"too long", "EURO", false, ErrCurrencyCodeFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CurrencyCode(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("CurrencyCode(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
