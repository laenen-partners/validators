package validators

import (
	"fmt"
	"regexp"
	"strings"
)

// vatPatterns maps EU country codes to their VAT number regex patterns.
var vatPatterns = map[string]*regexp.Regexp{
	"AT": regexp.MustCompile(`^ATU\d{8}$`),
	"BE": regexp.MustCompile(`^BE[01]\d{9}$`),
	"BG": regexp.MustCompile(`^BG\d{9,10}$`),
	"CY": regexp.MustCompile(`^CY\d{8}[A-Z]$`),
	"CZ": regexp.MustCompile(`^CZ\d{8,10}$`),
	"DE": regexp.MustCompile(`^DE\d{9}$`),
	"DK": regexp.MustCompile(`^DK\d{8}$`),
	"EE": regexp.MustCompile(`^EE\d{9}$`),
	"EL": regexp.MustCompile(`^EL\d{9}$`),
	"ES": regexp.MustCompile(`^ES[A-Z0-9]\d{7}[A-Z0-9]$`),
	"FI": regexp.MustCompile(`^FI\d{8}$`),
	"FR": regexp.MustCompile(`^FR[A-Z0-9]{2}\d{9}$`),
	"HR": regexp.MustCompile(`^HR\d{11}$`),
	"HU": regexp.MustCompile(`^HU\d{8}$`),
	"IE": regexp.MustCompile(`^IE\d[A-Z0-9+*]\d{5}[A-Z]{1,2}$`),
	"IT": regexp.MustCompile(`^IT\d{11}$`),
	"LT": regexp.MustCompile(`^LT(\d{9}|\d{12})$`),
	"LU": regexp.MustCompile(`^LU\d{8}$`),
	"LV": regexp.MustCompile(`^LV\d{11}$`),
	"MT": regexp.MustCompile(`^MT\d{8}$`),
	"NL": regexp.MustCompile(`^NL\d{9}B\d{2}$`),
	"PL": regexp.MustCompile(`^PL\d{10}$`),
	"PT": regexp.MustCompile(`^PT\d{9}$`),
	"RO": regexp.MustCompile(`^RO\d{2,10}$`),
	"SE": regexp.MustCompile(`^SE\d{12}$`),
	"SI": regexp.MustCompile(`^SI\d{8}$`),
	"SK": regexp.MustCompile(`^SK\d{10}$`),
	"XI": regexp.MustCompile(`^XI\d{9}$`),        // Northern Ireland
	"EU": regexp.MustCompile(`^EU\d{9}$`),         // EU VAT for non-established
	"CH": regexp.MustCompile(`^CHE\d{9}(TVA|MWST|IVA)$`), // Switzerland (not EU but commonly needed)
}

// VAT validates a European VAT identification number.
func VAT(value string) Result {
	if value == "" {
		return valid(nil)
	}

	cleaned := strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(value, " ", ""), ".", ""))

	if len(cleaned) < 4 {
		return invalid(ErrVATTooShort, "VAT number is too short", "vat", map[string]any{
			"value": value,
		})
	}

	// Extract 2-letter country prefix (or 3 for CHE).
	var cc string
	if strings.HasPrefix(cleaned, "CHE") {
		cc = "CH"
	} else {
		prefix := cleaned[:2]
		for _, c := range prefix {
			if c < 'A' || c > 'Z' {
				return invalid(ErrVATCountryInvalid, "VAT number must start with a country code", "vat", map[string]any{
					"value": value,
				})
			}
		}
		// Greece uses EL instead of GR.
		cc = prefix
	}

	pattern, ok := vatPatterns[cc]
	if !ok {
		return invalid(ErrVATCountryInvalid,
			fmt.Sprintf("Unsupported VAT country code: %s", cc),
			"vat", map[string]any{
				"value":        value,
				"country_code": cc,
			})
	}

	if !pattern.MatchString(cleaned) {
		return invalid(ErrVATFormatInvalid,
			fmt.Sprintf("Invalid VAT number format for %s", cc),
			"vat", map[string]any{
				"value":        value,
				"country_code": cc,
			})
	}

	return valid(map[string]any{
		"country_code": cc,
	})
}
