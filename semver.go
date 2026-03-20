package validators

import (
	"regexp"
	"strconv"
	"strings"
)

var semverRegex = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

// SemVer validates a semantic version string per semver.org.
func SemVer(value string) Result {
	if value == "" {
		return valid(nil)
	}

	v := strings.TrimSpace(value)
	// Allow optional leading "v" prefix (common convention).
	v = strings.TrimPrefix(v, "v")

	matches := semverRegex.FindStringSubmatch(v)
	if matches == nil {
		return invalid(ErrSemVerFormatInvalid, "Invalid semantic version format", "semver", map[string]any{
			"value": value,
		})
	}

	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	patch, _ := strconv.Atoi(matches[3])

	meta := map[string]any{
		"major": major,
		"minor": minor,
		"patch": patch,
	}
	if matches[4] != "" {
		meta["prerelease"] = matches[4]
	}
	if matches[5] != "" {
		meta["build"] = matches[5]
	}
	return valid(meta)
}
