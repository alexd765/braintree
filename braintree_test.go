package braintree

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

var bt *Braintree

func TestMain(m *testing.M) {
	bt = New()
	os.Exit(m.Run())
}

func TestLogger(t *testing.T) {
	logs := new(bytes.Buffer)
	bt2 := New()
	bt2.Logger = log.New(logs, "bt: ", 0)
	customer := &Customer{FirstName: "AA", LastName: "BB"}
	_, err := bt2.Customer().Create(customer)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	logsStr := logs.String()
	wantPre := "bt: >>> POST https://sandbox.braintreegateway.com/merchants/"
	wantSuff := "with payload: <Customer><addresses></addresses><first-name>AA</first-name><last-name>BB</last-name></Customer>\n"
	if !strings.HasPrefix(logsStr, wantPre) {
		t.Errorf("got: %s, want prefix: %s", logsStr, wantPre)
	}
	if !strings.HasSuffix(logsStr, wantSuff) {
		t.Errorf("got: %s, want suffix: %s", logsStr, wantSuff)
	}
}
