package validators

import (
	"fmt"
	"strconv"
	"strings"
)

// IPv4 validates an IPv4 address in dotted-decimal notation.
func IPv4(value string) Result {
	if value == "" {
		return valid(nil)
	}

	parts := strings.Split(value, ".")
	if len(parts) != 4 {
		return invalid(ErrIPv4FormatInvalid, "IPv4 address must have exactly 4 octets", "ipv4", map[string]any{
			"value":  value,
			"octets": len(parts),
		})
	}

	octets := make([]int, 4)
	for i, part := range parts {
		if part == "" {
			return invalid(ErrIPv4FormatInvalid, "IPv4 address contains empty octet", "ipv4", map[string]any{
				"value":    value,
				"position": i + 1,
			})
		}
		// Reject leading zeros (e.g., "01", "001").
		if len(part) > 1 && part[0] == '0' {
			return invalid(ErrIPv4OctetInvalid,
				fmt.Sprintf("IPv4 octet %d has leading zeros: %s", i+1, part),
				"ipv4", map[string]any{
					"value":    value,
					"position": i + 1,
					"octet":    part,
				})
		}
		n, err := strconv.Atoi(part)
		if err != nil || n < 0 || n > 255 {
			return invalid(ErrIPv4OctetInvalid,
				fmt.Sprintf("IPv4 octet %d is out of range (0-255): %s", i+1, part),
				"ipv4", map[string]any{
					"value":    value,
					"position": i + 1,
					"octet":    part,
				})
		}
		octets[i] = n
	}

	return valid(map[string]any{
		"octets": octets,
	})
}
