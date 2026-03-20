package validators

import (
	"math"
	"testing"
)

func TestLatLon(t *testing.T) {
	tests := []struct {
		name  string
		lat   float64
		lon   float64
		valid bool
		code  string
	}{
		{"valid origin", 0, 0, true, ""},
		{"valid Brussels", 50.8503, 4.3517, true, ""},
		{"valid min", -90, -180, true, ""},
		{"valid max", 90, 180, true, ""},
		{"lat too low", -91, 0, false, ErrLatitudeOutOfRange},
		{"lat too high", 91, 0, false, ErrLatitudeOutOfRange},
		{"lon too low", 0, -181, false, ErrLongitudeOutOfRange},
		{"lon too high", 0, 181, false, ErrLongitudeOutOfRange},
		{"NaN lat", math.NaN(), 0, false, ErrLatLonFormatInvalid},
		{"Inf lon", 0, math.Inf(1), false, ErrLatLonFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := LatLon(tt.lat, tt.lon)
			if r.Valid != tt.valid {
				t.Errorf("LatLon(%v, %v) valid=%v, want %v", tt.lat, tt.lon, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}
