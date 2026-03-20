package validators

import "testing"

func TestIPv4(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid", "192.168.1.1", true, ""},
		{"valid zeros", "0.0.0.0", true, ""},
		{"valid max", "255.255.255.255", true, ""},
		{"valid loopback", "127.0.0.1", true, ""},
		{"too few octets", "192.168.1", false, ErrIPv4FormatInvalid},
		{"too many octets", "192.168.1.1.1", false, ErrIPv4FormatInvalid},
		{"octet too high", "256.0.0.1", false, ErrIPv4OctetInvalid},
		{"leading zeros", "192.168.01.1", false, ErrIPv4OctetInvalid},
		{"empty octet", "192..1.1", false, ErrIPv4FormatInvalid},
		{"negative", "192.168.-1.1", false, ErrIPv4OctetInvalid},
		{"letters", "abc.def.ghi.jkl", false, ErrIPv4OctetInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := IPv4(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("IPv4(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
