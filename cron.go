package validators

import (
	"fmt"
	"strconv"
	"strings"
)

var cronFieldNames = []string{"minute", "hour", "day_of_month", "month", "day_of_week"}
var cronFieldRanges = [][2]int{
	{0, 59},  // minute
	{0, 23},  // hour
	{1, 31},  // day of month
	{1, 12},  // month
	{0, 7},   // day of week (0 and 7 both = Sunday)
}

// CRON validates a 5-field cron expression (minute hour day month weekday).
func CRON(value string) Result {
	if value == "" {
		return valid(nil)
	}

	trimmed := strings.TrimSpace(value)
	fields := strings.Fields(trimmed)

	if len(fields) != 5 {
		return invalid(ErrCRONFormatInvalid,
			fmt.Sprintf("CRON expression must have 5 fields, got %d", len(fields)),
			"cron", map[string]any{
				"value":  value,
				"fields": len(fields),
			})
	}

	for i, field := range fields {
		if err := validateCronField(field, cronFieldRanges[i][0], cronFieldRanges[i][1]); err != "" {
			return invalid(ErrCRONFieldInvalid,
				fmt.Sprintf("Invalid %s field: %s", cronFieldNames[i], err),
				"cron", map[string]any{
					"value":      value,
					"field_name": cronFieldNames[i],
					"field":      field,
					"reason":     err,
				})
		}
	}

	return valid(map[string]any{
		"minute":       fields[0],
		"hour":         fields[1],
		"day_of_month": fields[2],
		"month":        fields[3],
		"day_of_week":  fields[4],
	})
}

func validateCronField(field string, min, max int) string {
	// Handle comma-separated lists.
	parts := strings.Split(field, ",")
	for _, part := range parts {
		if err := validateCronPart(part, min, max); err != "" {
			return err
		}
	}
	return ""
}

func validateCronPart(part string, min, max int) string {
	if part == "*" {
		return ""
	}

	// Handle step values: */n or range/n.
	if strings.Contains(part, "/") {
		segments := strings.SplitN(part, "/", 2)
		if segments[0] != "*" {
			if err := validateCronRange(segments[0], min, max); err != "" {
				return err
			}
		}
		step, err := strconv.Atoi(segments[1])
		if err != nil || step < 1 {
			return fmt.Sprintf("invalid step value: %s", segments[1])
		}
		return ""
	}

	// Handle ranges: n-m.
	return validateCronRange(part, min, max)
}

func validateCronRange(part string, min, max int) string {
	if strings.Contains(part, "-") {
		bounds := strings.SplitN(part, "-", 2)
		lo, err1 := strconv.Atoi(bounds[0])
		hi, err2 := strconv.Atoi(bounds[1])
		if err1 != nil || err2 != nil {
			return fmt.Sprintf("non-numeric range: %s", part)
		}
		if lo < min || lo > max || hi < min || hi > max {
			return fmt.Sprintf("value out of range %d-%d: %s", min, max, part)
		}
		if lo > hi {
			return fmt.Sprintf("range start %d exceeds end %d", lo, hi)
		}
		return ""
	}

	n, err := strconv.Atoi(part)
	if err != nil {
		return fmt.Sprintf("non-numeric value: %s", part)
	}
	if n < min || n > max {
		return fmt.Sprintf("value %d out of range %d-%d", n, min, max)
	}
	return ""
}
