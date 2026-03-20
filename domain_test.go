package validators

import "testing"

func TestDomain(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid simple", "example.com", true, ""},
		{"valid subdomain", "sub.example.com", true, ""},
		{"valid deep subdomain", "a.b.c.example.com", true, ""},
		{"valid with hyphen", "my-domain.com", true, ""},
		{"valid trailing dot", "example.com.", true, ""},
		{"single label", "localhost", false, ErrDomainFormatInvalid},
		{"empty label", "example..com", false, ErrDomainLabelInvalid},
		{"label starts with hyphen", "-example.com", false, ErrDomainLabelInvalid},
		{"label ends with hyphen", "example-.com", false, ErrDomainLabelInvalid},
		{"numeric TLD", "example.123", false, ErrDomainFormatInvalid},
		{"single char TLD", "example.a", false, ErrDomainFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Domain(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("Domain(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
