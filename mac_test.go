package validators

import "testing"

func TestMAC(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid colon", "00:1A:2B:3C:4D:5E", true, ""},
		{"valid colon lowercase", "00:1a:2b:3c:4d:5e", true, ""},
		{"valid hyphen", "00-1A-2B-3C-4D-5E", true, ""},
		{"valid dot", "001A.2B3C.4D5E", true, ""},
		{"invalid format", "001A2B3C4D5E", false, ErrMACFormatInvalid},
		{"too short", "00:1A:2B", false, ErrMACFormatInvalid},
		{"invalid chars", "00:1A:2B:3C:4D:GG", false, ErrMACFormatInvalid},
		{"mixed separators", "00:1A-2B:3C:4D:5E", false, ErrMACFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := MAC(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("MAC(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
