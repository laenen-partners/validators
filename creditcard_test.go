package validators

import "testing"

func TestCreditCard(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
		code  string
	}{
		{"empty", "", true, ""},
		{"valid visa", "4111111111111111", true, ""},
		{"valid mastercard", "5500000000000004", true, ""},
		{"valid amex", "340000000000009", true, ""},
		{"valid with spaces", "4111 1111 1111 1111", true, ""},
		{"valid with dashes", "4111-1111-1111-1111", true, ""},
		{"too short", "41111111", false, ErrCreditCardTooShort},
		{"too long", "41111111111111111111", false, ErrCreditCardTooLong},
		{"invalid chars", "4111abcd11111111", false, ErrCreditCardInvalidChars},
		{"bad luhn", "4111111111111112", false, ErrCreditCardChecksumInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CreditCard(tt.input)
			if r.Valid != tt.valid {
				t.Errorf("CreditCard(%q) valid=%v, want %v", tt.input, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}

func TestCreditCard_Network(t *testing.T) {
	tests := []struct {
		input   string
		network string
	}{
		{"4111111111111111", "visa"},
		{"5500000000000004", "mastercard"},
		{"340000000000009", "amex"},
		{"6011000000000004", "discover"},
	}
	for _, tt := range tests {
		t.Run(tt.network, func(t *testing.T) {
			r := CreditCard(tt.input)
			if r.Metadata["network"] != tt.network {
				t.Errorf("CreditCard(%q) network=%v, want %v", tt.input, r.Metadata["network"], tt.network)
			}
		})
	}
}
