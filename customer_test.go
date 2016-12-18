package braintree

import (
	"crypto/rand"
	"encoding/base64"
	"testing"
	"time"
)

func TestCreateCustomer(t *testing.T) {
	t.Parallel()

	t.Run("noID", func(t *testing.T) {
		t.Parallel()
		want := &Customer{FirstName: "first"}
		got, err := bt.Customer().Create(want)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
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
		got, err := bt.Customer().Create(customer)
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
		customer, err := bt.Customer().Create(customer)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if err := bt.Customer().Delete(customer.ID); err != nil {
			t.Errorf("unexpected err: %s", err)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if err := bt.Customer().Delete("cus2"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404: Not Found", err)
		}
	})
}

func TestFindCustomer(t *testing.T) {
	t.Parallel()

	t.Run("existing", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.Customer().Find("cus1")
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if customer == nil {
			t.Error("customer unexpected nil")
		}
		if size := len(customer.Addresses); size != 2 {
			t.Fatalf("addresses: got: %d, want: 2", size)
		}
		if gotLast := customer.Addresses[0].LastName; gotLast != "last" {
			t.Errorf("got: %s, want: last", gotLast)
		}
		if gotLast2 := customer.Addresses[1].LastName; gotLast2 != "last2" {
			t.Errorf("got: %s, want: last2", gotLast2)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.Customer().Find("cus2")
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

		want := &Customer{ID: "cus1", Phone: random(), CreatedAt: time.Time{}}
		got, err := bt.Customer().Update(want)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
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
		got, err := bt.Customer().Update(customer)
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
		got, err := bt.Customer().Update(customer)
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
