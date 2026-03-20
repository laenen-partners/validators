package validators

import "testing"

func TestUUID(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		version int
		valid   bool
		code    string
	}{
		{"empty", "", 0, true, ""},
		{"valid v4", "550e8400-e29b-41d4-a716-446655440000", 0, true, ""},
		{"valid v4 uppercase", "550E8400-E29B-41D4-A716-446655440000", 0, true, ""},
		{"valid v4 check version", "550e8400-e29b-41d4-a716-446655440000", 4, true, ""},
		{"wrong version", "550e8400-e29b-41d4-a716-446655440000", 1, false, ErrUUIDVersionInvalid},
		{"no dashes", "550e8400e29b41d4a716446655440000", 0, false, ErrUUIDFormatInvalid},
		{"too short", "550e8400-e29b-41d4-a716", 0, false, ErrUUIDFormatInvalid},
		{"invalid chars", "550e8400-e29b-41d4-a716-44665544gggg", 0, false, ErrUUIDFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := UUID(tt.input, tt.version)
			if r.Valid != tt.valid {
				t.Errorf("UUID(%q, %d) valid=%v, want %v", tt.input, tt.version, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
