package validators

import "testing"

func TestEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		checkMX bool
		valid   bool
		code    string
	}{
		// Valid emails
		{name: "simple email", email: "test@example.com", valid: true},
		{name: "with subdomain", email: "test@mail.example.com", valid: true},
		{name: "with plus", email: "test+tag@example.com", valid: true},
		{name: "with dots in local", email: "first.last@example.com", valid: true},
		{name: "short TLD", email: "test@example.io", valid: true},
		{name: "long TLD", email: "test@example.museum", valid: true},
		{name: "numeric domain", email: "test@123.com", valid: true},
		{name: "hyphen in domain", email: "test@my-domain.com", valid: true},
		{name: "empty is valid", email: "", valid: true},

		// Invalid emails
		{name: "no TLD", email: "pascal@laenen", valid: false, code: ErrEmailFormatInvalid},
		{name: "no domain", email: "test@", valid: false, code: ErrEmailFormatInvalid},
		{name: "no at symbol", email: "testexample.com", valid: false, code: ErrEmailFormatInvalid},
		{name: "double at", email: "test@@example.com", valid: false, code: ErrEmailFormatInvalid},
		{name: "space in email", email: "test @example.com", valid: false, code: ErrEmailFormatInvalid},
		{name: "no local part", email: "@example.com", valid: false, code: ErrEmailFormatInvalid},
		{name: "trailing dot", email: "test@example.com.", valid: false, code: ErrEmailFormatInvalid},
		{name: "leading dot in domain", email: "test@.example.com", valid: false, code: ErrEmailFormatInvalid},
		{name: "double dot in domain", email: "test@example..com", valid: false, code: ErrEmailFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Email(tt.email, tt.checkMX)
			if result.Valid != tt.valid {
				t.Errorf("Email(%q) valid = %v, want %v", tt.email, result.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, result, tt.code)
			}
		})
	}
}

func TestEmail_Metadata(t *testing.T) {
	result := Email("user@example.com", false)
	if result.Metadata["domain"] != "example.com" {
		t.Errorf("Email domain = %v, want %q", result.Metadata["domain"], "example.com")
	}
}

func TestEmail_MXCheck(t *testing.T) {
	result := Email("test@gmail.com", true)
	if !result.Valid {
		t.Errorf("Email with MX check for gmail.com should be valid")
	}

	result = Email("test@thisdomain-does-not-exist-xyz123.com", true)
	if result.Valid {
		t.Error("Email with MX check for non-existent domain should be invalid")
	}
	assertErrorCode(t, result, ErrEmailDomainNoMX)
}
