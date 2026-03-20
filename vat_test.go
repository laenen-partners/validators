package validators

import "testing"

func TestVAT(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid DE", "DE123456789", true, ""},
		{"valid BE", "BE0123456789", true, ""},
		{"valid NL", "NL123456789B01", true, ""},
		{"valid FR", "FRAA123456789", true, ""},
		{"valid AT", "ATU12345678", true, ""},
		{"valid ES", "ESA1234567B", true, ""},
		{"valid IT", "IT12345678901", true, ""},
		{"valid PL", "PL1234567890", true, ""},
		{"valid CH", "CHE123456789MWST", true, ""},
		{"valid CH TVA", "CHE123456789TVA", true, ""},
		{"too short", "DE1", false, ErrVATTooShort},
		{"unknown country", "XX123456789", false, ErrVATCountryInvalid},
		{"bad format DE", "DE12345678", false, ErrVATFormatInvalid},
		{"bad format BE", "BE9123456789", false, ErrVATFormatInvalid},
		{"numeric prefix", "12345678901", false, ErrVATCountryInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := VAT(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("VAT(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
