package validators

import "testing"

func assertErrorCode(t *testing.T, r Result, code string) {
	t.Helper()
	if len(r.Errors) == 0 {
		t.Fatalf("expected error with code %q, got no errors", code)
	}
	if r.Errors[0].Code != code {
		t.Errorf("error code = %q, want %q", r.Errors[0].Code, code)
	}
}

func assertMetadata(t *testing.T, r Result, key string, want any) {
	t.Helper()
	got, ok := r.Metadata[key]
	if !ok {
		t.Fatalf("metadata missing key %q", key)
	}
	if got != want {
		t.Errorf("metadata[%q] = %v, want %v", key, got, want)
	}
}
