package validators

import "testing"

func TestLEI(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid", "7ZW8QJWVPR4P1J1KQY45", true, ""},
		{"valid Bloomberg", "549300DTUYXVMJXZNY75", true, ""},
		{"too short", "7ZW8QJWVPR4P1J1K", false, ErrLEILengthInvalid},
		{"too long", "7ZW8QJWVPR4P1J1KQY451", false, ErrLEILengthInvalid},
		{"invalid chars", "7ZW8QJWVPR4P1J1K!Y45", false, ErrLEIFormatInvalid},
		{"bad checksum", "7ZW8QJWVPR4P1J1KQY46", false, ErrLEIChecksumInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := LEI(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("LEI(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
