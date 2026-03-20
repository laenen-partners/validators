package validators

import (
	"fmt"
	"strings"
	"unicode"
)

// phoneLengths maps country calling codes to min/max subscriber digit counts (excluding the country code).
var phoneLengths = map[string][2]int{
	"1":   {10, 10}, // US, CA (includes country code in count for NANP)
	"7":   {10, 10}, // RU, KZ
	"20":  {10, 10}, // EG
	"27":  {9, 9},   // ZA
	"30":  {10, 10}, // GR
	"31":  {9, 9},   // NL
	"32":  {8, 9},   // BE
	"33":  {9, 9},   // FR
	"34":  {9, 9},   // ES
	"36":  {8, 9},   // HU
	"39":  {9, 11},  // IT
	"40":  {9, 9},   // RO
	"41":  {9, 9},   // CH
	"43":  {4, 13},  // AT
	"44":  {10, 10}, // GB
	"45":  {8, 8},   // DK
	"46":  {7, 13},  // SE
	"47":  {8, 8},   // NO
	"48":  {9, 9},   // PL
	"49":  {2, 13},  // DE (very variable)
	"51":  {9, 9},   // PE
	"52":  {10, 10}, // MX
	"53":  {8, 8},   // CU
	"54":  {10, 10}, // AR
	"55":  {10, 11}, // BR
	"56":  {9, 9},   // CL
	"57":  {10, 10}, // CO
	"58":  {10, 10}, // VE
	"60":  {9, 10},  // MY
	"61":  {9, 9},   // AU
	"62":  {9, 12},  // ID
	"63":  {10, 10}, // PH
	"64":  {8, 10},  // NZ
	"65":  {8, 8},   // SG
	"66":  {9, 9},   // TH
	"81":  {9, 10},  // JP
	"82":  {9, 10},  // KR
	"86":  {11, 11}, // CN
	"90":  {10, 10}, // TR
	"91":  {10, 10}, // IN
	"92":  {10, 10}, // PK
	"93":  {9, 9},   // AF
	"94":  {9, 9},   // LK
	"95":  {7, 10},  // MM
	"212": {9, 9},   // MA
	"213": {9, 9},   // DZ
	"216": {8, 8},   // TN
	"218": {9, 9},   // LY
	"220": {7, 7},   // GM
	"234": {10, 10}, // NG
	"254": {9, 9},   // KE
	"255": {9, 9},   // TZ
	"256": {9, 9},   // UG
	"260": {9, 9},   // ZM
	"263": {9, 9},   // ZW
	"351": {9, 9},   // PT
	"352": {4, 11},  // LU
	"353": {9, 9},   // IE
	"354": {7, 7},   // IS
	"358": {5, 12},  // FI
	"370": {8, 8},   // LT
	"371": {8, 8},   // LV
	"372": {7, 8},   // EE
	"380": {9, 9},   // UA
	"381": {8, 9},   // RS
	"385": {8, 9},   // HR
	"386": {8, 8},   // SI
	"387": {8, 8},   // BA
	"420": {9, 9},   // CZ
	"421": {9, 9},   // SK
	"852": {8, 8},   // HK
	"853": {8, 8},   // MO
	"855": {8, 9},   // KH
	"856": {8, 10},  // LA
	"880": {10, 10}, // BD
	"886": {9, 9},   // TW
	"960": {7, 7},   // MV
	"961": {7, 8},   // LB
	"962": {8, 9},   // JO
	"963": {8, 9},   // SY
	"964": {10, 10}, // IQ
	"965": {8, 8},   // KW
	"966": {9, 9},   // SA
	"968": {8, 8},   // OM
	"971": {9, 9},   // AE
	"972": {9, 9},   // IL
	"973": {8, 8},   // BH
	"974": {8, 8},   // QA
	"975": {8, 8},   // BT
	"976": {8, 8},   // MN
	"977": {10, 10}, // NP
	"992": {9, 9},   // TJ
	"993": {8, 8},   // TM
	"994": {9, 9},   // AZ
	"995": {9, 9},   // GE
	"996": {9, 9},   // KG
	"998": {9, 9},   // UZ
}

// Phone validates an international phone number in E.164-like format.
// Expects a leading + followed by country code and subscriber number.
func Phone(value string) Result {
	if value == "" {
		return valid(nil)
	}

	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || r == '+' {
			return r
		}
		if r == ' ' || r == '-' || r == '(' || r == ')' || r == '.' {
			return -1 // strip common formatting
		}
		return r
	}, value)

	if !strings.HasPrefix(cleaned, "+") {
		return invalid(ErrPhoneFormatInvalid, "Phone number must start with + and country code", "phone", map[string]any{
			"value": value,
		})
	}

	digits := cleaned[1:]
	for _, c := range digits {
		if !unicode.IsDigit(c) {
			return invalid(ErrPhoneInvalidChars, "Phone number contains invalid characters", "phone", map[string]any{
				"value": value,
			})
		}
	}

	if len(digits) < 7 {
		return invalid(ErrPhoneTooShort, "Phone number is too short", "phone", map[string]any{
			"value":  value,
			"digits": len(digits),
		})
	}
	if len(digits) > 15 {
		return invalid(ErrPhoneTooLong, "Phone number is too long (max 15 digits per E.164)", "phone", map[string]any{
			"value":  value,
			"digits": len(digits),
		})
	}

	// Try to match a country code (1-3 digits) and validate length.
	var countryCode string
	for _, length := range []int{1, 2, 3} {
		if length > len(digits) {
			break
		}
		cc := digits[:length]
		if _, ok := phoneLengths[cc]; ok {
			countryCode = cc
			break
		}
	}

	if countryCode != "" {
		bounds := phoneLengths[countryCode]
		subscriberLen := len(digits) - len(countryCode)
		if subscriberLen < bounds[0] || subscriberLen > bounds[1] {
			return invalid(ErrPhoneCountryInvalid,
				fmt.Sprintf("Phone number has invalid length for country code +%s", countryCode),
				"phone", map[string]any{
					"value":            value,
					"country_code":     "+" + countryCode,
					"subscriber_digits": subscriberLen,
					"expected_min":     bounds[0],
					"expected_max":     bounds[1],
				})
		}
	}

	meta := map[string]any{
		"digits": len(digits),
	}
	if countryCode != "" {
		meta["country_code"] = "+" + countryCode
	}
	return valid(meta)
}
