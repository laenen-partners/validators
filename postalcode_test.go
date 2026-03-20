package validators

import "testing"

func TestPostalCode(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		country string
		valid   bool
		code    string
	}{
		{"empty", "", "US", true, ""},
		{"valid US 5", "90210", "US", true, ""},
		{"valid US 5+4", "90210-1234", "US", true, ""},
		{"valid DE", "10115", "DE", true, ""},
		{"valid NL", "1234 AB", "NL", true, ""},
		{"valid NL no space", "1234AB", "NL", true, ""},
		{"valid BE", "1000", "BE", true, ""},
		{"valid GB", "SW1A 1AA", "GB", true, ""},
		{"valid GB no space", "SW1A1AA", "GB", true, ""},
		{"valid FR", "75001", "FR", true, ""},
		{"valid JP", "100-0001", "JP", true, ""},
		{"valid BR", "01001-000", "BR", true, ""},
		{"invalid US", "9021", "US", false, ErrPostalCodeFormatInvalid},
		{"invalid DE", "1011", "DE", false, ErrPostalCodeFormatInvalid},
		{"invalid BE", "100", "BE", false, ErrPostalCodeFormatInvalid},
		{"unsupported country", "12345", "ZZ", false, ErrPostalCodeCountryInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := PostalCode(tt.value, tt.country)
			if r.Valid != tt.valid {
				t.Errorf("PostalCode(%q, %q) valid=%v, want %v", tt.value, tt.country, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
