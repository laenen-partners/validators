package validators

import "testing"

func TestURL(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid https", "https://example.com", true, ""},
		{"valid http", "http://example.com", true, ""},
		{"valid with path", "https://example.com/path/to/page", true, ""},
		{"valid with query", "https://example.com?q=test", true, ""},
		{"valid with port", "https://example.com:8080", true, ""},
		{"valid with fragment", "https://example.com#section", true, ""},
		{"valid ftp", "ftp://files.example.com/file.txt", true, ""},
		{"no scheme", "example.com", false, ErrURLSchemeInvalid},
		{"invalid scheme", "gopher://example.com", false, ErrURLSchemeInvalid},
		{"no host", "https://", false, ErrURLHostMissing},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := URL(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("URL(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}

func TestURL_Metadata(t *testing.T) {
	r := URL("https://example.com:8080/api")
	assertMetadata(t, r, "scheme", "https")
	assertMetadata(t, r, "host", "example.com:8080")
	assertMetadata(t, r, "port", "8080")
	assertMetadata(t, r, "path", "/api")
}
