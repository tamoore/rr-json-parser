package roam

import (
	"reflect"
	"testing"
)

func AssertRoamNodes(t testing.TB, got, want RoamNodes) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
