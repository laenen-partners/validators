package validators

import "testing"

func TestDate(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid", "2024-01-15", true, ""},
		{"valid leap day", "2024-02-29", true, ""},
		{"valid end of month", "2024-12-31", true, ""},
		{"invalid leap day", "2023-02-29", false, ErrDateInvalid},
		{"invalid month", "2024-13-01", false, ErrDateInvalid},
		{"invalid day", "2024-01-32", false, ErrDateInvalid},
		{"invalid day zero", "2024-01-00", false, ErrDateInvalid},
		{"invalid month zero", "2024-00-15", false, ErrDateInvalid},
		{"wrong format", "01/15/2024", false, ErrDateFormatInvalid},
		{"wrong format short year", "24-01-15", false, ErrDateFormatInvalid},
		{"not a date", "hello", false, ErrDateFormatInvalid},
		{"april 31", "2024-04-31", false, ErrDateInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Date(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("Date(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
