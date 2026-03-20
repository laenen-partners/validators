package validators

import "testing"

func TestIPv6(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid full", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", true, ""},
		{"valid compressed", "2001:db8::1", true, ""},
		{"valid loopback", "::1", true, ""},
		{"valid all zeros", "::", true, ""},
		{"valid link-local", "fe80::1", true, ""},
		{"invalid", "not-an-ipv6", false, ErrIPv6FormatInvalid},
		{"ipv4", "192.168.1.1", false, ErrIPv6FormatInvalid},
		{"too many groups", "2001:db8:85a3:0:0:8a2e:370:7334:extra", false, ErrIPv6FormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := IPv6(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("IPv6(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
