package validators

import "testing"

func TestSemVer(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid simple", "1.2.3", true, ""},
		{"valid with v prefix", "v1.2.3", true, ""},
		{"valid with prerelease", "1.2.3-alpha.1", true, ""},
		{"valid with build", "1.2.3+build.123", true, ""},
		{"valid with both", "1.2.3-beta.2+build.456", true, ""},
		{"valid zeros", "0.0.0", true, ""},
		{"valid large", "100.200.300", true, ""},
		{"missing patch", "1.2", false, ErrSemVerFormatInvalid},
		{"extra segment", "1.2.3.4", false, ErrSemVerFormatInvalid},
		{"leading zero major", "01.2.3", false, ErrSemVerFormatInvalid},
		{"leading zero minor", "1.02.3", false, ErrSemVerFormatInvalid},
		{"letters only", "abc", false, ErrSemVerFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := SemVer(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("SemVer(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}

func TestSemVer_Metadata(t *testing.T) {
	r := SemVer("1.2.3-alpha.1+build.456")
	assertMetadata(t, r, "major", 1)
	assertMetadata(t, r, "minor", 2)
	assertMetadata(t, r, "patch", 3)
	assertMetadata(t, r, "prerelease", "alpha.1")
	assertMetadata(t, r, "build", "build.456")
}
