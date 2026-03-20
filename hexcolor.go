package validators

import (
	"regexp"
	"strings"
)

var hexColorRegex = regexp.MustCompile(`^#([0-9A-Fa-f]{3}|[0-9A-Fa-f]{4}|[0-9A-Fa-f]{6}|[0-9A-Fa-f]{8})$`)

// HexColor validates a CSS hex color code (#RGB, #RGBA, #RRGGBB, or #RRGGBBAA).
func HexColor(value string) Result {
	if value == "" {
		return valid(nil)
	}

	trimmed := strings.TrimSpace(value)

	if !hexColorRegex.MatchString(trimmed) {
		return invalid(ErrHexColorFormatInvalid, "Invalid hex color format (expected #RGB, #RGBA, #RRGGBB, or #RRGGBBAA)", "hex_color", map[string]any{
			"value": value,
		})
	}

	hex := trimmed[1:] // strip #
	var format string
	switch len(hex) {
	case 3:
		format = "shorthand"
	case 4:
		format = "shorthand-alpha"
	case 6:
		format = "full"
	case 8:
		format = "full-alpha"
	}

	return valid(map[string]any{
		"format":    format,
		"has_alpha": len(hex) == 4 || len(hex) == 8,
	})
}
