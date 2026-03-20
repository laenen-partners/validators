package validators

import "testing"

func TestBelgianNationalNumber(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid male 1985", "85073003328", true, ""},
		{"valid with dots and dash", "85.07.30-033.28", true, ""},
		{"too short", "8507300332", false, ErrBNNLengthInvalid},
		{"too long", "850730033280", false, ErrBNNLengthInvalid},
		{"non-digits", "8507300332a", false, ErrBNNFormatInvalid},
		{"bad checksum", "85073003329", false, ErrBNNChecksumInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := BelgianNationalNumber(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("BelgianNationalNumber(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}

func TestBelgianNationalNumber_Metadata(t *testing.T) {
	r := BelgianNationalNumber("85073003328")
	assertMetadata(t, r, "birth_year", 1985)
	assertMetadata(t, r, "birth_month", 7)
	assertMetadata(t, r, "birth_day", 30)
	assertMetadata(t, r, "gender", "male")
}
