package braintree

import (
	"crypto/rand"
	"encoding/base64"
	"testing"
)

var bt = New()

func TestFindCustomer(t *testing.T) {
	t.Parallel()

	t.Run("existing", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.FindCustomer("cus1")
		if err != nil {
			t.Errorf("unexpected err: %s", err)
		}
		if customer == nil {
			t.Error("customer unexpected nil")
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.FindCustomer("cus2")
		if err == nil || err.Error() != "404: not found" {
			t.Errorf("got: %v, want: 404: not found", err)
		}
		if customer != nil {
			t.Errorf("got: %+v, want: <nil>", customer)
		}
	})
}

func TestUpdateCustomer(t *testing.T) {
	t.Parallel()

	t.Run("existing", func(t *testing.T) {
		t.Parallel()

		want := &Customer{ID: "cus1", Phone: random()}
		got, err := bt.UpdateCustomer(want)
		if err != nil {
			t.Errorf("unexpected err: %s", err)
		}
		if got == nil {
			t.Fatal("customer unexpected nil")
		}
		if got.Phone != want.Phone {
			t.Errorf("got: %s, want: %s", got.Phone, want.Phone)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		customer := &Customer{ID: "cus2", Phone: random()}
		got, err := bt.UpdateCustomer(customer)
		if err == nil || err.Error() != "404: not found" {
			t.Errorf("got: %v, want: 404: not found", err)
		}
		if got != nil {
			t.Errorf("got: %+v, want: <nil>", got)
		}
	})

	t.Run("noID", func(t *testing.T) {
		t.Parallel()

		customer := &Customer{Phone: random()}
		got, err := bt.UpdateCustomer(customer)
		if err == nil || err.Error() != "404: not found" {
			t.Errorf("got: %v, want: 404: not found", err)
		}
		if got != nil {
			t.Errorf("got: %+v, want: <nil>", got)
		}
	})
}

func random() string {
	b := make([]byte, 8)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
