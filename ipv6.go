package validators

import (
	"net"
	"strings"
)

// IPv6 validates an IPv6 address.
func IPv6(value string) Result {
	if value == "" {
		return valid(nil)
	}

	trimmed := strings.TrimSpace(value)

	// Reject if it looks like IPv4.
	if !strings.Contains(trimmed, ":") {
		return invalid(ErrIPv6FormatInvalid, "Invalid IPv6 address format", "ipv6", map[string]any{
			"value": value,
		})
	}

	// Strip zone ID (e.g., "%eth0") for parsing.
	addr := trimmed
	if idx := strings.Index(addr, "%"); idx != -1 {
		addr = addr[:idx]
	}

	ip := net.ParseIP(addr)
	if ip == nil {
		return invalid(ErrIPv6FormatInvalid, "Invalid IPv6 address format", "ipv6", map[string]any{
			"value": value,
		})
	}

	// Ensure it's actually IPv6, not an IPv4 parsed by net.ParseIP.
	if ip.To4() != nil && !strings.Contains(addr, ":") {
		return invalid(ErrIPv6FormatInvalid, "Value is an IPv4 address, not IPv6", "ipv6", map[string]any{
			"value": value,
		})
	}

	return valid(map[string]any{
		"address": ip.String(),
	})
}
