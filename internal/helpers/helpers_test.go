package helpers

import (
	"reflect"
	"testing"
)

func assertReviews(t testing.TB, got, want roam.RoamNodes) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertError(t testing.TB, err error, wantError bool) {
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
