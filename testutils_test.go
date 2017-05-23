package braintree

import (
	"reflect"
	"testing"
)

func compareErrors(t *testing.T, got, want error) {
	if got == want {
		return
	}
	if g, w := reflect.TypeOf(got), reflect.TypeOf(want); g != w {
		t.Errorf("error types: got '%v'; want '%v'", g, w)
	}
	if got == nil || want == nil || got.Error() != want.Error() {
		t.Errorf("error: got '%v'; want '%v'", got, want)
	}
}
