package validators

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// BelgianNationalNumber validates a Belgian Rijksregisternummer (11 digits).
// Format: YY.MM.DD-SSS.CC where SSS is a sequence number and CC is a MOD-97 check.
func BelgianNationalNumber(value string) Result {
	if value == "" {
		return valid(nil)
	}

	// Strip dots, dashes, and spaces.
	cleaned := strings.Map(func(r rune) rune {
		if r == '.' || r == '-' || r == ' ' {
			return -1
		}
		return r
	}, value)

	if len(cleaned) != 11 {
		return invalid(ErrBNNLengthInvalid,
			fmt.Sprintf("Belgian national number must be 11 digits, got %d", len(cleaned)),
			"belgian_national_number", map[string]any{
				"value":  value,
				"length": len(cleaned),
			})
	}

	for _, c := range cleaned {
		if !unicode.IsDigit(c) {
			return invalid(ErrBNNFormatInvalid, "Belgian national number must contain only digits", "belgian_national_number", map[string]any{
				"value": value,
			})
		}
	}

	// Extract components.
	yy, _ := strconv.Atoi(cleaned[0:2])
	mm, _ := strconv.Atoi(cleaned[2:4])
	dd, _ := strconv.Atoi(cleaned[4:6])
	seq, _ := strconv.Atoi(cleaned[6:9])
	check, _ := strconv.Atoi(cleaned[9:11])

	// Validate month (00 is allowed for unknown, as are values > 20 for bis numbers).
	if mm > 12 && mm < 20 || mm > 32 {
		return invalid(ErrBNNDateInvalid, "Belgian national number has an invalid month", "belgian_national_number", map[string]any{
			"value": value,
			"month": mm,
		})
	}

	// Validate day.
	if dd > 31 {
		return invalid(ErrBNNDateInvalid, "Belgian national number has an invalid day", "belgian_national_number", map[string]any{
			"value": value,
			"day":   dd,
		})
	}

	// MOD-97 check: try with 19xx birth year first, then 20xx.
	base9 := cleaned[0:9]
	num19, _ := strconv.Atoi(base9)
	num20, _ := strconv.Atoi("2" + base9)

	expectedCheck19 := 97 - (num19 % 97)
	expectedCheck20 := 97 - (num20 % 97)

	born2000 := false
	if check == expectedCheck20 {
		born2000 = true
	} else if check != expectedCheck19 {
		return invalid(ErrBNNChecksumInvalid, "Invalid Belgian national number checksum", "belgian_national_number", map[string]any{
			"value": value,
		})
	}

	birthYear := 1900 + yy
	if born2000 {
		birthYear = 2000 + yy
	}

	meta := map[string]any{
		"birth_year": birthYear,
		"sequence":   seq,
	}
	if mm >= 1 && mm <= 12 {
		meta["birth_month"] = mm
	}
	if dd >= 1 && dd <= 31 {
		meta["birth_day"] = dd
	}
	// Odd sequence = male, even = female.
	if seq%2 == 1 {
		meta["gender"] = "male"
	} else {
		meta["gender"] = "female"
	}

	return valid(meta)
}
