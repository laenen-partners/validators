package validators

import (
	"fmt"
	"strings"
)

var currencies = map[string]bool{
	"AED": true, "AFN": true, "ALL": true, "AMD": true, "ANG": true, "AOA": true, "ARS": true, "AUD": true,
	"AWG": true, "AZN": true, "BAM": true, "BBD": true, "BDT": true, "BGN": true, "BHD": true, "BIF": true,
	"BMD": true, "BND": true, "BOB": true, "BRL": true, "BSD": true, "BTN": true, "BWP": true, "BYN": true,
	"BZD": true, "CAD": true, "CDF": true, "CHF": true, "CLP": true, "CNY": true, "COP": true, "CRC": true,
	"CUP": true, "CVE": true, "CZK": true, "DJF": true, "DKK": true, "DOP": true, "DZD": true, "EGP": true,
	"ERN": true, "ETB": true, "EUR": true, "FJD": true, "FKP": true, "GBP": true, "GEL": true, "GHS": true,
	"GIP": true, "GMD": true, "GNF": true, "GTQ": true, "GYD": true, "HKD": true, "HNL": true, "HTG": true,
	"HUF": true, "IDR": true, "ILS": true, "INR": true, "IQD": true, "IRR": true, "ISK": true, "JMD": true,
	"JOD": true, "JPY": true, "KES": true, "KGS": true, "KHR": true, "KMF": true, "KPW": true, "KRW": true,
	"KWD": true, "KYD": true, "KZT": true, "LAK": true, "LBP": true, "LKR": true, "LRD": true, "LSL": true,
	"LYD": true, "MAD": true, "MDL": true, "MGA": true, "MKD": true, "MMK": true, "MNT": true, "MOP": true,
	"MRU": true, "MUR": true, "MVR": true, "MWK": true, "MXN": true, "MYR": true, "MZN": true, "NAD": true,
	"NGN": true, "NIO": true, "NOK": true, "NPR": true, "NZD": true, "OMR": true, "PAB": true, "PEN": true,
	"PGK": true, "PHP": true, "PKR": true, "PLN": true, "PYG": true, "QAR": true, "RON": true, "RSD": true,
	"RUB": true, "RWF": true, "SAR": true, "SBD": true, "SCR": true, "SDG": true, "SEK": true, "SGD": true,
	"SHP": true, "SLE": true, "SOS": true, "SRD": true, "SSP": true, "STN": true, "SVC": true, "SYP": true,
	"SZL": true, "THB": true, "TJS": true, "TMT": true, "TND": true, "TOP": true, "TRY": true, "TTD": true,
	"TWD": true, "TZS": true, "UAH": true, "UGX": true, "USD": true, "UYU": true, "UZS": true, "VES": true,
	"VND": true, "VUV": true, "WST": true, "XAF": true, "XCD": true, "XOF": true, "XPF": true, "YER": true,
	"ZAR": true, "ZMW": true, "ZWL": true,
}

// CurrencyCode validates an ISO 4217 currency code.
func CurrencyCode(value string) Result {
	if value == "" {
		return valid(nil)
	}

	code := strings.ToUpper(strings.TrimSpace(value))

	if len(code) != 3 {
		return invalid(ErrCurrencyCodeFormatInvalid, "Currency code must be exactly 3 letters", "currency_code", map[string]any{
			"value":  value,
			"length": len(code),
		})
	}

	if !currencies[code] {
		return invalid(ErrCurrencyCodeUnknown,
			fmt.Sprintf("Unknown currency code: %s", code),
			"currency_code", map[string]any{
				"value": value,
			})
	}

	return valid(map[string]any{"code": code})
}
