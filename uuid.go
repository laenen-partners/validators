package validators

import (
	"fmt"
	"regexp"
	"strings"
)

var uuidRegex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)

// UUID validates a UUID string. If version > 0, also checks the version nibble.
func UUID(value string, version int) Result {
	if value == "" {
		return valid(nil)
	}

	lower := strings.ToLower(strings.TrimSpace(value))

	if !uuidRegex.MatchString(lower) {
		return invalid(ErrUUIDFormatInvalid, "Invalid UUID format", "uuid", map[string]any{
			"value": value,
		})
	}

	// The version nibble is at position 14 (the first char of the third group).
	detectedVersion := int(lower[14] - '0')

	if version > 0 && detectedVersion != version {
		return invalid(ErrUUIDVersionInvalid,
			fmt.Sprintf("Expected UUID version %d, got version %d", version, detectedVersion),
			"uuid", map[string]any{
				"value":            value,
				"expected_version": version,
				"actual_version":   detectedVersion,
			})
	}

	return valid(map[string]any{
		"version": detectedVersion,
	})
}
