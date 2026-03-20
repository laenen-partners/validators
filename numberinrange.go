package validators

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

// NumberInRange validates that a decimal number string falls within [min, max].
// Uses exact decimal arithmetic (math/big.Rat) — no floating point errors.
// Either min or max may be empty to leave that bound open.
func NumberInRange(value, min, max string) Result {
	if value == "" {
		return valid(nil)
	}

	trimmed := strings.TrimSpace(value)
	val, ok := new(big.Rat).SetString(trimmed)
	if !ok {
		return invalid(ErrNumberFormatInvalid, "Invalid number format", "number", map[string]any{
			"value": value,
		})
	}

	if min != "" {
		minVal, ok := new(big.Rat).SetString(strings.TrimSpace(min))
		if !ok {
			return invalid(ErrNumberFormatInvalid, "Invalid min number format", "number", map[string]any{
				"min": min,
			})
		}
		if val.Cmp(minVal) < 0 {
			return invalid(ErrNumberBelowMin,
				fmt.Sprintf("Number must be at least %s, got %s", min, trimmed),
				"number", map[string]any{
					"value": trimmed,
					"min":   min,
				})
		}
	}

	if max != "" {
		maxVal, ok := new(big.Rat).SetString(strings.TrimSpace(max))
		if !ok {
			return invalid(ErrNumberFormatInvalid, "Invalid max number format", "number", map[string]any{
				"max": max,
			})
		}
		if val.Cmp(maxVal) > 0 {
			return invalid(ErrNumberAboveMax,
				fmt.Sprintf("Number must be at most %s, got %s", max, trimmed),
				"number", map[string]any{
					"value": trimmed,
					"max":   max,
				})
		}
	}

	return valid(map[string]any{
		"value": trimmed,
	})
}

// NumberInRangeFloat validates that a float64 falls within [min, max].
// Convenience wrapper for callers who already have numeric types.
// Uses floating point comparison — caller accepts precision tradeoffs.
func NumberInRangeFloat(value, min, max float64) Result {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return invalid(ErrNumberFormatInvalid, "Value must be a finite number", "number", map[string]any{
			"value": value,
		})
	}

	if !math.IsNaN(min) && value < min {
		return invalid(ErrNumberBelowMin,
			fmt.Sprintf("Number must be at least %g, got %g", min, value),
			"number", map[string]any{
				"value": value,
				"min":   min,
			})
	}

	if !math.IsNaN(max) && value > max {
		return invalid(ErrNumberAboveMax,
			fmt.Sprintf("Number must be at most %g, got %g", max, value),
			"number", map[string]any{
				"value": value,
				"max":   max,
			})
	}

	return valid(map[string]any{
		"value": value,
	})
}
