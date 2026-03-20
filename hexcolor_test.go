package validators

import "testing"

func TestHexColor(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid 3", "#fff", true, ""},
		{"valid 4", "#fffa", true, ""},
		{"valid 6", "#ff00aa", true, ""},
		{"valid 8", "#ff00aa80", true, ""},
		{"valid uppercase", "#FF00AA", true, ""},
		{"valid mixed case", "#Ff00aA", true, ""},
		{"no hash", "ff00aa", false, ErrHexColorFormatInvalid},
		{"wrong length 5", "#ff00a", false, ErrHexColorFormatInvalid},
		{"wrong length 7", "#ff00aa8", false, ErrHexColorFormatInvalid},
		{"invalid chars", "#gghhii", false, ErrHexColorFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := HexColor(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("HexColor(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
