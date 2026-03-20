package validators

import (
	"net/url"
	"strings"
)

// allowedSchemes are the URL schemes considered valid.
var allowedSchemes = map[string]bool{
	"http":  true,
	"https": true,
	"ftp":   true,
	"ftps":  true,
}

// URL validates a URL string.
// Requires a scheme (http, https, ftp, ftps) and a host.
func URL(value string) Result {
	if value == "" {
		return valid(nil)
	}

	parsed, err := url.Parse(value)
	if err != nil {
		return invalid(ErrURLFormatInvalid, "Invalid URL format", "url", map[string]any{
			"value": value,
		})
	}

	scheme := strings.ToLower(parsed.Scheme)
	if scheme == "" || !allowedSchemes[scheme] {
		return invalid(ErrURLSchemeInvalid, "URL must use http, https, ftp, or ftps scheme", "url", map[string]any{
			"value":  value,
			"scheme": parsed.Scheme,
		})
	}

	if parsed.Host == "" {
		return invalid(ErrURLHostMissing, "URL must include a host", "url", map[string]any{
			"value": value,
		})
	}

	meta := map[string]any{
		"scheme": scheme,
		"host":   parsed.Host,
	}
	if parsed.Port() != "" {
		meta["port"] = parsed.Port()
	}
	if parsed.Path != "" && parsed.Path != "/" {
		meta["path"] = parsed.Path
	}
	return valid(meta)
}
