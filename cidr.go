package validators

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// CIDR validates an IP address in CIDR notation (e.g., "10.0.0.0/8", "fd00::/64").
func CIDR(value string) Result {
	if value == "" {
		return valid(nil)
	}

	trimmed := strings.TrimSpace(value)

	parts := strings.SplitN(trimmed, "/", 2)
	if len(parts) != 2 {
		return invalid(ErrCIDRFormatInvalid, "CIDR must contain an IP address and prefix length separated by /", "cidr", map[string]any{
			"value": value,
		})
	}

	ip := net.ParseIP(parts[0])
	if ip == nil {
		return invalid(ErrCIDRFormatInvalid, "CIDR contains an invalid IP address", "cidr", map[string]any{
			"value": value,
			"ip":    parts[0],
		})
	}

	prefix, err := strconv.Atoi(parts[1])
	if err != nil || prefix < 0 {
		return invalid(ErrCIDRPrefixInvalid, "CIDR prefix length must be a non-negative integer", "cidr", map[string]any{
			"value":  value,
			"prefix": parts[1],
		})
	}

	// Determine max prefix based on IP version.
	isV6 := ip.To4() == nil
	maxPrefix := 32
	ipVersion := "v4"
	if isV6 {
		maxPrefix = 128
		ipVersion = "v6"
	}

	if prefix > maxPrefix {
		return invalid(ErrCIDRPrefixInvalid,
			fmt.Sprintf("CIDR prefix length %d exceeds maximum %d for IP%s", prefix, maxPrefix, ipVersion),
			"cidr", map[string]any{
				"value":      value,
				"prefix":     prefix,
				"max_prefix": maxPrefix,
				"ip_version": ipVersion,
			})
	}

	return valid(map[string]any{
		"ip":         ip.String(),
		"prefix":     prefix,
		"ip_version": ipVersion,
	})
}
