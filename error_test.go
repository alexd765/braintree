package braintree

import "testing"

func TestError(t *testing.T) {
	err := &APIError{Code: 5, Message: "Johnny"}
	got := err.Error()
	want := "Code 5: Johnny"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
