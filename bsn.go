package validators

import (
	"fmt"
	"strings"
	"unicode"
)

// DutchBSN validates a Dutch Burgerservicenummer (9 digits, 11-check).
func DutchBSN(value string) Result {
	if value == "" {
		return valid(nil)
	}

	// Strip dots, dashes, and spaces.
	cleaned := strings.Map(func(r rune) rune {
		if r == '.' || r == '-' || r == ' ' {
			return -1
		}
		return r
	}, value)

	if len(cleaned) != 9 {
		return invalid(ErrBSNLengthInvalid,
			fmt.Sprintf("BSN must be 9 digits, got %d", len(cleaned)),
			"bsn", map[string]any{
				"value":  value,
				"length": len(cleaned),
			})
	}

	for _, c := range cleaned {
		if !unicode.IsDigit(c) {
			return invalid(ErrBSNFormatInvalid, "BSN must contain only digits", "bsn", map[string]any{
				"value": value,
			})
		}
	}

	// 11-check: 9*d1 + 8*d2 + 7*d3 + 6*d4 + 5*d5 + 4*d6 + 3*d7 + 2*d8 - 1*d9
	// must be divisible by 11 and not zero.
	weights := []int{9, 8, 7, 6, 5, 4, 3, 2, -1}
	sum := 0
	for i, w := range weights {
		sum += w * int(cleaned[i]-'0')
	}

	if sum <= 0 || sum%11 != 0 {
		return invalid(ErrBSNChecksumInvalid, "Invalid BSN checksum", "bsn", map[string]any{
			"value": value,
		})
	}

	return valid(nil)
}
