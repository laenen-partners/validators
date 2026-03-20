package validators

import (
	"fmt"
	"testing"
	"time"
)

func TestAgeAtLeast(t *testing.T) {
	now := time.Now()
	age18 := fmt.Sprintf("%04d-%02d-%02d", now.Year()-18, now.Month(), now.Day())
	age17 := fmt.Sprintf("%04d-%02d-%02d", now.Year()-17, now.Month(), now.Day())
	age21 := fmt.Sprintf("%04d-%02d-%02d", now.Year()-25, now.Month(), now.Day())

	tests := []struct {
		name     string
		value    string
		minYears int
		valid    bool
		code     string
	}{
		{"empty", "", 18, true, ""},
		{"exactly 18", age18, 18, true, ""},
		{"over 21", age21, 21, true, ""},
		{"under 18", age17, 18, false, ErrAgeAtLeastTooYoung},
		{"bad format", "not-a-date", 18, false, ErrDateFormatInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := AgeAtLeast(tt.value, tt.minYears)
			if r.Valid != tt.valid {
				t.Errorf("AgeAtLeast(%q, %d) valid=%v, want %v", tt.value, tt.minYears, r.Valid, tt.valid)
			}
			if !tt.valid {
				assertErrorCode(t, r, tt.code)
			}
		})
	}
}

func TestAgeAtLeast_Metadata(t *testing.T) {
	birthDate := fmt.Sprintf("%04d-06-15", time.Now().Year()-30)
	r := AgeAtLeast(birthDate, 18)
	if !r.Valid {
		t.Fatal("expected valid")
	}
	age, ok := r.Metadata["age"].(int)
	if !ok || age < 29 {
		t.Errorf("expected age >= 29, got %v", r.Metadata["age"])
	}
}
