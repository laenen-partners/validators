package validators

import "testing"

func TestSWIFT(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid 8 char", "DEUTDEFF", true, ""},
		{"valid 11 char", "DEUTDEFF500", true, ""},
		{"valid lowercase", "deutdeff", true, ""},
		{"too short", "DEUT", false, ErrSWIFTLengthInvalid},
		{"wrong length 9", "DEUTDEFF5", false, ErrSWIFTLengthInvalid},
		{"invalid chars", "DEUT12FF", false, ErrSWIFTFormatInvalid},
		{"numbers in bank code", "1234DEFF", false, ErrSWIFTFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := SWIFT(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("SWIFT(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}

func TestSWIFT_Metadata(t *testing.T) {
	r := SWIFT("DEUTDEFF500")
	if r.Metadata["bank_code"] != "DEUT" {
		t.Errorf("SWIFT bank_code = %v, want DEUT", r.Metadata["bank_code"])
	}
	if r.Metadata["country_code"] != "DE" {
		t.Errorf("SWIFT country_code = %v, want DE", r.Metadata["country_code"])
	}
}
