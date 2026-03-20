package validators

import "testing"

func TestStringLength(t *testing.T) {
	tests := []struct {
		name  string
		value string
		min   int
		max   int
		valid bool
		code  string
	}{
		{"empty no min", "", 0, 100, true, ""},
		{"empty with min", "", 1, 100, false, ErrStringTooShort},
		{"within range", "hello", 1, 10, true, ""},
		{"at min", "a", 1, 10, true, ""},
		{"at max", "1234567890", 1, 10, true, ""},
		{"too short", "ab", 3, 10, false, ErrStringTooShort},
		{"too long", "12345678901", 1, 10, false, ErrStringTooLong},
		{"no max", "a very long string indeed", 1, 0, true, ""},
		{"unicode runes", "héllo", 1, 5, true, ""},
		{"emoji counts as 1", "😀😀😀", 1, 3, true, ""},
		{"CJK characters", "你好世界", 1, 4, true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := StringLength(tt.value, tt.min, tt.max)
			if r.Valid != tt.valid {
				t.Errorf("StringLength(%q, %d, %d) valid=%v, want %v", tt.value, tt.min, tt.max, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}

func TestStringLength_Metadata(t *testing.T) {
	r := StringLength("héllo", 1, 10)
	assertMetadata(t, r, "length", 5)
}
