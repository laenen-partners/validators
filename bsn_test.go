package validators

import "testing"

func TestDutchBSN(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid", "111222333", true, ""},
		{"valid with dots", "1112.22.333", true, ""},
		{"too short", "12345678", false, ErrBSNLengthInvalid},
		{"too long", "1234567890", false, ErrBSNLengthInvalid},
		{"non-digits", "12345678a", false, ErrBSNFormatInvalid},
		{"bad checksum", "123456781", false, ErrBSNChecksumInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := DutchBSN(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("DutchBSN(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
