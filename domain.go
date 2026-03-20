package validators

import (
	"fmt"
	"regexp"
	"strings"
)

var domainLabelRegex = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$`)

// Domain validates a domain name per RFC 1035.
func Domain(value string) Result {
	if value == "" {
		return valid(nil)
	}

	d := strings.TrimSpace(value)
	// Allow trailing dot (FQDN), but strip for validation.
	d = strings.TrimSuffix(d, ".")

	if len(d) > 253 {
		return invalid(ErrDomainTooLong,
			fmt.Sprintf("Domain name exceeds 253 characters (got %d)", len(d)),
			"domain", map[string]any{
				"value":  value,
				"length": len(d),
			})
	}

	labels := strings.Split(d, ".")
	if len(labels) < 2 {
		return invalid(ErrDomainFormatInvalid, "Domain must have at least two labels", "domain", map[string]any{
			"value": value,
		})
	}

	for _, label := range labels {
		if label == "" {
			return invalid(ErrDomainLabelInvalid, "Domain contains an empty label", "domain", map[string]any{
				"value": value,
			})
		}
		if len(label) > 63 {
			return invalid(ErrDomainLabelInvalid,
				fmt.Sprintf("Domain label exceeds 63 characters: %s", label),
				"domain", map[string]any{
					"value": value,
					"label": label,
				})
		}
		if !domainLabelRegex.MatchString(label) {
			return invalid(ErrDomainLabelInvalid,
				fmt.Sprintf("Invalid domain label: %s", label),
				"domain", map[string]any{
					"value": value,
					"label": label,
				})
		}
	}

	// TLD must be at least 2 characters and not all digits.
	tld := labels[len(labels)-1]
	allDigits := true
	for _, c := range tld {
		if c < '0' || c > '9' {
			allDigits = false
			break
		}
	}
	if len(tld) < 2 || allDigits {
		return invalid(ErrDomainFormatInvalid, "Domain must end with a valid TLD", "domain", map[string]any{
			"value": value,
			"tld":   tld,
		})
	}

	return valid(map[string]any{
		"labels": len(labels),
		"tld":    tld,
	})
}
