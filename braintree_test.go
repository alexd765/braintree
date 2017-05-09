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
	var err error
	bt, err = New()
	if err != nil {
		log.Fatalf("error %s", err)
	}
	os.Exit(m.Run())
}

// doesn't run in parallel to other tests
func TestNew(t *testing.T) {

	for _, key := range [3]string{"BRAINTREE_MERCHANT_ID", "BRAINTREE_PUBLIC_KEY", "BRAINTREE_PRIVATE_KEY"} {
		value := os.Getenv(key)
		os.Setenv(key, "")
		want := "env " + key + " not set"
		if _, err := New(); err == nil || err.Error() != want {
			t.Errorf("got: %v, want: %s", err, want)
		}
		os.Setenv(key, value)
	}
}

func TestLogger(t *testing.T) {
	t.Parallel()
	logs := new(bytes.Buffer)
	bt2, err := New()
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	bt2.Logger = log.New(logs, "bt: ", 0)
	_, err = bt2.Customer().Create(CustomerInput{FirstName: "AA", LastName: "BB"})
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	logsStr := logs.String()
	wantPre := "bt: >>> POST https://sandbox.braintreegateway.com/merchants/"
	wantSuff := "with payload: <customer><first-name>AA</first-name><last-name>BB</last-name></customer>\n"
	if !strings.HasPrefix(logsStr, wantPre) {
		t.Errorf("got: %s, want prefix: %s", logsStr, wantPre)
	}
	if !strings.HasSuffix(logsStr, wantSuff) {
		t.Errorf("got: %s, want suffix: %s", logsStr, wantSuff)
	}
}
