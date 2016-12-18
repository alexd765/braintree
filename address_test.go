package braintree

import (
	"reflect"
	"testing"
)

func TestCreateAddress(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()

		customer := &Customer{FirstName: "test", LastName: "create address"}
		customer, err := bt.Customer().Create(customer)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		address := &Address{CustomerID: customer.ID, StreetAddress: "street"}
		got, err := bt.Address().Create(address)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		address.ID = got.ID
		address.CreatedAt = got.CreatedAt
		address.UpdatedAt = got.UpdatedAt
		if !reflect.DeepEqual(got, address) {
			t.Errorf("got: %+v\nwant: %+v", got, address)
		}
	})

	t.Run("withoutID", func(t *testing.T) {
		t.Parallel()

		address := &Address{StreetAddress: "street"}
		if _, err := bt.Address().Create(address); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})
}

func TestFindAddress(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()

		got, err := bt.Address().Find("cus1", "qd")
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		want := &Address{CustomerID: "cus1", ID: "qd", FirstName: "first", LastName: "last"}
		want.CreatedAt = got.CreatedAt
		want.UpdatedAt = got.UpdatedAt
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %+v\nwant: %+v", got, want)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if _, err := bt.Address().Find("cus1", "bla"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})

}
