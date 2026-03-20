package validators

import "testing"

func TestCIDR(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid v4", "10.0.0.0/8", true, ""},
		{"valid v4 host", "192.168.1.1/32", true, ""},
		{"valid v6", "fd00::/64", true, ""},
		{"valid v6 full", "2001:db8::/32", true, ""},
		{"no prefix", "10.0.0.0", false, ErrCIDRFormatInvalid},
		{"bad IP", "999.999.999.999/8", false, ErrCIDRFormatInvalid},
		{"prefix too large v4", "10.0.0.0/33", false, ErrCIDRPrefixInvalid},
		{"prefix too large v6", "fd00::/129", false, ErrCIDRPrefixInvalid},
		{"negative prefix", "10.0.0.0/-1", false, ErrCIDRPrefixInvalid},
		{"non-numeric prefix", "10.0.0.0/abc", false, ErrCIDRPrefixInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CIDR(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("CIDR(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
