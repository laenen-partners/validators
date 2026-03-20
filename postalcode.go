package validators

import (
	"fmt"
	"regexp"
	"strings"
)

var postalCodePatterns = map[string]*regexp.Regexp{
	"US": regexp.MustCompile(`^\d{5}(-\d{4})?$`),
	"CA": regexp.MustCompile(`^[A-Z]\d[A-Z]\s?\d[A-Z]\d$`),
	"GB": regexp.MustCompile(`^[A-Z]{1,2}\d[A-Z0-9]?\s?\d[A-Z]{2}$`),
	"DE": regexp.MustCompile(`^\d{5}$`),
	"FR": regexp.MustCompile(`^\d{5}$`),
	"IT": regexp.MustCompile(`^\d{5}$`),
	"ES": regexp.MustCompile(`^\d{5}$`),
	"PT": regexp.MustCompile(`^\d{4}(-\d{3})?$`),
	"NL": regexp.MustCompile(`^\d{4}\s?[A-Z]{2}$`),
	"BE": regexp.MustCompile(`^\d{4}$`),
	"LU": regexp.MustCompile(`^\d{4}$`),
	"CH": regexp.MustCompile(`^\d{4}$`),
	"AT": regexp.MustCompile(`^\d{4}$`),
	"DK": regexp.MustCompile(`^\d{4}$`),
	"NO": regexp.MustCompile(`^\d{4}$`),
	"SE": regexp.MustCompile(`^\d{3}\s?\d{2}$`),
	"FI": regexp.MustCompile(`^\d{5}$`),
	"PL": regexp.MustCompile(`^\d{2}-\d{3}$`),
	"CZ": regexp.MustCompile(`^\d{3}\s?\d{2}$`),
	"SK": regexp.MustCompile(`^\d{3}\s?\d{2}$`),
	"HU": regexp.MustCompile(`^\d{4}$`),
	"RO": regexp.MustCompile(`^\d{6}$`),
	"BG": regexp.MustCompile(`^\d{4}$`),
	"HR": regexp.MustCompile(`^\d{5}$`),
	"SI": regexp.MustCompile(`^\d{4}$`),
	"RS": regexp.MustCompile(`^\d{5}$`),
	"BA": regexp.MustCompile(`^\d{5}$`),
	"IE": regexp.MustCompile(`^[A-Z0-9]{3}\s?[A-Z0-9]{4}$`),
	"GR": regexp.MustCompile(`^\d{3}\s?\d{2}$`),
	"TR": regexp.MustCompile(`^\d{5}$`),
	"RU": regexp.MustCompile(`^\d{6}$`),
	"AU": regexp.MustCompile(`^\d{4}$`),
	"NZ": regexp.MustCompile(`^\d{4}$`),
	"JP": regexp.MustCompile(`^\d{3}-?\d{4}$`),
	"KR": regexp.MustCompile(`^\d{5}$`),
	"CN": regexp.MustCompile(`^\d{6}$`),
	"IN": regexp.MustCompile(`^\d{6}$`),
	"BR": regexp.MustCompile(`^\d{5}-?\d{3}$`),
	"MX": regexp.MustCompile(`^\d{5}$`),
	"AR": regexp.MustCompile(`^[A-Z]\d{4}[A-Z]{3}$`),
	"ZA": regexp.MustCompile(`^\d{4}$`),
	"IL": regexp.MustCompile(`^\d{7}$`),
	"SA": regexp.MustCompile(`^\d{5}(-\d{4})?$`),
	"AE": regexp.MustCompile(`^\d{5}$`), // Dubai and others; some emirates don't use postal codes
	"SG": regexp.MustCompile(`^\d{6}$`),
}

// PostalCode validates a postal/ZIP code for a given ISO 3166-1 alpha-2 country code.
func PostalCode(value, countryCode string) Result {
	if value == "" {
		return valid(nil)
	}

	cc := strings.ToUpper(strings.TrimSpace(countryCode))
	cleaned := strings.ToUpper(strings.TrimSpace(value))

	pattern, ok := postalCodePatterns[cc]
	if !ok {
		return invalid(ErrPostalCodeCountryInvalid,
			fmt.Sprintf("Postal code validation not supported for country: %s", cc),
			"postal_code", map[string]any{
				"value":        value,
				"country_code": cc,
			})
	}

	if !pattern.MatchString(cleaned) {
		return invalid(ErrPostalCodeFormatInvalid,
			fmt.Sprintf("Invalid postal code format for %s", cc),
			"postal_code", map[string]any{
				"value":        value,
				"country_code": cc,
			})
	}

	return valid(map[string]any{
		"country_code": cc,
	})
}
