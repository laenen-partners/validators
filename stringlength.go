package validators

import (
	"fmt"
	"unicode/utf8"
)

// StringLength validates that a string's length in Unicode rune count falls within [min, max].
// Uses rune count, not byte count, so multi-byte characters are counted correctly.
// Set min to 0 to skip minimum check. Set max to 0 to skip maximum check.
func StringLength(value string, min, max int) Result {
	if value == "" && min == 0 {
		return valid(nil)
	}

	length := utf8.RuneCountInString(value)

	if min > 0 && length < min {
		return invalid(ErrStringTooShort,
			fmt.Sprintf("Must be at least %d characters, got %d", min, length),
			"string", map[string]any{
				"value":      value,
				"length":     length,
				"min_length": min,
			})
	}

	if max > 0 && length > max {
		return invalid(ErrStringTooLong,
			fmt.Sprintf("Must be at most %d characters, got %d", max, length),
			"string", map[string]any{
				"value":      value,
				"length":     length,
				"max_length": max,
			})
	}

	return valid(map[string]any{
		"length": length,
	})
}
