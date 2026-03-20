package validators

import (
	"net"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$`)

// Email validates an email address.
// If checkMX is true, it verifies the domain has MX records.
func Email(value string, checkMX bool) Result {
	if value == "" {
		return valid(nil)
	}

	// RFC 5321: total address must not exceed 254 characters.
	if len(value) > 254 {
		return invalid(ErrEmailFormatInvalid, "Email address exceeds 254 characters", "email", map[string]any{
			"value":  value,
			"length": len(value),
		})
	}

	if !emailRegex.MatchString(value) {
		return invalid(ErrEmailFormatInvalid, "Invalid email format", "email", map[string]any{
			"value": value,
		})
	}

	// The regex guarantees exactly one @, so SplitN is safe here.
	parts := strings.SplitN(value, "@", 2)
	local := parts[0]
	domain := parts[1]

	// RFC 5321: local part must not exceed 64 characters.
	if len(local) > 64 {
		return invalid(ErrEmailFormatInvalid, "Email local part exceeds 64 characters", "email", map[string]any{
			"value":        value,
			"local_length": len(local),
		})
	}

	if checkMX {
		mxRecords, err := net.LookupMX(domain)
		if err != nil || len(mxRecords) == 0 {
			return invalid(ErrEmailDomainNoMX, "Domain does not accept email", "email", map[string]any{
				"value":  value,
				"domain": domain,
			})
		}
	}

	return valid(map[string]any{
		"domain": domain,
	})
}
