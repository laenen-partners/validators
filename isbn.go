package validators

import (
	"strings"
	"unicode"
)

// ISBN validates an ISBN-10 or ISBN-13.
func ISBN(value string) Result {
	if value == "" {
		return valid(nil)
	}

	// Strip hyphens and spaces.
	cleaned := strings.Map(func(r rune) rune {
		if r == '-' || r == ' ' {
			return -1
		}
		return r
	}, value)

	switch len(cleaned) {
	case 10:
		return validateISBN10(value, cleaned)
	case 13:
		return validateISBN13(value, cleaned)
	default:
		return invalid(ErrISBNFormatInvalid, "ISBN must be 10 or 13 digits", "isbn", map[string]any{
			"value":  value,
			"length": len(cleaned),
		})
	}
}

func validateISBN10(original, cleaned string) Result {
	// ISBN-10: first 9 must be digits, last can be digit or X.
	for i := 0; i < 9; i++ {
		if !unicode.IsDigit(rune(cleaned[i])) {
			return invalid(ErrISBNFormatInvalid, "ISBN-10 contains invalid characters", "isbn", map[string]any{
				"value": original,
			})
		}
	}
	last := rune(cleaned[9])
	if !unicode.IsDigit(last) && last != 'X' && last != 'x' {
		return invalid(ErrISBNFormatInvalid, "ISBN-10 check digit must be 0-9 or X", "isbn", map[string]any{
			"value": original,
		})
	}

	// Checksum: sum of (10*d1 + 9*d2 + ... + 1*d10) mod 11 == 0.
	sum := 0
	for i := 0; i < 9; i++ {
		sum += (10 - i) * int(cleaned[i]-'0')
	}
	if last == 'X' || last == 'x' {
		sum += 10
	} else {
		sum += int(last - '0')
	}

	if sum%11 != 0 {
		return invalid(ErrISBNChecksumInvalid, "Invalid ISBN-10 checksum", "isbn", map[string]any{
			"value":  original,
			"format": "isbn-10",
		})
	}

	return valid(map[string]any{
		"format": "isbn-10",
	})
}

func validateISBN13(original, cleaned string) Result {
	for _, c := range cleaned {
		if !unicode.IsDigit(c) {
			return invalid(ErrISBNFormatInvalid, "ISBN-13 contains invalid characters", "isbn", map[string]any{
				"value": original,
			})
		}
	}

	// Must start with 978 or 979.
	if !strings.HasPrefix(cleaned, "978") && !strings.HasPrefix(cleaned, "979") {
		return invalid(ErrISBNFormatInvalid, "ISBN-13 must start with 978 or 979", "isbn", map[string]any{
			"value": original,
		})
	}

	// Checksum: alternating weights 1 and 3, sum mod 10 == 0.
	sum := 0
	for i, c := range cleaned {
		d := int(c - '0')
		if i%2 == 0 {
			sum += d
		} else {
			sum += d * 3
		}
	}

	if sum%10 != 0 {
		return invalid(ErrISBNChecksumInvalid, "Invalid ISBN-13 checksum", "isbn", map[string]any{
			"value":  original,
			"format": "isbn-13",
		})
	}

	return valid(map[string]any{
		"format": "isbn-13",
	})
}
