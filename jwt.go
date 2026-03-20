package validators

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

// JWT validates the structure of a JSON Web Token (three base64url-encoded segments).
// It does NOT verify the cryptographic signature.
func JWT(value string) Result {
	if value == "" {
		return valid(nil)
	}

	trimmed := strings.TrimSpace(value)
	parts := strings.Split(trimmed, ".")

	if len(parts) != 3 {
		return invalid(ErrJWTFormatInvalid, "JWT must have exactly 3 segments separated by dots", "jwt", map[string]any{
			"value":    value,
			"segments": len(parts),
		})
	}

	for i, name := range []string{"header", "payload", "signature"} {
		if parts[i] == "" {
			return invalid(ErrJWTSegmentInvalid, "JWT "+name+" segment is empty", "jwt", map[string]any{
				"value":   value,
				"segment": name,
			})
		}
	}

	// Validate header is decodable JSON with "alg".
	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return invalid(ErrJWTSegmentInvalid, "JWT header is not valid base64url", "jwt", map[string]any{
			"value":   value,
			"segment": "header",
		})
	}

	var header map[string]any
	if err := json.Unmarshal(headerBytes, &header); err != nil {
		return invalid(ErrJWTSegmentInvalid, "JWT header is not valid JSON", "jwt", map[string]any{
			"value":   value,
			"segment": "header",
		})
	}

	alg, _ := header["alg"].(string)
	if alg == "" {
		return invalid(ErrJWTSegmentInvalid, "JWT header missing 'alg' field", "jwt", map[string]any{
			"value":   value,
			"segment": "header",
		})
	}

	// Validate payload is decodable base64url JSON.
	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return invalid(ErrJWTSegmentInvalid, "JWT payload is not valid base64url", "jwt", map[string]any{
			"value":   value,
			"segment": "payload",
		})
	}

	var payload map[string]any
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return invalid(ErrJWTSegmentInvalid, "JWT payload is not valid JSON", "jwt", map[string]any{
			"value":   value,
			"segment": "payload",
		})
	}

	// Validate signature is decodable base64url.
	if _, err := base64.RawURLEncoding.DecodeString(parts[2]); err != nil {
		return invalid(ErrJWTSegmentInvalid, "JWT signature is not valid base64url", "jwt", map[string]any{
			"value":   value,
			"segment": "signature",
		})
	}

	meta := map[string]any{
		"algorithm": alg,
	}
	if typ, ok := header["typ"].(string); ok {
		meta["type"] = typ
	}
	return valid(meta)
}
