package validators

import (
	"strconv"
	"strings"
	"unicode"
)

// CreditCard validates a credit card number using the Luhn algorithm.
// Detects the card network from the IIN prefix.
func CreditCard(value string) Result {
	if value == "" {
		return valid(nil)
	}

	// Strip spaces and dashes.
	cleaned := strings.Map(func(r rune) rune {
		if r == ' ' || r == '-' {
			return -1
		}
		return r
	}, value)

	for _, c := range cleaned {
		if !unicode.IsDigit(c) {
			return invalid(ErrCreditCardInvalidChars, "Credit card number contains invalid characters", "credit_card", map[string]any{
				"value": value,
			})
		}
	}

	if len(cleaned) < 12 {
		return invalid(ErrCreditCardTooShort, "Credit card number is too short", "credit_card", map[string]any{
			"value":  value,
			"length": len(cleaned),
		})
	}
	if len(cleaned) > 19 {
		return invalid(ErrCreditCardTooLong, "Credit card number is too long", "credit_card", map[string]any{
			"value":  value,
			"length": len(cleaned),
		})
	}

	// Luhn checksum.
	var sum int
	double := false
	for i := len(cleaned) - 1; i >= 0; i-- {
		d := int(cleaned[i] - '0')
		if double {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	if sum%10 != 0 {
		return invalid(ErrCreditCardChecksumInvalid, "Credit card number failed Luhn check", "credit_card", map[string]any{
			"value": value,
		})
	}

	network := detectCardNetwork(cleaned)
	meta := map[string]any{
		"length": len(cleaned),
	}
	if network != "" {
		meta["network"] = network
	}
	return valid(meta)
}

func detectCardNetwork(number string) string {
	switch {
	case strings.HasPrefix(number, "4"):
		return "visa"
	case hasPrefixRange(number, 51, 55) || hasPrefixRange(number, 2221, 2720):
		return "mastercard"
	case strings.HasPrefix(number, "34") || strings.HasPrefix(number, "37"):
		return "amex"
	case strings.HasPrefix(number, "6011") || strings.HasPrefix(number, "65") || hasPrefixRange(number, 644, 649):
		return "discover"
	case strings.HasPrefix(number, "3528") || hasPrefixRange(number, 3528, 3589):
		return "jcb"
	case strings.HasPrefix(number, "36") || strings.HasPrefix(number, "300") || strings.HasPrefix(number, "301") ||
		strings.HasPrefix(number, "302") || strings.HasPrefix(number, "303") || strings.HasPrefix(number, "304") ||
		strings.HasPrefix(number, "305"):
		return "diners"
	case strings.HasPrefix(number, "62"):
		return "unionpay"
	default:
		return ""
	}
}

func hasPrefixRange(number string, low, high int) bool {
	s := strconv.Itoa(high)
	width := len(s)
	if len(number) < width {
		return false
	}
	prefix, err := strconv.Atoi(number[:width])
	if err != nil {
		return false
	}
	return prefix >= low && prefix <= high
}
