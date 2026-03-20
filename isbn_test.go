package validators

import "testing"

func TestISBN(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid isbn-10", "0306406152", true, ""},
		{"valid isbn-10 with X", "007462542X", true, ""},
		{"valid isbn-10 dashes", "0-306-40615-2", true, ""},
		{"valid isbn-13", "9780306406157", true, ""},
		{"valid isbn-13 dashes", "978-0-306-40615-7", true, ""},
		{"valid isbn-13 979", "9791034304547", true, ""},
		{"wrong length", "12345", false, ErrISBNFormatInvalid},
		{"bad checksum isbn-10", "0306406153", false, ErrISBNChecksumInvalid},
		{"bad checksum isbn-13", "9780306406158", false, ErrISBNChecksumInvalid},
		{"invalid chars isbn-10", "030640615A", false, ErrISBNFormatInvalid},
		{"isbn-13 wrong prefix", "1234567890123", false, ErrISBNFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ISBN(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("ISBN(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
