package validators

import (
	"fmt"
	"math/big"
	"strings"
	"unicode"
)

var ibanLengths = map[string]int{
	"AL": 28, "AD": 24, "AT": 20, "AZ": 28, "BH": 22,
	"BY": 28, "BE": 16, "BA": 20, "BR": 29, "BG": 22,
	"CR": 22, "HR": 21, "CY": 28, "CZ": 24, "DK": 18,
	"DO": 28, "TL": 23, "EG": 29, "SV": 28, "EE": 20,
	"FO": 18, "FI": 18, "FR": 27, "GE": 22, "DE": 22,
	"GI": 23, "GR": 27, "GL": 18, "GT": 28, "HU": 28,
	"IS": 26, "IQ": 23, "IE": 22, "IL": 23, "IT": 27,
	"JO": 30, "KZ": 20, "XK": 20, "KW": 30, "LV": 21,
	"LB": 28, "LI": 21, "LT": 20, "LU": 20, "MT": 31,
	"MR": 27, "MU": 30, "MC": 27, "MD": 24, "ME": 22,
	"NL": 18, "MK": 19, "NO": 15, "PK": 24, "PS": 29,
	"PL": 28, "PT": 25, "QA": 29, "RO": 24, "LC": 32,
	"SM": 27, "ST": 25, "SA": 24, "RS": 22, "SC": 31,
	"SK": 24, "SI": 19, "ES": 24, "SD": 18, "SE": 24,
	"CH": 21, "TN": 24, "TR": 26, "UA": 29, "AE": 23,
	"GB": 22, "VA": 22, "VG": 24,
}

// IBAN validates an International Bank Account Number.
func IBAN(value string) Result {
	if value == "" {
		return valid(nil)
	}

	iban := strings.ToUpper(strings.ReplaceAll(value, " ", ""))

	if len(iban) < 5 {
		return invalid(ErrIBANTooShort, "IBAN is too short", "iban", map[string]any{
			"value":  value,
			"length": len(iban),
		})
	}

	for _, c := range iban {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return invalid(ErrIBANInvalidChars, "IBAN contains invalid characters", "iban", map[string]any{
				"value": value,
			})
		}
	}

	cc := iban[:2]
	if !unicode.IsLetter(rune(cc[0])) || !unicode.IsLetter(rune(cc[1])) {
		return invalid(ErrIBANCountryInvalid, "Invalid country code", "iban", map[string]any{
			"value":        value,
			"country_code": cc,
		})
	}

	if expected, ok := ibanLengths[cc]; ok {
		if len(iban) != expected {
			return invalid(ErrIBANLengthMismatch,
				fmt.Sprintf("IBAN length for %s must be %d, got %d", cc, expected, len(iban)),
				"iban", map[string]any{
					"value":           value,
					"country_code":    cc,
					"expected_length": expected,
					"actual_length":   len(iban),
				})
		}
	}

	rearranged := iban[4:] + iban[:4]
	var digits strings.Builder
	for _, c := range rearranged {
		if unicode.IsLetter(c) {
			digits.WriteString(big.NewInt(int64(c - 'A' + 10)).String())
		} else {
			digits.WriteByte(byte(c))
		}
	}

	num := new(big.Int)
	num.SetString(digits.String(), 10)
	mod := new(big.Int).Mod(num, big.NewInt(97))
	if mod.Int64() != 1 {
		return invalid(ErrIBANChecksumInvalid, "Invalid IBAN checksum", "iban", map[string]any{
			"value":        value,
			"country_code": cc,
		})
	}

	return valid(map[string]any{
		"country_code": cc,
	})
}
