package validators

import (
	"fmt"
	"math"
)

// LatLon validates a geographic coordinate pair (latitude, longitude) in decimal degrees.
func LatLon(lat, lon float64) Result {
	if math.IsNaN(lat) || math.IsNaN(lon) || math.IsInf(lat, 0) || math.IsInf(lon, 0) {
		return invalid(ErrLatLonFormatInvalid, "Latitude and longitude must be finite numbers", "latlon", map[string]any{
			"latitude":  lat,
			"longitude": lon,
		})
	}

	if lat < -90 || lat > 90 {
		return invalid(ErrLatitudeOutOfRange,
			fmt.Sprintf("Latitude must be between -90 and 90, got %g", lat),
			"latlon", map[string]any{
				"latitude":  lat,
				"longitude": lon,
			})
	}

	if lon < -180 || lon > 180 {
		return invalid(ErrLongitudeOutOfRange,
			fmt.Sprintf("Longitude must be between -180 and 180, got %g", lon),
			"latlon", map[string]any{
				"latitude":  lat,
				"longitude": lon,
			})
	}

	return valid(map[string]any{
		"latitude":  lat,
		"longitude": lon,
	})
}
