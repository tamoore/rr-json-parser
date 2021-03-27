package test

import (
	"testing"
)

func AssertError(t testing.TB, err error, wantError bool) {
	t.Helper()
	if err == nil {
		if wantError {
			t.Errorf("wanted error but recv nil: %v", err)
		}
	}

	if err != nil {
		if !wantError {
			t.Errorf("did not want error but recv: %v", err)
		}
	}
}
