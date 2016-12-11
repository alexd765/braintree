package braintree

import (
	"os"
	"testing"
)

var bt *Braintree

func TestMain(m *testing.M) {
	bt = New()
	os.Exit(m.Run())
}
