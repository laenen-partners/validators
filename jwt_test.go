package validators

import (
	"encoding/base64"
	"testing"
)

func TestJWT(t *testing.T) {
	// Build a valid JWT for testing.
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	payload := base64.RawURLEncoding.EncodeToString([]byte(`{"sub":"1234567890","name":"Test"}`))
	sig := base64.RawURLEncoding.EncodeToString([]byte("signature"))
	validJWT := header + "." + payload + "." + sig

	// Header without alg.
	badHeader := base64.RawURLEncoding.EncodeToString([]byte(`{"typ":"JWT"}`))
	noAlgJWT := badHeader + "." + payload + "." + sig

	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid", validJWT, true, ""},
		{"only two segments", "abc.def", false, ErrJWTFormatInvalid},
		{"four segments", "a.b.c.d", false, ErrJWTFormatInvalid},
		{"empty header", "." + payload + "." + sig, false, ErrJWTSegmentInvalid},
		{"bad base64 header", "!!!." + payload + "." + sig, false, ErrJWTSegmentInvalid},
		{"no alg", noAlgJWT, false, ErrJWTSegmentInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := JWT(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("JWT(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}

func TestJWT_Metadata(t *testing.T) {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"RS256","typ":"JWT"}`))
	payload := base64.RawURLEncoding.EncodeToString([]byte(`{"sub":"123"}`))
	sig := base64.RawURLEncoding.EncodeToString([]byte("sig"))

	r := JWT(header + "." + payload + "." + sig)
	assertMetadata(t, r, "algorithm", "RS256")
	assertMetadata(t, r, "type", "JWT")
}
