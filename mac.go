package validators

import (
	"regexp"
	"strings"
)

var macColonRegex = regexp.MustCompile(`^([0-9A-Fa-f]{2}:){5}[0-9A-Fa-f]{2}$`)
var macHyphenRegex = regexp.MustCompile(`^([0-9A-Fa-f]{2}-){5}[0-9A-Fa-f]{2}$`)
var macDotRegex = regexp.MustCompile(`^[0-9A-Fa-f]{4}\.[0-9A-Fa-f]{4}\.[0-9A-Fa-f]{4}$`)

// MAC validates a 48-bit MAC address in colon, hyphen, or dot notation.
func MAC(value string) Result {
	if value == "" {
		return valid(nil)
	}

	trimmed := strings.TrimSpace(value)

	var format string
	switch {
	case macColonRegex.MatchString(trimmed):
		format = "colon"
	case macHyphenRegex.MatchString(trimmed):
		format = "hyphen"
	case macDotRegex.MatchString(trimmed):
		format = "dot"
	default:
		return invalid(ErrMACFormatInvalid, "Invalid MAC address format (expected XX:XX:XX:XX:XX:XX, XX-XX-XX-XX-XX-XX, or XXXX.XXXX.XXXX)", "mac", map[string]any{
			"value": value,
		})
	}

	return valid(map[string]any{
		"format": format,
	})
}
