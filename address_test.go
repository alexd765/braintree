package braintree

import (
	"reflect"
	"testing"
)

func TestCreateAddress(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.Customer().Create(CustomerInput{FirstName: "test", LastName: "create address"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		got, err := bt.Address().Create(customer.ID, AddressInput{StreetAddress: "street"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if got.StreetAddress != "street" {
			t.Errorf("StreetAddress: got %s, want street", got.StreetAddress)
		}
	})

	t.Run("withoutID", func(t *testing.T) {
		t.Parallel()

		if _, err := bt.Address().Create("", AddressInput{StreetAddress: "street"}); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		_, err := bt.Address().Create("cus1", AddressInput{})
		apiErr, ok := err.(*APIError)
		if !ok {
			t.Fatalf("expected APIError")
		}
		if apiErr == nil || apiErr.Code != 81801 {
			t.Errorf("got %v, want error code 81801", apiErr)
		}
	})
}

func TestDeleteAddress(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()
		customer, err := bt.Customer().Create(CustomerInput{FirstName: "test", LastName: "create address"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		address, err := bt.Address().Create(customer.ID, AddressInput{StreetAddress: "street"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		if err := bt.Address().Delete(customer.ID, address.ID); err != nil {
			t.Errorf("unexpected err: %s", err)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if err := bt.Address().Delete("cus1", "bla"); err == nil || err.Error() != "404 Not Found" {
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

func TestUpdateAddress(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()
		customer, err := bt.Customer().Create(CustomerInput{FirstName: "test", LastName: "update address"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		address, err := bt.Address().Create(customer.ID, AddressInput{StreetAddress: "street"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		got, err := bt.Address().Update(customer.ID, address.ID, AddressInput{FirstName: "test2"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if got.FirstName != "test2" {
			t.Errorf("FirstName: got %s, want test2", got.FirstName)
		}
		if got.StreetAddress != "street" {
			t.Errorf("StreetAddress: got %s, want street", got.StreetAddress)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if _, err := bt.Address().Update("cus1", "bla", AddressInput{FirstName: "first"}); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})
}
