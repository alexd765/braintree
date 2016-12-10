package braintree

import (
	"crypto/rand"
	"encoding/base64"
	"testing"
)

var bt = New()

func TestCreateCustomer(t *testing.T) {
	t.Parallel()

	t.Run("noID", func(t *testing.T) {
		t.Parallel()
		want := &Customer{FirstName: "first"}
		got, err := bt.CreateCustomer(want)
		if err != nil {
			t.Errorf("unexpected err: %s", err)
		}
		if got.FirstName != want.FirstName {
			t.Errorf("FirstName: got: %s, want: %s", got.FirstName, want.FirstName)
		}
		if got.ID == "" {
			t.Errorf("ID: got empty, want nonempty")
		}
	})

	t.Run("existing", func(t *testing.T) {
		t.Parallel()
		customer := &Customer{ID: "cus1", FirstName: "first"}
		got, err := bt.CreateCustomer(customer)
		if err == nil || err.Error() != "422 Unprocessable Entity" {
			t.Errorf("got: %v, want: 422 Unprocessable Entity", err)
		}
		if got != nil {
			t.Errorf("got: %+v, want: <nil>", got)
		}
	})
}

func TestDeleteCustomer(t *testing.T) {
	t.Parallel()

	t.Run("existing", func(t *testing.T) {
		t.Parallel()

		customer := &Customer{FirstName: "first"}
		customer, err := bt.CreateCustomer(customer)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if err := bt.DeleteCustomer(customer.ID); err != nil {
			t.Errorf("unexpected err: %s", err)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if err := bt.DeleteCustomer("cus2"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404: Not Found", err)
		}
	})
}

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
		if err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404: Not Found", err)
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
		if err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
		if got != nil {
			t.Errorf("got: %+v, want: <nil>", got)
		}
	})

	t.Run("noID", func(t *testing.T) {
		t.Parallel()

		customer := &Customer{Phone: random()}
		got, err := bt.UpdateCustomer(customer)
		if err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
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
