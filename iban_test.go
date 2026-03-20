package validators

import "testing"

func TestIBAN(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid GB", "GB29 NWBK 6016 1331 9268 19", true, ""},
		{"valid DE", "DE89370400440532013000", true, ""},
		{"valid NL", "NL91ABNA0417164300", true, ""},
		{"valid BE", "BE68539007547034", true, ""},
		{"valid FR", "FR7630006000011234567890189", true, ""},
		{"valid ES", "ES9121000418450200051332", true, ""},
		{"too short", "GB29", false, ErrIBANTooShort},
		{"bad checksum", "GB00 NWBK 6016 1331 9268 19", false, ErrIBANChecksumInvalid},
		{"wrong length for country", "GB29 NWBK 6016 13", false, ErrIBANLengthMismatch},
		{"invalid chars", "GB29!NWBK60161331926819", false, ErrIBANInvalidChars},
		{"lowercase valid", "gb29 nwbk 6016 1331 9268 19", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := IBAN(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("IBAN(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}

func TestIBAN_Metadata(t *testing.T) {
	r := IBAN("DE89370400440532013000")
	if r.Metadata["country_code"] != "DE" {
		t.Errorf("IBAN country_code = %v, want DE", r.Metadata["country_code"])
	}
}
