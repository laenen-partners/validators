package validators

import "testing"

func TestCRON(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"every minute", "* * * * *", true, ""},
		{"specific time", "30 8 * * 1", true, ""},
		{"range", "0 9-17 * * *", true, ""},
		{"step", "*/15 * * * *", true, ""},
		{"list", "0 8,12,18 * * *", true, ""},
		{"complex", "0,30 9-17/2 1-15 1,6 1-5", true, ""},
		{"too few fields", "* * *", false, ErrCRONFormatInvalid},
		{"too many fields", "* * * * * *", false, ErrCRONFormatInvalid},
		{"minute out of range", "60 * * * *", false, ErrCRONFieldInvalid},
		{"hour out of range", "0 24 * * *", false, ErrCRONFieldInvalid},
		{"day out of range", "0 0 32 * *", false, ErrCRONFieldInvalid},
		{"month out of range", "0 0 * 13 *", false, ErrCRONFieldInvalid},
		{"weekday out of range", "0 0 * * 8", false, ErrCRONFieldInvalid},
		{"non-numeric", "abc * * * *", false, ErrCRONFieldInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CRON(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("CRON(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
