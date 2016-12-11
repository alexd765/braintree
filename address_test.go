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
		customer, err := bt.CreateCustomer(customer)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		address := &Address{CustomerID: customer.ID, StreetAddress: "street"}
		got, err := bt.CreateAddress(address)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		address.ID = got.ID
		if !reflect.DeepEqual(got, address) {
			t.Errorf("got: %+v\nwant: %+v", got, address)
		}
	})

	t.Run("withoutID", func(t *testing.T) {
		t.Parallel()

		address := &Address{StreetAddress: "street"}
		if _, err := bt.CreateAddress(address); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})
}
