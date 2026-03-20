package validators

import (
	"fmt"
	"regexp"
	"strings"
)

var swiftRegex = regexp.MustCompile(`^[A-Z]{4}[A-Z]{2}[A-Z0-9]{2}([A-Z0-9]{3})?$`)

// SWIFT validates a SWIFT/BIC code.
func SWIFT(value string) Result {
	if value == "" {
		return valid(nil)
	}

	code := strings.ToUpper(strings.TrimSpace(value))

	if len(code) != 8 && len(code) != 11 {
		return invalid(ErrSWIFTLengthInvalid,
			fmt.Sprintf("SWIFT code must be 8 or 11 characters, got %d", len(code)),
			"swift", map[string]any{
				"value":         value,
				"actual_length": len(code),
			})
	}

	if !swiftRegex.MatchString(code) {
		return invalid(ErrSWIFTFormatInvalid, "Invalid SWIFT code format", "swift", map[string]any{
			"value": value,
		})
	}

	return valid(map[string]any{
		"bank_code":    code[:4],
		"country_code": code[4:6],
		"location":     code[6:8],
	})
}
