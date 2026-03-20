package validators

import "testing"

func TestPhone(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid US", "+12025551234", true, ""},
		{"valid UK", "+442071234567", true, ""},
		{"valid BE", "+32412345678", true, ""},
		{"valid NL", "+31612345678", true, ""},
		{"valid DE", "+4930123456", true, ""},
		{"valid FR", "+33612345678", true, ""},
		{"valid with spaces", "+32 412 345 678", true, ""},
		{"valid with dashes", "+32-412-345-678", true, ""},
		{"valid with parens", "+32(412)345678", true, ""},
		{"missing plus", "32412345678", false, ErrPhoneFormatInvalid},
		{"too short", "+1234", false, ErrPhoneTooShort},
		{"too long", "+1234567890123456", false, ErrPhoneTooLong},
		{"letters", "+32abc345678", false, ErrPhoneInvalidChars},
		{"wrong length for country", "+3212345678901", false, ErrPhoneCountryInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Phone(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("Phone(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
