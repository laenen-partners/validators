package validators

import (
	"math/big"
	"regexp"
	"strings"
	"unicode"
)

var leiRegex = regexp.MustCompile(`^[A-Z0-9]{18}\d{2}$`)

// LEI validates a Legal Entity Identifier (ISO 17442).
// 20 alphanumeric characters with a MOD-97 check.
func LEI(value string) Result {
	if value == "" {
		return valid(nil)
	}

	code := strings.ToUpper(strings.TrimSpace(value))

	if len(code) != 20 {
		return invalid(ErrLEILengthInvalid, "LEI must be exactly 20 characters", "lei", map[string]any{
			"value":  value,
			"length": len(code),
		})
	}

	if !leiRegex.MatchString(code) {
		return invalid(ErrLEIFormatInvalid, "LEI contains invalid characters", "lei", map[string]any{
			"value": value,
		})
	}

	// MOD-97 check: convert letters to digits (A=10..Z=35), compute mod 97.
	var digits strings.Builder
	for _, c := range code {
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
		return invalid(ErrLEIChecksumInvalid, "Invalid LEI checksum", "lei", map[string]any{
			"value": value,
		})
	}

	return valid(map[string]any{
		"lou_prefix": code[:4],
	})
}
